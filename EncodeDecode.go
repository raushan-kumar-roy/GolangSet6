package main

import (
	"fmt"
	"strings"
)

func encode(s string) []byte {
	encoded := make([]byte, len(s))

	for i, c := range s {
		encoded[i] = byte(c + 1)
	}

	return encoded
}

func decode(encoded []byte) string {
	decoded := make([]byte, len(encoded))

	for i, b := range encoded {
		decoded[i] = b - 1
	}

	return string(decoded)
}

func main() {
	input := "Encode example with GoLang"
	encoded := encode(input)
	decoded := decode(encoded)

	fmt.Println("Original string:", input)
	fmt.Println("Encoded bytes:", encoded)
	fmt.Println("Decoded string:", decoded)

	fmt.Println("Encoded and decoded strings are equal:", strings.Compare(input, decoded) == 0)
}
