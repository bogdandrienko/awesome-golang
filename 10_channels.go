package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func channels1() {
	var msg chan string
	fmt.Println(msg)

	//msg <- "Chanel write" // TODO deadlock! because chanel is close

	msg = make(chan string)
	fmt.Println(msg)

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "Chanel write"
	}()
	//

	value := <-msg
	fmt.Println(value)
}

func channels2() {
	msg := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "Chanel write 1"
		msg <- "Chanel write 2"
		msg <- "Chanel write 3"
	}()
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

func channels3() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Channel 1. after 200ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			message2 <- "Channel 2. after 1000ms"
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		default:
		}
	}
}

func channels4() {
	t := time.Now()

	wg := &sync.WaitGroup{}
	users := make(chan User2, 1000)
	go generateUsers2(1000, users)

	for user := range users {
		wg.Add(1)
		go saveUserInfo3(user, wg)
	}

	wg.Wait()

	fmt.Println("noConcurrent3 Elapsed time: ", time.Since(t).String())
}

type User2 struct {
	id    int
	email string
	logs  []logItem2
}

func (u User2) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email : %s\n Activity Log:\n", u.id, u.email)
	for idx, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", idx, item.action, item.timestamp)
	}
	return out
}

func saveUserInfo3(user User2, wg *sync.WaitGroup) error {
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

var actions2 = []string{
	"action 1",
	"action 2",
	"action 3",
	"action 4",
	"action 5",
}

type logItem2 struct {
	action    string
	timestamp time.Time
}

func generateUsers2(count int, users chan User2) {
	for i := 0; i < count; i += 1 {
		users <- User2{
			id:    i + 1,
			email: fmt.Sprintf("example%d@gmail.com", i+1),
			logs: []logItem2{
				{action: actions2[rand.Intn(4)], timestamp: time.Now()},
				{action: actions2[rand.Intn(4)], timestamp: time.Now()},
				{action: actions2[rand.Intn(4)], timestamp: time.Now()},
			},
		}
		time.Sleep(time.Millisecond * 10)
	}

	close(users)
}

func main() {
	channels4()
}
