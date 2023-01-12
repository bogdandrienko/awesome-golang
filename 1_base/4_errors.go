package main

import "fmt"

func main() {
	// TODO panic
	//row := make([]int, 5)
	//row[10] = 12 // panic: runtime error: index out of range [10] with length 5
	//fmt.Println(row)

	//panic("raise panic") // panic: raise panic

	// TODO defer
	//defer printMessage3() // last call +50 ms for program work
	//fmt.Println("print message before")

	// TODO recover
	defer handlePanic()
	panic("custom raise panic") // panic: raise panic
}

func printMessage3() {
	fmt.Println("print message after")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("handlePanic success worked")
}
