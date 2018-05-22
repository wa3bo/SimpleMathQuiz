package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

//variante 1: von csv nach reihe
//variante 2: randmon generieren und auswerten

func readCSV(path string) []Calc {
	var values []Calc

	file, err := os.Open(path)
	checkErr(err)

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	checkErr(err)

	for _, line := range lines {
		res, _ := strconv.Atoi(line[1])

		data := Calc{
			Res:  res,
			Show: line[0],
		}
		values = append(values, data)
	}
	return values
}

func wirteCSV(data []string) {
	file, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	checkErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	errr := writer.Write(data)
	checkErr(errr)

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
