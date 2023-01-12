package main

import (
	"fmt"
	"sync"
	"time"
)

func race1() {
	counter := 0
	for i := 0; i < 1000; i += 1 {
		go func() {
			counter += 1
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Println(counter)

	// go run -race main.gp // TODO detect race condition
}

func race2() {
	c := counter{count: 0, mu: new(sync.Mutex)}
	for i := 0; i < 1000; i += 1 {
		go func() {
			c.inc()
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Println(c.value())

	// go run -race main.gp // TODO detect race condition
}

type counter struct {
	count int
	mu    *sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.count += 1
	c.mu.Unlock()
}

func (c *counter) value() int {
	//c.mu.Lock()
	//defer c.mu.Unlock()
	//return c.count

	c.mu.Lock()
	value := c.count
	c.mu.Unlock()
	return value
}

func main() {
	//race1()
	race2()
}
