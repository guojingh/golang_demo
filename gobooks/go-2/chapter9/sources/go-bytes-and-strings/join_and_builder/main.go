package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	// json
	s := []string{"I", "love", "Go"}
	fmt.Println(strings.Join(s, " "))
	b := [][]byte{[]byte("I"), []byte("love"), []byte("Go")}
	fmt.Printf("%q\n", bytes.Join(b, []byte(" ")))
	fmt.Println("=====================")

	// Builder - string
	s = []string{"I", "love", "go"}
	var builder strings.Builder
	for i, w := range s {
		builder.WriteString(w)
		if i != len(s)-1 {
			builder.WriteString(" ")
		}
	}
	fmt.Printf("%s\n", builder.String())

	// Buffer - bytes
	b = [][]byte{[]byte("I"), []byte("love"), []byte("go")}
	var buf bytes.Buffer
	for i2, w2 := range b {
		buf.Write(w2)
		if i2 != len(b)-1 {
			buf.WriteString(" ")
		}
	}
	fmt.Printf("%s\n", buf.String())
}
