package main

import (
	"fmt"
)

type Calc struct {
	Res  int
	Show string
}

func main() {
	var values = readCSV("./test.csv")
	fmt.Println(values)
	fmt.Println("haha")

}
