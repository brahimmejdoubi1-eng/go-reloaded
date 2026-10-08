package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input file> <output file>")
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[2], ".go"){
		fmt.Println("Error: output file cannot be a .go file")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	text := string(data)

	text = Hex(text)
	text = Bin(text)
	text = ToUpper(text)
	text = ToLower(text)
	text = Cap(text)
	text = Articles(text)
	text = Quotes(text)
	text = Punct(text)

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Done! Wrote output to", os.Args[2])
}
