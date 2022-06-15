package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "golu"
	records := readCsvFile("./read.csv")

	city := getCity(records, input)
	fmt.Println(city)

}

func getCity(records [][]string, input string) string {
	var cityName string
	for i := 0; i < len(records); i++ {
		if records[i][0] == input {
			cityName = records[i][1]
			break
		}

	}
	return cityName
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
