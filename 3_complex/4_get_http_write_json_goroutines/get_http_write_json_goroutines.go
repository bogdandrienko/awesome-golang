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
	"sync"
	"time"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func getData(url string) (Todo, error) {
	resp, err := http.DefaultClient.Get(url)
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

func writeTodoToJson(todo Todo) error {
	todoJson, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("3_complex/4_get_http_write_json_goroutines/to/%d.json", todo.Id),
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

func oneTodoAction(url string) {
	todo, err := getData(url)
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

	//time.Sleep(time.Millisecond * 100)
}

func oneTodoActionGo(ulr string, wg *sync.WaitGroup) {
	todo, err := getData(ulr)
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

	//time.Sleep(time.Millisecond * 100)

	wg.Done()
}

func oneTodoActionGoPool(id int, jobs <-chan string, results chan<- int) {
	for url := range jobs {
		todo, err := getData(url)
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

		//time.Sleep(time.Millisecond * 100)

		results <- 1
	}
}

func runSync(urls []string) {
	for _, url := range urls {
		oneTodoAction(url)
	}
}

func runGoWaitGroup(urls []string) {
	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go oneTodoActionGo(url, wg)
	}

	wg.Wait()
}

func runGoThreadPool(urls []string, workerCount int) {
	jobs := make(chan string, len(urls))
	results := make(chan int, len(urls))

	for i := 0; i < workerCount; i += 1 {
		go oneTodoActionGoPool(i+1, jobs, results)
	}

	for i := 0; i < len(urls); i += 1 {
		jobs <- urls[i]
	}
	close(jobs)

	for i := 0; i < len(urls); i += 1 {
		<-results
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

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (todo Todo) GetString() string {
	return fmt.Sprintf("id: %d | userId: %d | title: %s | completed %t",
		todo.Id, todo.UserId, todo.Title, todo.Completed)
}

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	t := time.Now()

	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/11",
		"https://jsonplaceholder.typicode.com/todos/12",
		"https://jsonplaceholder.typicode.com/todos/13",
		"https://jsonplaceholder.typicode.com/todos/14",
		"https://jsonplaceholder.typicode.com/todos/15",
		"https://jsonplaceholder.typicode.com/todos/21",
		"https://jsonplaceholder.typicode.com/todos/22",
		"https://jsonplaceholder.typicode.com/todos/23",
		"https://jsonplaceholder.typicode.com/todos/24",
		"https://jsonplaceholder.typicode.com/todos/25",
	}

	//runSync(urls) // 1.510791851s
	//runGoWaitGroup(urls) // 100.976268ms
	runGoThreadPool(urls, 4*2+1) // 201.224126ms

	fmt.Printf("ELAPSED TIME: " + time.Since(t).String())
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
