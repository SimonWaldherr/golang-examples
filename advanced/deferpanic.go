package main

import (
	"fmt"
)

func main() {
	func() {
		defer func() {
			fmt.Println("1st lvl defer")
		}()
		fmt.Println("1st lvl start")
		func() {
			defer func() {
				fmt.Println("2nd lvl defer")
				fmt.Println("recover", recover())
			}()
			fmt.Println("2nd lvl start")
			func() {
				defer func() {
					fmt.Println("3rd lvl defer")
				}()
				fmt.Println("3rd lvl start")
				func() {
					defer func() {
						fmt.Println("4th lvl defer")
					}()
					fmt.Println("4th lvl start")
					fmt.Println("panic")
					panic(0)
					fmt.Println("4th lvl finish")
				}()
				fmt.Println("3rd lvl finish")
			}()
			fmt.Println("2nd lvl finish")
		}()
		fmt.Println("1st lvl finish")
	}()
}
