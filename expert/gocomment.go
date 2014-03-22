package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) >= 2 {
		var output string
		var i int
		input := os.Args[1]
		gosrcs, _ := ioutil.ReadFile(input)
		gosrc := string(gosrcs)
		gorun, _ := exec.Command("go", "build", input).Output()
		gostr := string(gorun)
		aline := regexp.MustCompile("[^\\n]*\\n")
		line, _ := regexp.Compile(input + ":(\\d+): ")
		stre := line.FindAllStringSubmatch(gostr, -1)
		stra := aline.FindAllString(gosrc, -1)
		j := 0
		for i = 0; i < len(stra); i++ {
			if len(stre) > j {
				lc, _ := strconv.ParseInt(stre[j][1], 10, 0)
				if i == int(lc)-1 {
					output += "//" + stra[i]
					j++
				} else {
					fmt.Println(i)
					fmt.Println(stre[j][1])
					output += stra[i]
				}
			} else {
				output += stra[i]
			}
		}
		outputb := []byte(output)
		ioutil.WriteFile(input, outputb, 0755)
	}
}
