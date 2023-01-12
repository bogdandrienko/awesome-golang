package main

import (
	"context"
	"fmt"
	"time"
)

// context.Background - on the top level
// context.TODO - unknown task
// context.Value - use small and set not needed values

func context1() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()
	parse1(ctx)
}

func parse1(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		default:

		}
	}
}

func context2() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	ctx = context.WithValue(ctx, "id", 1)

	parse2(ctx)
}

func parse2(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Println(id.(int))
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		default:

		}
	}
}

func main() {
	//context1()
	context2()
}
