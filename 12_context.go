package main

import (
	"context"
	"fmt"
	"time"
)

func context1() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*1)
	parse(ctx)
}

func parse(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}

func main() {
	context1()
}
