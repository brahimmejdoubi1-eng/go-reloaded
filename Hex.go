package main

import (
	"strconv"
	"strings"
)

// Hex handles "(hex)": it converts the previous word from hexadecimal
// to decimal and removes the marker.
func Hex(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(hex)" {
			if i >= 1 {
				l, err := strconv.ParseInt(r[i-1], 16, 64)
				if err == nil {
					r[i-1] = strconv.FormatInt(l, 10)
				}
			}
			r = append(r[:i], r[i+1:]...)
			i--
		}
	}
	return strings.Join(r, " ")
}
