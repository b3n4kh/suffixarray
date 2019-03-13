package main

import (
	"fmt"
	"index/suffixarray"
	"strings"
)

func mysuffix() {
	data := []byte("CAGGGTAGATGTATACGT")
	lookup := []byte("ATG")
	length := len(lookup)
	index := New(data)
	offsets := index.lookupAll(lookup)
	fmt.Println(string(data))
	for _, offset := range offsets {
		spaces := strings.Repeat(" ", offset)
		fmt.Printf("%s%s", spaces, string(data[offset:offset+length]))
	}
}

func suffix() {
	data := []byte("CAGGGTAGATGTATACGT")
	lookup := []byte("ATG")
	length := len(lookup)
	index := suffixarray.New(data)
	offsets := index.Lookup(lookup, -1)
	fmt.Println(string(data))
	for _, offset := range offsets {
		spaces := strings.Repeat(" ", offset)
		fmt.Printf("%s%s", spaces, string(data[offset:offset+length]))
	}
}

func main() {
	//suffix()
	//fmt.Println("\nMysuffix: ")
	mysuffix()
}
