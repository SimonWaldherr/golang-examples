package main

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LogInit(infoHandle, warningHandle, errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	LogInit(os.Stdout, os.Stderr, os.Stderr)
	Info.Println("this is only a info message")
	Warning.Println("you can print warnings the same way")
	Error.Println("and even errors work as expected")

	print("\n")

	Info.Printf(`
		You can do everything, a normal log.X command can do
		In addition you can specify the output stream: stdout or stderr
		And define additional prefix und suffix strings.
	`)
}
