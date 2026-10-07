package main

import (
	"strings"
)

func Quotes(str string) string {
	r := strings.Fields(str)
	openindex := -1
	for i := 0; i < len(r); i++ {
		if r[i] == "'" {
			if openindex == -1 {
				openindex = i
			} else {
				r[openindex+1] = "'" + r[openindex+1]
				r[i-1] = r[i-1] + "'"
				r = append(r[:openindex], append(r[openindex+1:i], r[i+1:]...)...)
			}
		}
	}
	return strings.Join(r, " ")
}
