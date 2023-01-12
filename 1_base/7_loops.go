package main

import (
	"fmt"
	"reflect"
)

func printVariable(variables ...any) {
	/*
		TODO Функция, принимающая массив любых переменных и выводящая их на экран, вместе с типом
	*/

	// TODO Цикл по элементам [_(индекс), value(значение)]
	for _, value := range variables {
		fmt.Println(value, reflect.TypeOf(value))
	}
}

func main() {

	// TODO for loop
	row := make([]int, 0)
	for x := 0; x < 10+1; x++ {
		row = append(row, x)
	}
	fmt.Println(row)

	// TODO my realization
	//matrix := make([][]string, 0)
	//for i := 1; i < 10+1; i++ {
	//	local := make([]string, 0)
	//	for j := 1; j < 10+1; j++ {
	//		local = append(local, fmt.Sprintf("%d|%d", i, j))
	//	}
	//	matrix = append(matrix, local)
	//}
	//fmt.Println(matrix)

	// TODO custom realization
	counter := 0
	matrix := make([][]string, 10)
	for i := 1; i < 10+1; i++ {
		matrix[i-1] = make([]string, 10)
		for j := 1; j < 10+1; j++ {
			counter += 1
			matrix[i-1][j-1] = fmt.Sprintf("%d|%d|%d", i, j, counter)
		}
	}
	fmt.Println(matrix)

	// TODO for loop 1
	numbers1 := []int{10, 20, 30, 40, 50}
	for index1 := 0; index1 < len(numbers1); index1 += 1 {
		fmt.Println(index1, numbers1[index1])
	}

	// TODO for range 1
	numbers2 := []int{10, 20, 30, 40, 50}
	for index2 := range numbers2 {
		fmt.Println(index2, numbers2[index2])
	}

	// TODO for range 2
	numbers3 := []int{10, 20, 30, 40, 50}
	for index3, value3 := range numbers3 {
		fmt.Println(index3, value3)
	}

	// TODO for loop an break
	for i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if i%2 == 0 {
			continue
		}
		if i > 6 {
			break
		}
		fmt.Println(i)
	}

	// TODO loop
	sum := 0
	start := 1
	stop := 50
	for {
		if start > stop {
			break
		}
		sum += start
		start += 1
	}
	fmt.Println(sum)

	// TODO loop for map
	users := map[string]int{
		"Vasya":  15,
		"Petya":  30,
		"Kostya": 40,
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}
