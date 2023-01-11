package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////
func readOneSimple() {
	// todo open file
	jsonFile, err := os.Open("3_complex/1_read_json_write_txt/from/todo.json")
	if err != nil {
		fmt.Println(err)
	}

	// todo after close file
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(jsonFile)

	// todo read file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// todo deserialize bytes
	var todo Todo
	if json.Unmarshal(byteValue, &todo) != nil {
		log.Fatal(err)
		return
	}

	// todo print data
	fmt.Println(todo, reflect.TypeOf(todo))
}

func readOneComplex(isPrint bool) (Todo, error) {
	jsonFile, err := os.Open("3_complex/1_read_json_write_txt/from/todo.json")
	if err != nil {
		return Todo{}, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(jsonFile)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Todo{}, err
	}
	var todo Todo
	if json.Unmarshal(byteValue, &todo) != nil {
		log.Fatal(err)
		return Todo{}, err
	}

	if isPrint {
		fmt.Println(todo, reflect.TypeOf(todo))
	}
	return todo, nil
}

func readOneComplexStart() {
	todo, err := readOneComplex(false)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(todo, reflect.TypeOf(todo))

	writeTxtSimple(todo)

	err = writeTxtComplex(todo)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func readManySimple() {
	// todo open file
	jsonFile, err := os.Open("3_complex/1_read_json_write_txt/from/todos.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// todo after close file
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(jsonFile)

	// todo read file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// todo deserialize bytes
	var todos []Todo
	if json.Unmarshal(byteValue, &todos) != nil {
		log.Fatal(err)
		return
	}

	// todo print data
	for index, todo := range todos {
		fmt.Println(fmt.Sprintf("#%d", index), todo, reflect.TypeOf(todo))
	}
}

func readManyComplex(isPrint bool) ([]Todo, error) {
	jsonFile, err := os.Open("3_complex/1_read_json_write_txt/from/todos.json")
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(jsonFile)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	if json.Unmarshal(byteValue, &todos) != nil {
		log.Fatal(err)
		return nil, err
	}

	if isPrint {
		for index, todo := range todos {
			fmt.Println(fmt.Sprintf("#%d", index), todo, reflect.TypeOf(todo))
		}
	}
	return todos, err
}

func readManyComplexStart() {
	todos, err := readManyComplex(false)
	if err != nil {
		log.Fatal(err)
		return
	}
	for index, todo := range todos {
		fmt.Println(fmt.Sprintf("#%d", index), todo, reflect.TypeOf(todo))
	}

	for _, todo := range todos {
		writeTxtSimple(todo)

		err = writeTxtComplex(todo)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	err = writeSliceTxtComplex(todos)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func writeTxtSimple(todo Todo) {
	var filename string
	filename = fmt.Sprintf("3_complex/1_read_json_write_txt/to/%d.txt", todo.Id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)

	_, err = file.WriteString(todo.GetString())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeTxtComplex(todo Todo) error {
	file, err := os.OpenFile(fmt.Sprintf("3_complex/1_read_json_write_txt/to/%d.txt", todo.Id), os.O_RDWR|os.O_CREATE, 0644)
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

	_, err = file.WriteString(todo.GetString())
	if err != nil {
		return err
	}

	return nil
}

func writeSliceTxtComplex(todos []Todo) error {
	file, err := os.OpenFile("3_complex/1_read_json_write_txt/to/0.txt", os.O_RDWR|os.O_CREATE, 0644)
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

	for _, todo := range todos {
		_, err = file.WriteString(todo.GetString())
		if err != nil {
			return err
		}
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
	readOneSimple()
	readOneComplexStart()
	readManySimple()
	readManyComplexStart()
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
