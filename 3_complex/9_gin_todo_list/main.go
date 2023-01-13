package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	app "golang-study/3_complex/9_gin_todo_list/pkg/app"
	user "golang-study/3_complex/9_gin_todo_list/pkg/user"
	"log"
	"net/http"
	"strconv"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
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

// Create
func createBookHandler(c *gin.Context) {
	var book Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	dbConnection, err := app.Utils.CreateDbPgConnection()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			app.Utils.ErrorHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		_ = dbTransaction.Rollback()
	}(dbTransaction)

	_, err = dbTransaction.Exec("insert into books (title, author) values ($1, $2);", book.Title, book.Author)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	c.Status(http.StatusCreated)
}

// Read
func getBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	dbConnection, err := app.Utils.CreateDbPgConnection()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			app.Utils.ErrorHandler(err)
			return
		}
	}(dbConnection)

	row, err := dbConnection.Query("select * from books where id = $1", id)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	book := Book{}
	row.Next()
	err = row.Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	err = row.Err()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	c.JSON(http.StatusOK, book)
}

// Read all
func getBooksHandler(c *gin.Context) {
	dbConnection, err := app.Utils.CreateDbPgConnection()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			app.Utils.ErrorHandler(err)
		}
	}(dbConnection)

	rows, err := dbConnection.Query("select * from books order by id asc;")
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			app.Utils.ErrHandlerWithContext(err, c)
			return
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	c.JSON(http.StatusOK, books)
}

// Update
func updateBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	var book Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	dbConnection, err := app.Utils.CreateDbPgConnection()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			app.Utils.ErrorHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		_ = dbTransaction.Rollback()
	}(dbTransaction)

	_, err = dbTransaction.Exec("update books set title=$1, author=$2 where id = $3;", book.Title, book.Author, id)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	c.Status(http.StatusOK)
}

// Delete
func deleteBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	dbConnection, err := app.Utils.CreateDbPgConnection()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			app.Utils.ErrorHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		_ = dbTransaction.Rollback()
	}(dbTransaction)

	_, err = dbTransaction.Exec("delete from books where id = $1;", id)
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		app.Utils.ErrHandlerWithContext(err, c)
		return
	}

	c.Status(http.StatusOK)
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func run() {
	err := app.App{}.NewApp([]app.Route{
		// user(s)
		{"POST", "/users", user.RegisterUserHandler, false},    // Create   | curl -v -X POST 127.0.0.1:8080/users -d '{"username":"admin", "password": "admin"}'
		{"POST", "/users/login", user.LoginUserHandler, false}, // Create   | curl -v -X POST 127.0.0.1:8080/users/login -d '{"username":"admin", "password": "admin"}'

		// book(s)
		{"POST", "/books", createBookHandler, false},      // Create   | curl -v -X POST 127.0.0.1:8080/books -d '{"title":"Amon Ra","author":"V.Pelevin"}'
		{"GET", "/books/:id", getBookHandler, false},      // Read     | curl -v -X GET 127.0.0.1:8080/books/1
		{"GET", "/books", getBooksHandler, false},         // Read all | curl -v -X GET 127.0.0.1:8080/books
		{"PUT", "/books/:id", updateBookHandler, false},   // Update   | curl -v -X PUT 127.0.0.1:8080/books/1 -d '{"title":"War and peace","author":"N.Tolstoy"}'
		{"DELETE", "/books/:id", deleteBookHandler, true}, // Delete   | curl -v -H 'x-id:1' -X DELETE 127.0.0.1:8080/books/1
	})
	if err != nil {
		app.Utils.ErrorHandler(err)
		log.Fatal(err)
		return
	}
}

func main() {
	run()

	//user.CreateUsersDatabase()
	//user.CreateTokensDatabase()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
