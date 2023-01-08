package main

import (
	"fmt"
	"reflect"
)

func inlines() {
	var1 := 12

	var1 += 1
	var1 -= 1

	var1 *= 1
	var1 /= 1

	var1 %= 1

	fmt.Println(var1, reflect.TypeOf(var1))
}

func equals() {
	var2 := false

	var2 = 12 == 1
	var2 = 12 != 1

	var2 = 12 > 1
	var2 = 12 < 1

	var2 = 12 >= 1
	var2 = 12 <= 1

	fmt.Println(var2, reflect.TypeOf(var2))
}

func main() {
	inlines()
	equals()
}
