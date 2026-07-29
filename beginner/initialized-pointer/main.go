// Description: Initialized pointers with new expressions (Go 1.26)
// Tags: new, pointer, optional value, Go 1.26
package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Name    *string `json:"name,omitempty"`
	Retries *int    `json:"retries,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

func main() {
	// Since Go 1.26, new accepts an expression and returns a pointer to its value.
	// A separate helper such as ptr(3) is no longer necessary.
	config := Config{
		Name:    new("demo"),
		Retries: new(2 + 1),
		Enabled: new(true),
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	fmt.Printf("the retry value is %d\n", *config.Retries)
}
