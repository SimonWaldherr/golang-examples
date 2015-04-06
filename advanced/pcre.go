package main

import (
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	m := pcre.MustCompile("abc", pcre.CASELESS).MatcherString("Abc", 0)
	fmt.Printf("pcre 1: %v\n", m.Matches())

	m = pcre.MustCompile("abc", 0).MatcherString("Abc", 0)
	fmt.Printf("pcre 2: %v\n", m.Matches())

	m = pcre.MustCompile("<([A-Z][A-Z0-9]*)\\b[^>]*>.*?</\\1>", pcre.CASELESS).MatcherString("<H1>foobar</H1>", 0)
	fmt.Printf("pcre 3: %v\n", m.Matches())

	m = pcre.MustCompile("<([A-Z][A-Z0-9]*)\\b[^>]*>.*?</\\1>", pcre.CASELESS).MatcherString("<H1>foobar</H2>", 0)
	fmt.Printf("pcre 4: %v\n", m.Matches())

	m = pcre.MustCompile("abc", 0).MatcherString("Abc", 0)
	fmt.Printf("pcre 5: %v\n", m.Matches())

	re := pcre.MustCompile("foo", 0)
	result := re.ReplaceAll([]byte("I like foods."), []byte("car"), 0)
	fmt.Printf("pcre 6: %v\n", string(result))

	result = re.ReplaceAll([]byte("food fight fools foo"), []byte("car"), 0)
	fmt.Printf("pcre 7: %v\n", string(result))
}
