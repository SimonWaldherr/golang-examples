package main

import (
	"encoding/json"
	"fmt"
)

const jsonString = `
	[
		{
			"type": "group",
			"value": [
				"Lorem",
				"Ipsum",
				"dolor",
				"sit",
				["A", "m", "e", "t"]
			]
		},
		{
			"type": "value",
			"value": "Hello World"
		},
		{
			"type": "value",
			"value": "foobar"
		}
	]
`

func jsonforeach(in *interface{}, handler func(*string, *int, *interface{}, int)) {
	eachJsonValue(in, handler, 0)
}

func eachJsonValue(node *interface{}, handler func(*string, *int, *interface{}, int), depth int) {
	if node == nil {
		return
	}
	o, isObject := (*node).(map[string]interface{})
	if isObject {
		for k, v := range o {
			handler(&k, nil, &v, depth)
			eachJsonValue(&v, handler, depth+1)
		}
	}
	a, isArray := (*node).([]interface{})
	if isArray {
		for i, x := range a {
			handler(nil, &i, &x, depth)
			eachJsonValue(&x, handler, depth+1)
		}
	}
}

func main() {
	var j interface{}
	err := json.Unmarshal([]byte(jsonString), &j)
	if err == nil {
		jsonforeach(&j, func(key *string, index *int, value *interface{}, depth int) {
			for i := 0; i < depth; i++ {
				fmt.Print("  ")
			}
			v := *value
			switch v.(type) {
			case string:
				if key != nil {
					fmt.Printf("OBJECT: key=%q, value=%#v\n", *key, *value)
				} else {
					fmt.Printf("ARRAY: index=%d, value=%#v\n", *index, *value)
				}
			default:
				if key != nil {
					fmt.Printf("%v\n", *key)
				} else {
					fmt.Println("")
				}
			}
		})
	} else {
		fmt.Println(err)
	}
}
