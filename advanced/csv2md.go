package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func fixedLengthBefore(str string, spacer string, length int) string {
	spacer = spacer[:1]
	l := length - len(str)
	if l > 0 {
		return strings.Repeat(spacer, l) + str
	}
	if l == 0 {
		return str
	}
	return str[:length]
}

func convertCSVtoMD(filename string) string {
	fp, _ := os.Open(filename)

	var row int
	var columnLength = map[int]int{}
	var data = map[int][]string{}

	csvReader := csv.NewReader(bufio.NewReader(fp))
	csvReader.Comma = ';'
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		for k := range record {
			clen := len(record[k])
			if columnLength[k] < clen {
				columnLength[k] = clen
			}
		}
		data[row] = record
		row++
	}

	var md string
	var mdHeader string
	header := true
	for row = 0; row < len(data); row++ {
		for colKey := range data[row] {
			if colKey != 0 {
				md += "|"
			}
			md += fixedLengthBefore(data[row][colKey], " ", columnLength[colKey])
			if header {
				if colKey != 0 {
					mdHeader += "|"
				}
				mdHeader += strings.Repeat("-", columnLength[colKey])
			}
		}
		md += "\n"
		if header {
			md += mdHeader + "\n"
			header = false
		}
	}
	return md
}

func main() {
	fmt.Println(convertCSVtoMD("example.csv"))
}
