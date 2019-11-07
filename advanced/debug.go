package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"runtime/debug"
	"runtime/trace"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

func source(lines [][]byte, n int) []byte {
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.Trim(lines[n], " \t")
}

func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

func debugExample() []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		line--
		fn, _ := path.Split(file)
		fmt.Fprintf(buf, "%s:%s:%d: %s\n", fn, function(pc), line+1, source(lines, line))
	}
	return buf.Bytes()
}

func fn3() {
	fmt.Printf("%v\n", string(debugExample()))
}

func fn2() {
	fn3()
}

func fn() {
	log.Printf("log inside a function")
	debug.PrintStack()
	log.Printf("%v\n", string(debug.Stack()))
	fn2()
}

func main() {
	var b bytes.Buffer
	trace.Start(&b)
	log.SetFlags(log.Ltime | log.Lshortfile)

	log.Println("first log output")
	log.Printf("second log output\n")
	fn()
	log.Printf("tace: %v\n", b.String())
}
