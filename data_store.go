package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	empData := [][]string{
		{"Name", "Address"},
		{"Smith", "Newyork"},
		{"William", "Paris"},
		{"Rose", "London"},
	}
	
	csvFile, err := os.Create("employee.csv")
	
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	
	csvwriter := csv.NewWriter(csvFile)
	
	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}