package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/bogdandrienko/gin-rest-api-v3/pkg/app"
	"github.com/bogdandrienko/gin-rest-api-v3/pkg/tasks"
	_ "github.com/lib/pq"
	"log"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO global /////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func run() {
	err := app.NewApp([]app.RouteS{
		// users(s)
		{"GET", "/api/users/check_captcha", app.CheckCaptchaHandler, false}, // Create   | curl -v -X POST 127.0.0.1:8080/users -d '{"username":"admin", "password": "admin"}'
		{"POST", "/api/users/register", app.RegisterUserHandler, false},     // Create   | curl -v -X POST 127.0.0.1:8080/users -d '{"username":"admin", "password": "admin"}'
		{"POST", "/api/users/login", app.LoginUserHandler, false},           // Create   | curl -v -X POST 127.0.0.1:8080/users/login -d '{"username":"admin", "password": "admin"}'
		{"GET", "/api/users/", app.GetAllUsersHandler, false},               // Create   | curl -v -X POST 127.0.0.1:8080/users/login -d '{"username":"admin", "password": "admin"}'

		// tasks(s)
		{"POST", "/api/tasks", tasks.CreateTaskHandler, true},       // Create   | curl -v -X POST 127.0.0.1:8080/tasks -d '{"title":"Amon Ra","author":"V.Pelevin"}'
		{"GET", "/api/tasks/:id", tasks.ReadTaskHandler, true},      // Read     | curl -v -X GET 127.0.0.1:8080/tasks/1
		{"GET", "/api/tasks", tasks.ReadTasksHandler, true},         // Read all | curl -v -X GET 127.0.0.1:8080/tasks
		{"PUT", "/api/tasks/:id", tasks.UpdateTaskHandler, true},    // Update   | curl -v -X PUT 127.0.0.1:8080/tasks/1 -d '{"title":"War and peace","author":"N.Tolstoy"}'
		{"DELETE", "/api/tasks/:id", tasks.DeleteTaskHandler, true}, // Delete   | curl -v -H 'x-id:1' -X DELETE 127.0.0.1:8080/tasks/1
	})
	if err != nil {
		app.ErrorHandler(err)
		log.Fatal(err)
		return
	}
}

func main() {
	// tsx home add data
	// tsx add fields
	// tsx search, filter... fields
	// gin response html
	// gin response static
	// tsx home add data
	// add file field
	// add redis-cache

	run()

	//app.CreateUsersDatabase()
	//app.CreateTokensDatabase()
	//tasks.CreateTasksDatabase()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO extra //////////////////////////////////////////////////////////////////////////////////////////////////////////

//
