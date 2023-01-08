package main

import (
	"fmt"
	"math"
	"reflect"
)

var global1 bool // TODO global variable
var _ bool       // TODO unused variable

func booleans() {
	var bool1 bool                     // TODO declare variable
	fmt.Println(bool1)                 // TODO print variable in console
	fmt.Println(reflect.TypeOf(bool1)) // TODO print variable type in console

	bool1 = false // TODO set variable
	fmt.Println(bool1, reflect.TypeOf(bool1))

	var bool2 bool = true // TODO initialize (declare + set) variable
	fmt.Println(bool2, reflect.TypeOf(bool2))

	bool3 := false // TODO fast initialize (declare + set) variable
	fmt.Println(bool3, reflect.TypeOf(bool3))

	const bool4 bool = true // TODO initialize (declare + set) constant(unchangeable) variable
	fmt.Println(bool4, reflect.TypeOf(bool4))

	const (
		bool5 bool = false
		bool6 bool = true
		bool7 bool = false
	) // TODO complex initialize variables
	fmt.Println(bool5, reflect.TypeOf(bool5))
	fmt.Println(bool6, reflect.TypeOf(bool6))
	fmt.Println(bool7, reflect.TypeOf(bool7))

	b1, b2, b3 := false, false, true // TODO cascade initialize variables
	fmt.Println(b1, reflect.TypeOf(b1))
	fmt.Println(b2, reflect.TypeOf(b2))
	fmt.Println(b3, reflect.TypeOf(b3))
}

func chars() {
	var string1 string = "Golang"
	fmt.Println(string1, reflect.TypeOf(string1))

	var rune1 rune = 'G'
	fmt.Println(rune1, reflect.TypeOf(rune1))
}

func numbers() {
	var int1 int = 12
	fmt.Println(int1, reflect.TypeOf(int1))

	var int2 int8 = 64
	fmt.Println(int2, reflect.TypeOf(int2))

	var int3 int16 = 64
	fmt.Println(int3, reflect.TypeOf(int3))

	var int4 int32 = 64
	fmt.Println(int4, reflect.TypeOf(int4))

	var int5 int64 = 64
	fmt.Println(int5, reflect.TypeOf(int5))

	var int6 uint = 12
	fmt.Println(int6, reflect.TypeOf(int6))

	var int7 uint8 = 64
	fmt.Println(int7, reflect.TypeOf(int7))

	var int9 uint16 = 64
	fmt.Println(int9, reflect.TypeOf(int9))

	var int10 uint32 = 64
	fmt.Println(int10, reflect.TypeOf(int10))

	var int11 uint64 = 64
	fmt.Println(int11, reflect.TypeOf(int11))

	var int12 uint32 = 64
	fmt.Println(int12, reflect.TypeOf(int12))

	var float1 float32 = 64.0
	fmt.Println(float1, reflect.TypeOf(float1))

	var float2 float64 = 64.0
	fmt.Println(float2, reflect.TypeOf(float2))
}

func arrays() {
	var arr1 [3]int = [3]int{}
	fmt.Println(arr1, reflect.TypeOf(arr1), len(arr1), cap(arr1))

	var arr2 = [5]bool{true, true, true, false, false}
	fmt.Println(arr2, reflect.TypeOf(arr2), len(arr2), cap(arr2))

	var arr3 = [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr3, reflect.TypeOf(arr3), len(arr3), cap(arr3))
}

func slices() {
	var slice1 []int = []int{1, 2, 3}
	fmt.Println(slice1, reflect.TypeOf(slice1), len(slice1), cap(slice1))

	slice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(slice2, reflect.TypeOf(slice2), len(slice2), cap(slice2))

	slice3 := []int{}
	fmt.Println(slice3, reflect.TypeOf(slice3), len(slice3), cap(slice3))

	slice4 := make([]int, 5)
	fmt.Println(slice4, reflect.TypeOf(slice4), len(slice4), cap(slice4))

	slice5 := make([]int, 5, 10)
	fmt.Println(slice5, reflect.TypeOf(slice5), len(slice5), cap(slice5))
}

