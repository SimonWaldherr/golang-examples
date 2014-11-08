package main

import (
	"compress/gzip"
	"flag"
	"io"
	"os"
)

var n = flag.Int("n", 6, "Compression level (0-9)")

func main() {
	flag.Parse()
	c, _ := gzip.NewWriterLevel(os.Stdout, *n)

	io.Copy(c, os.Stdin)
	c.Close()
}
