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
	var counter int
	for {
		counter = 0
		fmt.Println("random vlaues -> type 1 \nvalues from csv -> type 2")
		var input int = 0
		fmt.Printf(">>")
		fmt.Scanln(&input)
		if input == 1 {
			var randoemValue = generateOutPut()
			handleOutput(randoemValue)
		} else if input == 2 {
			var values = readCSV("./test.csv")
			handleOutput(values)
		} else {
			fmt.Println("okay, bye!")
			os.Exit(3)
		}

		if resultOutput(counter) {
			continue
		} else {
			fmt.Println("bye")
			break
		}
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

func handleOutput(data []Calc) int {
	var counter = 0
	for _, ele := range data {
		fmt.Println(ele.Show)
		var input int = 0
		fmt.Printf(">>")
		fmt.Scanln(&input)
		if input == ele.Res {
			counter++
		}
	}
	return counter
}

func resultOutput(counter int) bool {
	if counter <= 5 {
		fmt.Println("you are really stupid")
	} else if counter > 5 && counter <= 8 {
		fmt.Println("ok")
	} else if counter > 8 {
		fmt.Println("okok")
	}
	fmt.Println("would you like to do it again? (y n)")
	var input string
	fmt.Printf(">> ")
	fmt.Scanln(&input)
	if input == "y" {
		return true
	} else {
		return false
	}
}
