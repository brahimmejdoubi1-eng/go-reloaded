package main

import (
	"strconv"
	"strings"
)

// Bin handles "(bin)": it converts the previous word from binary
// to decimal and removes the marker.
func Bin(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(bin)" {
			if i >= 1 {
				n, err := strconv.ParseInt(r[i-1], 2, 64)
				if err == nil {
					r[i-1] = strconv.FormatInt(n, 10)
				}
			}
			r = append(r[:i], r[i+1:]...)
			i--
		}
	}
	return strings.Join(r, " ")
}
