package main

import "fmt"

func main() {
	message := "Hello!"
	fmt.Println(message)
	printMessage1(message)
	fmt.Println(message)
	printMessage2(&message)
	fmt.Println(message)

	number := 5
	var p *int

	p = &number
	fmt.Println(p)
	fmt.Println(&number)

	*p = 10
	fmt.Println(number)

}

func printMessage1(message string) {
	message += " World!"
	fmt.Println(message)
}

func printMessage2(message *string) {
	*message += " World!"
	fmt.Println(*message)
}
