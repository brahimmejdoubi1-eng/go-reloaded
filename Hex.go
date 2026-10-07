package main

import (
	// "fmt"
	"strconv"
	"strings"
)

func Hex(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(hex)" {
			l, _ := strconv.ParseInt(r[i-1], 16, 64)
			r[i-1] = strconv.FormatInt(l, 10)
			r = append(r[:i], r[i+1:]...)

		}
	}
	return strings.Join(r, " ")
}

// func main() {
// 	fmt.Println(Hex("1E (hex) files were added"))
// }
