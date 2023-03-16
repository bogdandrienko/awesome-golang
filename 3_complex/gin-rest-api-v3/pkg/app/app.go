package app

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type ConfigS struct {
	// Gin
	DebugGin bool
	HostGin  string
	PortGin  int

	// Postgresql
	DriverNamePgSQL string
	HostPgSQL       string
	PortPgSQL       int
	UserPgSQL       string
	PasswordPgSQL   string
	DatabasePgSQL   string
	SslModePgSQL    string

	// Additional
	HashCost             int
	HashSalt             string
	TokenAccessLifetime  time.Duration
	TokenRefreshLifetime time.Duration
}

type RouteS struct {
	Method       string
	RelativePath string
	Function     gin.HandlerFunc
	Auth         bool
}

type User struct {
	// main
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`

	// roles
	IsModerator bool `json:"is_moderator"`
	IsAdmin     bool `json:"is_admin"`
	IsSuperuser bool `json:"is_superuser"`

	// additional
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Patronymic        string    `json:"patronymic"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Avatar            string    `json:"avatar"`
	DatetimeJoined    time.Time `json:"datetime_joined"`
	DatetimeLastLogin time.Time `json:"datetime_last_login"`
}

type Token struct {
	// main
	Id       int    `json:"id"`
	Username string `json:"username"`
	Access   string `json:"access"`
	Refresh  string `json:"refresh"`

	// additional
	DatetimeCreated time.Time `json:"datetime_created"`
}

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

