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
	// protecting the code
	if !strings.HasSuffix(os.Args[2], ".txt") {
		fmt.Println("Error: output file must be .txt file")
		os.Exit(1)
	}
	// read the input file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	// Aplly every rule to each line
	text := string(data)
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = Hex(line)
		line = Bin(line)
		line = ToUpper(line)
		line = ToLower(line)
		line = Cap(line)
		line = Articles(line)
		line = Punct(line)
		line = Quotes(line)
		lines[i] = line
	}
	text = strings.Join(lines, "\n")
	// write result
	err = os.WriteFile(os.Args[2], []byte(text), 0o644)
	if err != nil {
		fmt.Println("error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Done! Wrote output to", os.Args[2])
}
