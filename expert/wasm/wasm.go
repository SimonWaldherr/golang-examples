package main

import (
	"fmt"
	"github.com/SimonWaldherr/ColorConverterGo"
	"github.com/SimonWaldherr/golibs/as"
	"strings"
	"syscall/js"
)

func rgb2cmykWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		input := args[0].String()
		cmyk, _ := rgb2cmyk(input)
		return cmyk
	})
	return jsonFunc
}

func rgb2cmyk(input string) (string, error) {
	inp := strings.Split(input, ",")
	if len(inp) > 2 {
		r, g, b := int(as.Int(inp[0])), int(as.Int(inp[1])), int(as.Int(inp[2]))
		c, m, y, k := colorconverter.RGB2CMYK(r, g, b)
		cmyk := fmt.Sprintf("CMYK(%v,%v,%v,%v)", c, m, y, k)
		return cmyk, nil
	}
	return "", nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("rgb2cmyk", rgb2cmykWrapper())
	<-make(chan bool)
}
