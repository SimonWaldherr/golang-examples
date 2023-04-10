// Description: compress and decompress a file using flate package
// Tags: compress, decompress, compress, compress file, decompress file, compress file using flate, decompress file using flate, compress file using flate package, decompress file using flate package
package main

import (
	"compress/flate"
	"io"
	"os"
)

func decompress(inputFile, outputFile string) {
	i, _ := os.Open(inputFile)
	defer i.Close()
	f := flate.NewReader(i)
	defer f.Close()
	o, _ := os.Create(outputFile)
	defer o.Close()
	io.Copy(o, f)
}

func compress(inputFile, outputFile string) {
	i, _ := os.Open(inputFile)
	defer i.Close()
	o, _ := os.Create(outputFile)
	defer o.Close()
	f, _ := flate.NewWriter(o, flate.BestCompression)
	defer f.Close()
	io.Copy(f, i)
}

func main() {
	compress("compress.go", "compress.min")
	decompress("compress.min", "compress_1.go")
}
