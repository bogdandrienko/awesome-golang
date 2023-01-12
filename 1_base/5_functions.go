package main

import (
	"errors"
	"fmt"
	"reflect"
)

var msg string

func init() {
	msg = "from init"
	fmt.Println("World!")
}

func printHello() {
	fmt.Println("Hello")
}

func printMessage(message string) {
	fmt.Println(message)
}

func printAny(variables ...any) {
	for _, value := range variables {
		fmt.Println(value, reflect.TypeOf(value))
	}
}

func getSum(val1 int, val2 int) int {
	result := val1 + val2
	return result
}

func checkAdult(age int) (string, bool) {
	if age < 18 {
		return "Входите", true
	}
	return "Стойте", false
}

func findMin(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("len of param 'numbers' is 0")
	}
	min := numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}

	return min, nil
}

func anonymousFunction() {
	func() {
		fmt.Println("I'm a anonymous")
	}()
}

func incrementValue() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	printHello()

	printMessage("Hello World!")

	result := getSum(12, 13)
	message := fmt.Sprintf("Result is %d", result)
	printMessage(message)

	mes, error1 := checkAdult(17)
	printAny(mes, error1)

	minVal, error2 := findMin()
	if error2 != nil {
		fmt.Println(error2)
	} else {
		fmt.Println(minVal, reflect.TypeOf(minVal))
	}

	minVal2, error3 := findMin(1, 7, 3, 9, -12)
	if error3 != nil {
		fmt.Println(error3)
	} else {
		fmt.Println(minVal2, reflect.TypeOf(minVal2))
	}

	anonymousFunction()

	funcType1 := incrementValue()
	res1 := funcType1()
	fmt.Println(res1)
	res2 := funcType1()
	fmt.Println(res2)
	res3 := funcType1()
	fmt.Println(res3)

	fmt.Println(msg)
}
