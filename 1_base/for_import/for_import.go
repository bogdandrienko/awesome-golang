package for_import

import "fmt"

func main() {

}

func FuncFromImport(message string) {
	fmt.Println("It's message is echo from import :" + message)
}

type Krug struct {
	Radius float32
}
