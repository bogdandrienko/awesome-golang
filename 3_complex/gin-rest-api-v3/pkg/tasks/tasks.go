package tasks

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"errors"
	"fmt"
	"github.com/bogdandrienko/gin-rest-api-v3/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
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

func CreateTasksDatabase() {
	err := app.ExecuteInsertOrDeleteDb("create table tasks (id serial not null unique, title varchar(255) not null unique);")
	if err != nil {
		app.ErrorHandler(err)
		return
	}

	fmt.Println("database 'tasks' successfully added")
}

func CreateTaskHandler(context *gin.Context) {
	// get and check param
	title := context.PostForm("title")
	if title == "" {
		app.ErrHandlerWithContext(errors.New("title incorrect"), context)
		return
	}

	// insert to db
	err := app.ExecuteInsertOrDeleteDb("insert into tasks (title) values ($1);", title)
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	context.JSON(http.StatusCreated, map[string]string{"response": "successfully created"})
}

func ReadTaskHandler(context *gin.Context) {
	// get id from url
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	// create object
	task := Task{}

	// select from db
	err = app.ExecuteSelectOneDb([]any{&task.Id, &task.Title}, "select id, title from tasks where id = $1", id)
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	context.JSON(http.StatusOK, map[string]any{"response": task})
}

func ReadTasksHandler(context *gin.Context) {
	// select from db
	rows, err := app.ExecuteRowsDb("select id, title from tasks order by id asc;")
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	// create objects
	tasks := make([]Task, 0)

	// fulling objects
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.Title)
		if err != nil {
			app.ErrHandlerWithContext(err, context)
			return
		}
		tasks = append(tasks, task)
	}

	context.JSON(http.StatusOK, map[string]map[string]any{"response": {"list": tasks}})
}

func UpdateTaskHandler(context *gin.Context) {
	// get id from url
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	// get and check param
	title := context.PostForm("title")
	if title == "" {
		app.ErrHandlerWithContext(errors.New("title incorrect"), context)
		return
	}

	// update into db
	err = app.ExecuteInsertOrDeleteDb("update tasks set title=$1 where id = $2;", title, id)
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"response": "successfully updated"})
}

func DeleteTaskHandler(context *gin.Context) {
	// get id from url
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	// delete from db
	err = app.ExecuteInsertOrDeleteDb("delete from tasks where id = $1;", id)
	if err != nil {
		app.ErrHandlerWithContext(err, context)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"response": "successfully deleted"})
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	CreateTasksDatabase()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
