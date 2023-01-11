package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func getDataOneSimple() {
	var url string
	url = "https://jsonplaceholder.typicode.com/todos/1"

	// todo get json data from api
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)

	// todo read bytes from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// todo fast write json from web to json-file
	file, err := os.OpenFile("3_complex/2_get_http_write_json/to/_1.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)
	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(resp.StatusCode, string(body))
	fmt.Println(todo.GetString())

	example, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err = os.OpenFile("3_complex/2_get_http_write_json/to/__1.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)

	_, err = file.Write(example)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getDataOneComplex() (Todo, error) {
	resp, err := http.DefaultClient.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return Todo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func getDataOneComplexStart() {
	todo, err := getDataOneComplex()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(todo.GetString())

	err = writeTodoToJson(todo)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getDataManyComplex() ([]Todo, error) {
	resp, err := http.DefaultClient.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return []Todo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Todo{}, err
	}

	var todos []Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		return []Todo{}, err
	}

	return todos, nil
}

func getDataManyComplexStart() {
	todos, err := getDataManyComplex()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(todos)

	err = writeTodosToJson(todos)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, todo := range todos {
		err = writeTodoToJson(todo)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func getDataSimple() {
	var url string
	url = "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.OpenFile("3_complex/2_get_http_write_json/to/0.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)
	_, err = file.Write(body)
	if err != nil {
		return
	}

	fmt.Println(resp.StatusCode, string(body))

	var todos []Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, todo := range todos {
		fmt.Println(todo.GetString())
	}

	example, _ := json.Marshal(todos)
	file, err = os.OpenFile("3_complex/2_get_http_write_json/to/1.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)
	_, err = file.Write(example)
	if err != nil {
		return
	}
}

func writeTodoToJson(todo Todo) error {
	todoJson, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("3_complex/2_get_http_write_json/to/%d.json", todo.Id),
		os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)

	_, err = file.Write(todoJson)
	if err != nil {
		return err
	}

	return nil
}

func writeTodosToJson(todos []Todo) error {
	todosJson, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	file, err := os.OpenFile("3_complex/2_get_http_write_json/to/0.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)

	_, err = file.Write(todosJson)
	if err != nil {
		return err
	}

	return nil
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

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (todo Todo) GetString() string {
	return fmt.Sprintf("id: %d | userId: %d | title: %s | completed %t\n",
		todo.Id, todo.UserId, todo.Title, todo.Completed)
}

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	//getDataOneSimple()
	//getDataOneComplexStart()
	getDataManyComplexStart()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
