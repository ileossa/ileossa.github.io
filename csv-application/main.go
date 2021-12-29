package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	filePath := "./ressources/People.csv"
	fmt.Println("Hello World!")
	writeCSV(filePath)
	persoCSV(filePath)
}

func writeCSV(path string) {
	println("Start write csv")
	f, e := os.Create(path)
	if e != nil {
		fmt.Println(e)
	}

	writer := csv.NewWriter(f)
	var data = [][]string{
		{"Name", "Age", "Occupation"},
		{"Sally", "22", "Nurse"},
		{"Joe", "43", "Sportsman"},
		{"Louis", "39", "Author"},
	}

	e = writer.WriteAll(data)
	e = writer.Write
	if e != nil {
		fmt.Println(e)
	}
}

func persoCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records)
}
