// Description: Advanced struct example
// Tags: struct, advanced, struct advanced, struct advanced, advanced, advanced struct, advanced struct, json
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	Address   map[string]Address
}

type Address struct {
	Street     string
	Number     int
	Addition   string
	PostalCode int
	City       string
	Country    string
}

func main() {
	// Create a new person
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Birthday:  time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
		Address: map[string]Address{
			"home": {
				Street:     "Mainstreet",
				Number:     1,
				Addition:   "A",
				PostalCode: 1234,
				City:       "Amsterdam",
				Country:    "Netherlands",
			},
		},
	}

	// Print the person
	fmt.Printf("%v, %T, %#v\n", p, p, p)
	// Output: {John Doe 1980-01-01 00:00:00 +0000 UTC map[home:{Mainstreet 1 A 1234 Amsterdam Netherlands}]}, main.Person, main.Person{FirstName:"John", LastName:"Doe", Birthday:time.Time{wall:0x0, ext:63704560000, loc:(*time.Location)(nil)}, Address:map[string]main.Address{"home":main.Adress{Street:"Mainstreet", Number:1, Addition:"A", PostalCode:1234, City:"Amsterdam", Country:"Netherlands"}}}

	jsonPerson, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(jsonPerson))
	// Output:
	// {
	// 	"FirstName": "John",
	// 	"LastName": "Doe",
	// 	"Birthday": "1980-01-01T00:00:00Z",
	// 	"Address": {
	// 		"home": {
	// 			"Street": "Mainstreet",
	// 			"Number": 1,
	// 			"Addition": "A",
	// 			"PostalCode": 1234,
	// 			"City": "Amsterdam",
	// 			"Country": "Netherlands"
	// 		}
	// 	}
	// }
}
