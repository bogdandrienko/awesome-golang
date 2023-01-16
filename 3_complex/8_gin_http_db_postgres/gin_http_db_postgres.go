package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

const (
	DEBUG bool   = true
	HOST  string = "0.0.0.0"
	PORT  int    = 8080
)

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func errContextHandler(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	fmt.Printf("Error: [%s]", err.Error())
}

func errHandler(err error) {
	fmt.Printf("Error: [%s]", err.Error())
}

func createDbConnection() (*sql.DB, error) {
	dbConnection, err := sql.Open("postgres", "host=127.0.0.1 port=5432 users=pgs_usr password=12345Qwerty! dbname=pgs_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

// Create
func createBookHandler(c *gin.Context) {
	var book Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	dbConnection, err := createDbConnection()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		err = dbTransaction.Rollback()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbTransaction)

	_, err = dbTransaction.Exec("insert into books (title, author) values ($1, $2);", book.Title, book.Author)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		errContextHandler(err, c)
		return
	}

	c.Status(http.StatusCreated)
}

// Read
func getBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errContextHandler(err, c)
		return
	}

	dbConnection, err := createDbConnection()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbConnection)

	row, err := dbConnection.Query("select * from books where id = $1", id)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	book := Book{}
	row.Next()
	err = row.Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	err = row.Err()
	if err != nil {
		errContextHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, book)
}

// Read all
func getBooksHandler(c *gin.Context) {
	dbConnection, err := createDbConnection()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbConnection)

	rows, err := dbConnection.Query("select * from books order by id asc")
	if err != nil {
		errContextHandler(err, c)
		return
	}

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			errContextHandler(err, c)
			return
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		errContextHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, books)
}

// Update
func updateBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errContextHandler(err, c)
		return
	}

	var book Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	dbConnection, err := createDbConnection()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		err := dbTransaction.Rollback()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbTransaction)

	_, err = dbTransaction.Exec("update books set title=$1, author=$2 where id = $3;", book.Title, book.Author, id)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		errContextHandler(err, c)
		return
	}

	c.Status(http.StatusOK)
}

// Delete
func deleteBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errContextHandler(err, c)
		return
	}

	dbConnection, err := createDbConnection()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbConnection *sql.DB) {
		err = dbConnection.Close()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbConnection)

	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		errContextHandler(err, c)
		return
	}
	defer func(dbTransaction *sql.Tx) {
		err = dbTransaction.Rollback()
		if err != nil {
			errHandler(err)
			return
		}
	}(dbTransaction)

	_, err = dbTransaction.Exec("delete from books where id = $1;", id)
	if err != nil {
		errContextHandler(err, c)
		return
	}

	err = dbTransaction.Commit()
	if err != nil {
		errContextHandler(err, c)
		return
	}

	c.Status(http.StatusOK)
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	if DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// Create
	r.POST("/books", createBookHandler)

	// Read
	r.GET("/books/:id", getBookHandler)

	// Read all
	r.GET("/books", getBooksHandler)

	// Update
	r.PUT("/books/:id", updateBookHandler)

	// Delete
	r.DELETE("/books/:id", deleteBookHandler)

	err := r.Run(fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Fatal(err)
		return
	}
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// Create
// curl -v -X POST 127.0.0.1:8080/books -d '{"title":"Amon Ra","author":"V.Pelevin"}'

// Read
// curl -v -X GET 127.0.0.1:8080/books/1

// Read all
// curl -v -X GET 127.0.0.1:8080/books

// Update
// curl -v -X PUT 127.0.0.1:8080/books/1 -d '{"title":"War and peace","author":"N.Tolstoy"}'

// Delete
// curl -v -X DELETE 127.0.0.1:8080/books/1

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
