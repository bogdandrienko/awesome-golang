package app

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"C"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type App struct {
	temp bool
}

type config struct {
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
	HashCost int
	HashSalt string
}

type Route struct {
	Method       string
	RelativePath string
	Function     gin.HandlerFunc
	Auth         bool
}

type utils struct {
	temp bool
}

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (a App) NewApp(routes []Route) error {
	// todo Configure engine
	if Config.DebugGin {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// todo Create engine
	engine := gin.New()

	// todo Binding routes
	for _, route := range routes {
		switch route.Method {
		case "POST": // Create
			engine.POST(route.RelativePath, Utils.Middleware(route.Function, route))
		case "GET": // Read
			engine.GET(route.RelativePath, Utils.Middleware(route.Function, route))
		case "PUT": // Update
			engine.PUT(route.RelativePath, Utils.Middleware(route.Function, route))
		case "DELETE": // Delete
			engine.DELETE(route.RelativePath, Utils.Middleware(route.Function, route))
		default:
			Utils.ErrorHandler(errors.New("unknown method"))
		}
	}

	// todo Run engine
	return engine.Run(fmt.Sprintf("%s:%d", Config.HostGin, Config.PortGin))
}

func (c config) ReadConfig(filename string) *config {
	// TODO NEED read and parse from .ENV file
	//return config{DebugGin: c.DebugGin, HostGin: c.HostGin, PORT: c.PORT}
	_ = filename
	return &c
}

func (u utils) Middleware(next gin.HandlerFunc, route Route) gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Printf("[%s] %s\n", context.Request.Method, context.Request.RequestURI)

		if route.Auth {
			userID := context.Request.Header.Get("x-id")
			if userID == "" {
				log.Printf("[%s] %s - error: userID is not provided\n", context.Request.Method, context.Request.RequestURI)
				Utils.ErrHandlerWithContext(errors.New("user id is not provided"), context)
				return
			}
		}

		//
		//ctx := r.Context()
		//ctx = context.WithValue(ctx, "id", userID)
		//
		//r = r.WithContext(ctx)

		next(context)
	}
}

func (u utils) ErrorHandler(err error) {
	fmt.Printf("error: %s", err.Error())
}

func (u utils) ErrHandlerWithContext(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	u.ErrorHandler(err)
}

func (u utils) CreateDbPgConnection() (*sql.DB, error) {
	source := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Config.HostPgSQL, Config.PortPgSQL, Config.UserPgSQL, Config.PasswordPgSQL, Config.DatabasePgSQL, Config.SslModePgSQL,
	)
	dbConnection, err := sql.Open(Config.DriverNamePgSQL, source)
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

func (u utils) ExecuteSelectOneDb(object []any, query string, args ...any) error {
	dbConnection, err := Utils.CreateDbPgConnection()
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

func (u utils) ExecuteInsertDb(query string, args ...any) error {
	dbConnection, err := Utils.CreateDbPgConnection()
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

func (u utils) GetRequestBodyData(c *gin.Context) (map[string]any, error) {
	var data = make(map[string]any)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

var Config = configConstructor()
var Utils = utils{}

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func configConstructor() *config {
	c := &config{
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
		HashCost: 14,
		HashSalt: "Qwerty!12345",
	}
	c.ReadConfig("../.env")
	return c
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