var ConfigV = ConfigConstructor()

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewApp(routes []RouteS) error {
	// todo Configure engine
	if ConfigV.DebugGin {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// todo Create engine
	//engine := gin.New()
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
		//AllowAllOrigins: true,
		AllowOrigins: []string{"*"},
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "http://127.0.0.1:3000"
		//},
		//AllowMethods:     []string{"PUT", "PATCH"},
		//AllowHeaders:     []string{"Origin"},
		//ExposeHeaders:    []string{"Content-Length"},
		//AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// todo Binding routes
	for _, route := range routes {
		switch route.Method {
		case "POST": // Create
			engine.POST(route.RelativePath, Middleware(route.Function, route))
		case "GET": // Read
			engine.GET(route.RelativePath, Middleware(route.Function, route))
		case "PUT": // Update
			engine.PUT(route.RelativePath, Middleware(route.Function, route))
		case "DELETE": // Delete
			engine.DELETE(route.RelativePath, Middleware(route.Function, route))
		default:
			ErrorHandler(errors.New("unknown method"))
		}
	}

	// todo Run engine
	return engine.Run(fmt.Sprintf("%s:%d", ConfigV.HostGin, ConfigV.PortGin))
}

func Middleware(next gin.HandlerFunc, route RouteS) gin.HandlerFunc {

	time.Sleep(time.Millisecond * 1500)

	return func(context *gin.Context) {
		//log.Printf("[%s] %s\n", context.Request.Method, context.Request.RequestURI)

		if route.Auth {
			// get config.headers.Authorization
			authorizationHeader := context.GetHeader("Authorization")
			if authorizationHeader == "" {
				ErrHandlerWithContext(errors.New("authorization failed"), context)
				return
			}

			// get JWT=$Qwerty!1234567
			tokenArr := strings.Split(authorizationHeader, "=")
			if len(tokenArr) < 2 {
				ErrHandlerWithContext(errors.New("authorization failed"), context)
				return
			}

			// check token
			status, err := CheckAccessToken(tokenArr[1])
			if err != nil {
				ErrHandlerWithContext(errors.New("authorization failed"), context)
				return
			}

			if !status {
				ErrHandlerWithContext(errors.New("authorization failed"), context)
				return
			}

			//userID := context.Request.Header.Get("x-id")
			//if userID == "" {
			//	log.Printf("[%s] %s - error: userID is not provided\n", context.Request.Method, context.Request.RequestURI)
			//	ErrHandlerWithContext(errors.New("users id is not provided"), context)
			//	return
			//}
		}

		//
		//ctx := r.Context()
		//ctx = context.WithValue(ctx, "id", userID)
		//
		//r = r.WithContext(ctx)

		next(context)
	}
}

func ErrorHandler(err error) {
	fmt.Printf("error: %s", err.Error())
}

func ErrHandlerWithContext(err error, context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	ErrorHandler(err)
}

func CreateDbPgConnection() (*sql.DB, error) {
	source := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		ConfigV.HostPgSQL, ConfigV.PortPgSQL, ConfigV.UserPgSQL, ConfigV.PasswordPgSQL, ConfigV.DatabasePgSQL, ConfigV.SslModePgSQL,
	)
	dbConnection, err := sql.Open(ConfigV.DriverNamePgSQL, source)
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

func ExecuteSelectOneDb(object []any, query string, args ...any) error {
	dbConnection, err := CreateDbPgConnection()
	if err != nil {
		return err
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			return
		}
	}(dbConnection)

	rows, err := dbConnection.Query(query, args...)
	if err != nil {
		return err
	}

	rows.Next()
	err = rows.Scan(object...)
	if err != nil {
		return err
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func ExecuteSelectManyDb(objects *[]string, query string, args ...any) error {
	dbConnection, err := CreateDbPgConnection()
	if err != nil {
		return err
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			return
		}
	}(dbConnection)

	rows, err := dbConnection.Query(query, args...)
	if err != nil {
		return err
	}

	usernames := make([]string, 0)
	for rows.Next() {
		var obj string
		err = rows.Scan(&obj)
		if err != nil {
			return err
		}

		usernames = append(usernames, obj)
	}
	fmt.Println(usernames)
	*objects = usernames

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func ExecuteRowsDb(query string, args ...any) (*sql.Rows, error) {
	dbConnection, err := CreateDbPgConnection()
	if err != nil {
		return nil, err
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			return
		}
	}(dbConnection)

	rows, err := dbConnection.Query(query, args...)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func ExecuteInsertOrDeleteDb(query string, args ...any) error {
	dbConnection, err := CreateDbPgConnection()
	if err != nil {
		return err
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		return err
	}
	defer func(dbTransaction *sql.Tx) {
		_ = dbTransaction.Rollback()
	}(dbTransaction)

	_, err = dbTransaction.Exec(query, args...)
	if err != nil {
		return err
	}

	err = dbTransaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetRequestBodyData(context *gin.Context) (map[string]any, error) {
	var data = make(map[string]any)
	err := context.ShouldBindJSON(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ConfigConstructor() *ConfigS {
	config_ := &ConfigS{
		// Gin
		HostGin:  "127.0.0.1",
		PortGin:  8080,
		DebugGin: true,

		// Postgresql
		DriverNamePgSQL: "postgres",
		HostPgSQL:       "127.0.0.1",
		PortPgSQL:       5432,
		UserPgSQL:       "pgs_usr",
		PasswordPgSQL:   "12345Qwerty!",
		DatabasePgSQL:   "pgs_db",
		SslModePgSQL:    "disable",

		// Additional
		HashCost:             14,
		HashSalt:             "Qwerty!12345",
		TokenAccessLifetime:  time.Minute * 10,
		TokenRefreshLifetime: time.Minute * 60 * 24,
	}
	ReadConfig(config_, "../.env")
	return config_
}

func ReadConfig(config *ConfigS, filename string) *ConfigS {
	// TODO NEED read and parse from .ENV file
	//return config{DebugGin: c.DebugGin, HostGin: c.HostGin, PORT: c.PORT}
	_ = filename
	return config
}

func HashPassword(passwords ...string) (string, error) {
	password := ""
	for _, pwd := range passwords {
		password += pwd
	}
	if len(password) < 4 {
		return "", errors.New("password very simple")
	} else if len(password) > 41 {
		password = password[:40]
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), ConfigV.HashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func CreateToken(user *User) (string, string, error) {
	// TODO NEED check token date is not expired
	a := user.Username
	b := user.DatetimeJoined.String()
	c := ConfigV.HashSalt

	access, err := HashPassword(a, b, c)
	if err != nil {
		return "", "", err
	}

	refresh, err := HashPassword(c, b, a)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func CheckAccessToken(accessToken string) (bool, error) {
	// find users in db
	var username string
	var datetimeCreated time.Time
	err := ExecuteSelectOneDb([]any{&username, &datetimeCreated},
		"select username, datetime_created from tokens where access=$1;", accessToken)
	if err != nil {
		return false, err
	}
	if username == "" {
		return false, errors.New("token not found")
	}

	if time.Now().Sub(datetimeCreated).Minutes() > ConfigV.TokenAccessLifetime.Minutes() {
		return false, errors.New("token lifetime expired")
	}

	return true, nil
}

func CreateUsersDatabase() {
	query := "create table users (id serial not null unique, username varchar(255) not null unique, password varchar(255) not null, datetime_joined timestamp not null default now(), is_active bool default 'true');"
	err := ExecuteInsertOrDeleteDb(query)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Println("database 'users' successfully added")
}

func CreateTokensDatabase() {
	query := "create table tokens (id serial not null unique, username varchar(255) not null unique, access varchar(255) not null, refresh varchar(255) not null, datetime_created timestamp not null default now());"
	err := ExecuteInsertOrDeleteDb(query)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Println("database 'tokens' successfully added")
}

func CheckCaptchaHandler(context *gin.Context) {
	// get data from request
	//_, err := GetRequestBodyData(context)
	//if err != nil {
	//	ErrHandlerWithContext(err, context)
	//	return
	//}

	context.JSON(http.StatusOK, map[string]string{"response": "success"})
}

func RegisterUserHandler(context *gin.Context) {
	// get and check param
	username := context.PostForm("username")
	if username == "" {
		ErrHandlerWithContext(errors.New("username incorrect"), context)
		return
	}

	// get and check param
	password := context.PostForm("password")
	if password == "" {
		ErrHandlerWithContext(errors.New("password incorrect"), context)
		return
	}

	// hash password
	password, err := HashPassword(password)
	if password == "" {
		ErrHandlerWithContext(err, context)
		return
	}

	// write to db
	err = ExecuteInsertOrDeleteDb("insert into users (username, password) values ($1, $2);", username, password)
	if err != nil {
		ErrHandlerWithContext(err, context)
		return
	}

	time.Sleep(time.Millisecond * 2000)

	// return
	context.JSON(http.StatusCreated, map[string]string{"response": "success"})
}

func LoginUserHandler(context *gin.Context) {
	// get and check username
	username := context.PostForm("username")
	if username == "" {
		ErrHandlerWithContext(errors.New("username incorrect"), context)
		return
	}

	// get and check password
	password := context.PostForm("password")
	if password == "" {
		ErrHandlerWithContext(errors.New("password incorrect"), context)
		return
	}

	// find users in db
	user := User{}
	err := ExecuteSelectOneDb([]any{&user.Id, &user.Username, &user.Password},
		"select id, username, password from users where username=$1;", username)
	if err != nil {
		ErrHandlerWithContext(err, context)
		return
	}
	if user.Username == "" {
		ErrHandlerWithContext(errors.New("users not found"), context)
		return
	}

	// check hash password
	success := CheckPasswordHash(password, user.Password)
	if success != true {
		ErrHandlerWithContext(errors.New("password didn't match"), context)
		return
	}

	// generate tokens
	access, refresh, err := CreateToken(&user)
	if success != true {
		ErrHandlerWithContext(err, context)
		return
	}

	// write tokens to db
	err = ExecuteInsertOrDeleteDb("insert into tokens (username, access, refresh) values ($1, $2, $3);", user.Username, access, refresh)
	if err != nil {
		err = ExecuteInsertOrDeleteDb("update tokens set access=$1, refresh=$2 where username = $3;", access, refresh, user.Username)
		if err != nil {
			ErrHandlerWithContext(err, context)
			return
		}
	}

	// return
	context.JSON(http.StatusOK, map[string]map[string]string{"response": {"access": access, "refresh": refresh}})
}

func GetAllUsersHandler(context *gin.Context) {
	usernames := make([]string, 0)
	err := ExecuteSelectManyDb(&usernames, "select username from users;")
	if err != nil {
		ErrHandlerWithContext(err, context)
		return
	}

	context.JSON(http.StatusOK, map[string]map[string]any{"response": {"list": usernames}})
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	CreateUsersDatabase()
	CreateTokensDatabase()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
