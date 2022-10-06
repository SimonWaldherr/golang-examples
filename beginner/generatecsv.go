package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"Language", "Version", "Date"},
		{"Java", "1.19", "2022-08-02"},
		{"Go", "Paris", "2022-09-22"},
	}

	csvfile, err := os.Create("languages.csv")

	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()

	csvfile.Close()
}
