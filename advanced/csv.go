package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadCSVfromFile(filename string) (map[int][]string, map[string]int) {
	fp, _ := os.Open(filename)
	return loadCSV(bufio.NewReader(fp))
}

func LoadCSVfromString(csv string) (map[int][]string, map[string]int) {
	fp := strings.NewReader(csv)
	return loadCSV(fp)
}

func loadCSV(reader io.Reader) (map[int][]string, map[string]int) {
	var row int
	var head = map[int][]string{}
	var data = map[int][]string{}

	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if row == 0 {
			head[row] = record
		} else {
			data[row] = record
		}
		row++
	}
	return data, GetHead(head)
}

func GetHead(data map[int][]string) map[string]int {
	head := make(map[string]int, len(data[0]))
	for pos, name := range data[0] {
		head[name] = pos
	}
	return head
}

var userdata string = `id;name;email
0;John Doe;jDoe@example.org
1;Jane Doe;jane.doe@example.com
2;Max Mustermann;m.mustermann@alpha.tld`

func main() {
	csvmap, k := LoadCSVfromString(userdata)
	for _, user := range csvmap {
		fmt.Println(user[k["name"]])
	}
}
