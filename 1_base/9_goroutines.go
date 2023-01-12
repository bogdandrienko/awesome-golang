package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func concurrent1() {
	go fmt.Println("concurrent task")
	go fmt.Println("concurrent task")
	go fmt.Println("concurrent task")

	time.Sleep(time.Second * 1) // TODO need to await start goroutines
	fmt.Println("sync task")
}

func concurrent2() {
	go func() {
		time.Sleep(time.Second * 1)
		go fmt.Println("another concurrent task")
	}()
	go fmt.Println("concurrent task")
	go fmt.Println("concurrent task")
	go fmt.Println("concurrent task")

	time.Sleep(time.Second * 2)
	fmt.Println("sync task")
}

var actions = []string{
	"action 1",
	"action 2",
	"action 3",
	"action 4",
	"action 5",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User1 struct {
	id    int
	email string
	logs  []logItem
}

func (u User1) getActivityInfo() string {
	out := fmt.Sprintf("Id: %d | Email : %s\n Activity Log:\n", u.id, u.email)
	for idx, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", idx, item.action, item.timestamp)
	}

	return out
}

func noConcurrent3() { // 11.216206713s
	users := generateUsers(1000)

	t := time.Now()

	for _, user := range users {
		saveUserInfo1(user)
	}
	fmt.Println("noConcurrent3 Elapsed time: ", time.Since(t).String())
}

func Concurrent3() { // 74.788483ms
	users := generateUsers(1000)

	t := time.Now()

	wg := &sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		go saveUserInfo2(user, wg)
	}

	wg.Wait()

	fmt.Println("noConcurrent3 Elapsed time: ", time.Since(t).String())
}

func generateUsers(count int) []User1 {
	users := make([]User1, count)

	for i := 0; i < count; i += 1 {
		users[i] = User1{
			id:    i + 1,
			email: fmt.Sprintf("example%d@gmail.com", i+1),
			logs: []logItem{
				{action: actions[rand.Intn(4)], timestamp: time.Now()},
				{action: actions[rand.Intn(4)], timestamp: time.Now()},
				{action: actions[rand.Intn(4)], timestamp: time.Now()},
			},
		}
	}

	return users
}

func saveUserInfo1(user User1) error {
	fmt.Printf("Writing file for user: %d\n", user.id)

	time.Sleep(time.Millisecond * 10)

	filename := fmt.Sprintf("temp/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}
	return nil
}

func saveUserInfo2(user User1, wg *sync.WaitGroup) error {
	fmt.Printf("Writing file for user: %d\n", user.id)

	time.Sleep(time.Millisecond * 10)

	filename := fmt.Sprintf("temp/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}

	wg.Done()

	return nil
}

func main() {
	//concurrent1()
	//concurrent2()
	//noConcurrent3()
	Concurrent3()
}
