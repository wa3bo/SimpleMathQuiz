package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Calc struct {
	Res  int
	Show string
}

func main() {

	fmt.Println("random vlaues -> type 1 \nvalues from csv -> type 2")
	var input int = 0
	fmt.Printf(">>")
	fmt.Scanln(&input)
	if input == 1 {
		var randoemValue = generateOutPut()
		fmt.Println(randoemValue)
	} else if input == 2 {
		var values = readCSV("./test.csv")
		fmt.Println(values)
	} else {
		fmt.Println("okay, bey!")
		os.Exit(3)
	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func generateOutPut() []Calc {
	var value1 = 0
	var value2 = 0
	var values []Calc
	for i := 0; i <= 20; i++ {
		value1 = random(1, 10)
		value2 = random(1, 10)
		if i%2 == 0 && i != 0 {
			show := strconv.Itoa(value1) + "+" + strconv.Itoa(value2)
			data := Calc{
				Res:  value1 + value2,
				Show: show,
			}
			values = append(values, data)

		}

	}
	return values

}