func maps() {
	var map1 = make(map[string]string)
	fmt.Println(map1, reflect.TypeOf(map1), len(map1))

	map2 := make(map[any]any)
	fmt.Println(map2, reflect.TypeOf(map2), len(map2))

	map3 := map[string]int{"Bogdan": 24, "Borya": 32}
	fmt.Println(map3, reflect.TypeOf(map3), len(map3))

	map3["Alice"] = 20 // TODO set to map
	fmt.Println(map3, reflect.TypeOf(map3), len(map3))

	fmt.Println(map3["Bogdan"]) // TODO get from map

	delete(map3, "Bogdan") // TODO delete from map
	fmt.Println(map3, reflect.TypeOf(map3), len(map3))
}

func structs() {
	var user1 struct {
		name string
		age  int
	}
	fmt.Printf("%+v\n", user1)

	user2 := struct {
		name   string
		age    int
		sex    string
		weight int
		height int
	}{"Vasya", 24, "Male", 75, 175}
	fmt.Printf("%+v\n", user2)

	user3 := User{"Petya", 35, "Male", 90, 182}
	fmt.Printf("%+v\n", user3)

	fmt.Println(user3.name)

	user4 := NewUser("Alice", "Woman", 22, 50, 165)
	fmt.Printf("%+v\n", user4)

	fmt.Println(user4.getName())

	user4.setNameCopied("Ivan")
	fmt.Println(user4.getName())
	user4.setNameOriginal("Ivan")
	fmt.Println(user4.getName())
}

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func (u *User) getName() string {
	return u.name
}

func (u User) setNameCopied(name string) {
	u.name = name
}

func (u *User) setNameOriginal(name string) {
	u.name = name
}

func interfaces() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	printAnything(12)
	printAnything("12")
	printAnything(true)

	printAnythings(12, "12", false)

	square2 := Square{12}
	printShapeAreaAndPerimeter(square2)
	//circle2 := Circle{12}
	//printShapeAreaAndPerimeter(circle2) // TODO not full implemented
}

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type ShapeComplex interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}

func generics() {
	a := []int64{1, 2, 3, 4, 5}
	b := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(sumN(b))
	fmt.Println(searchElement([]string{"1", "2", "3"}, "2"))
	printAnyGen("Hello world!")
}

func sum[N int64 | float64](input []N) N {
	var result N
	for _, number := range input {
		result += number
	}
	return result
}

type Number interface {
	int64 | float64
}

func sumN[N Number](input []N) N {
	var result N
	for _, number := range input {
		result += number
	}
	return result
}

func searchElement[c comparable](elements []c, searchEl c) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}
	return false
}

func printAnyGen[A any](input A) {
	fmt.Println(input)
}

func main() {
	/*
		TODO Entry point of app
	*/

	//fmt.Println(global1, reflect.TypeOf(global1))
	//booleans()
	//chars()
	//numbers()
	//arrays()
	//slices()
	//maps()
	//structs()
	//interfaces()
	generics()
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printShapeAreaAndPerimeter(shape ShapeComplex) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func printAnything(i interface{}) {
	fmt.Println(i, reflect.TypeOf(i))
}

func printAnythings(i ...interface{}) {
	for _, value := range i {
		fmt.Println(value, reflect.TypeOf(value))
	}
}

func typeSwitch(i interface{}) {
	// TODO types casting
	str1, ok := i.(string)
	if !ok {
		fmt.Println("not a string")
	} else {
		fmt.Println(len(str1))
	}

	// TODO type switch
	switch t := i.(type) {
	case bool:
		fmt.Println("Это булево значение", t)
	case string:
		fmt.Println("Это строка", t)
	case int:
		fmt.Println("Это целочисленного значение", t)
	case float32:
		fmt.Println("Это число с плавающей точкой", t)
	default:
		fmt.Println("Это другой тип данных", t)
	}
}
