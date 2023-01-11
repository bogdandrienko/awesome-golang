package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		userID := r.Header.Get("x-id")
		if userID == "" {
			log.Printf("[%s] %s - error: userID is not provided\n", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userID)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userID is invalid\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("[%s] %s by UserID %s\n", r.Method, r.URL, userID)
		next(w, r)
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// todo read file
		jsonFile, err := os.OpenFile("3_complex/5_native_http_server/db.json", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}(jsonFile)

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var todos []Todo
		if json.Unmarshal(byteValue, &todos) != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// todo convert file to json
		resp, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		// todo json response
		_, err = w.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		// todo read file
		jsonFile, err := os.OpenFile("3_complex/5_native_http_server/db.json", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}(jsonFile)

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var todos []Todo
		if json.Unmarshal(byteValue, &todos) != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// todo read request
		reqBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		var todo Todo
		if err = json.Unmarshal(reqBytes, &todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// todo append todo
		todos = append(todos, todo)

		// todo convert file to json
		todoJson, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// todo write to file
		file, err := os.OpenFile("3_complex/5_native_http_server/db.json", os.O_RDWR|os.O_CREATE, 0644)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}(file)
		_, err = file.Write(todoJson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
	case http.MethodPut:
	//
	//
	case http.MethodDelete:
	//
	//
	default:
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//var (
//	todos = []Todo{{1, 1, "Wash a cat", true}}
//)

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {

	const (
		host string = "127.0.0.1"
		port int    = 8080
	)

	http.HandleFunc("/todos", authMiddleware(loggerMiddleware(handleTodos)))

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// curl -v -X GET -H 'x-id:1' 127.0.0.1:8080/todos
// curl -v -X POST -H 'x-id:1' 127.0.0.1:8080/todos -d '{"userId":2,"id":2,"title":"11111111111111","completed":true}'

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
