package main

import (
	"errors"
	"fmt"
	"log"
)

func onlyIf(age int) {
	if age > 18 {
		fmt.Println("Вы совершеннолетний")
	}
}

func ifElse(age int) {
	if age > 18 {
		fmt.Println("Вы совершеннолетний")
	} else {
		fmt.Println("Вы не совершеннолетний")
	}
}

func ifElseIf(age int) {
	if age > 18 {
		fmt.Println("Вы совершеннолетний")
	} else if age == 17 && age > 0 {
		fmt.Println("Вы почти совершеннолетний")
	} else {
		fmt.Println("Вы не совершеннолетний")
	}
}

func switchCase(fruit string) {
	switch fruit {
	case "banana":
		fmt.Println("It's a banana")
	case "kiwi":
		fmt.Println("It's a banana")
	default:
		fmt.Println("It's unknown fruit")
	}
}

func raiseErrors(age int) (bool, error) {
	if age > 18 {
		return true, nil
	}
	return false, errors.New("you are too young")
}

func main() {
	age := 17
	onlyIf(age)
	ifElse(age)
	ifElseIf(age)

	switchCase("banana")
	switchCase("banana1")

	res1, ok1 := raiseErrors(age)
	if ok1 != nil {
		fmt.Println(res1, ok1)
		fmt.Println("Вы не совершеннолетний")
		log.Fatal(ok1)
	} else {
		fmt.Println(res1, ok1)
		fmt.Println("Вы совершеннолетний")
	}
}
