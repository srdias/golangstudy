package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	csvFile, err := os.Open("emp.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		fmt.Println("Name: " + line[0] + "; Age: " + line[1])
	}
}
