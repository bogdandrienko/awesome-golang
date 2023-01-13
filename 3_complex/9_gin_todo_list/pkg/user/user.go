package user

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-study/3_complex/9_gin_todo_list/pkg/app"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

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

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func HashPassword(passwords ...string) (string, error) {
	password := ""
	for _, pwd := range passwords {
		password += pwd
	}
	if len(password) < 4 {
		return "", errors.New("password very simple")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password[:40]), app.Config.HashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func CreateToken(user *User) (string, string, error) {
	// TODO NEED check token date is not expired
	a := user.Username
	b := user.DatetimeJoined.String()
	c := app.Config.HashSalt

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

func CheckExpired() bool {
	// TODO NEED check token date is not expired
	return true
}

func CreateUsersDatabase() {
	query := "create table users (id serial not null unique, username varchar(255) not null unique, password varchar(255) not null, datetime_joined timestamp not null default now(), is_active bool default 'true');"
	err := app.Utils.ExecuteInsertDb(query)
	if err != nil {
		app.Utils.ErrorHandler(err)
		return
	}

	fmt.Println("database 'users' successfully added")
}

func CreateTokensDatabase() {
	query := "create table tokens (id serial not null unique, username varchar(255) not null unique, access varchar(255) not null, refresh varchar(255) not null, datetime_created timestamp not null default now());"
	err := app.Utils.ExecuteInsertDb(query)
	if err != nil {
		app.Utils.ErrorHandler(err)
		return
	}

	fmt.Println("database 'tokens' successfully added")
}

func RegisterUserHandler(c *gin.Context) {
	// get data from request
	data, err := app.Utils.GetRequestBodyData(c)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	// check username
	username := data["username"].(string)
	if username == "" {
		app.Utils.ErrHandlerWithContext(errors.New("username incorrect"), c)
		return
	}

	// check password
	password := data["password"].(string)
	if password == "" {
		app.Utils.ErrHandlerWithContext(errors.New("password incorrect"), c)
		return
	}

	// hash password
	password, err = HashPassword(password)
	if password == "" {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	// write to db
	err = app.Utils.ExecuteInsertDb("insert into users (username, password) values ($1, $2);", username, password)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	// return
	c.Status(http.StatusCreated)
}

func LoginUserHandler(c *gin.Context) {
	// get data from request
	data, err := app.Utils.GetRequestBodyData(c)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	// check username
	username := data["username"].(string)
	if username == "" {
		app.Utils.ErrHandlerWithContext(errors.New("username incorrect"), c)
		return
	}

	// check password
	password := data["password"].(string)
	if password == "" {
		app.Utils.ErrHandlerWithContext(errors.New("password incorrect"), c)
		return
	}

	// find user in db
	user := User{}
	err = app.Utils.ExecuteSelectOneDb([]any{&user.Id, &user.Username, &user.Password},
		"select id, username, password from users where username=$1;", username)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	if user.Username == "" {
		app.Utils.ErrHandlerWithContext(errors.New("user not found"), c)
		return
	}

	// check hash password
	success := CheckPasswordHash(password, user.Password)
	if success != true {
		app.Utils.ErrHandlerWithContext(errors.New("password didn't match"), c)
		return
	}

	// generate tokens
	access, refresh, err := CreateToken(&user)
	if success != true {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	response := map[string]string{"access": access, "refresh": refresh}

	// write tokens to db
	err = app.Utils.ExecuteInsertDb("insert into tokens (username, access, refresh) values ($1, $2, $3);", user.Username, access, refresh)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	// return
	c.JSON(http.StatusOK, response)
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
