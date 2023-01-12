package main

import (
	"fmt"
	"golang-study/1_base/for_import"
	"time"
)

func main() {
	krug := for_import.Krug{Radius: 12}
	fmt.Println(krug)
	for_import.FuncFromImport("i'm ready to call because i started than capitalize letter")

	t := time.Now()
	fmt.Println(t)

	hour := time.Hour * 3
	fmt.Println(hour)
}
