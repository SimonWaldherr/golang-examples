package di

import (
	"fmt"
	"io"
)

// PrintCheckList to show D.I.
//
// dependency Injection makes testing easier
// by creating functions that get all their
// dependencies are arguments
func PrintCheckList(writer io.Writer, list []string) {
	for _, item := range list {
		fmt.Fprintln(writer, item)
	}
}
