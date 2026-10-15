package main

import (
	"strings"
)

// Quotes handels pairs of "'": the quotes are glued to the fisrt and last
// word between them, with no spaces.
func Quotes(str string) string {
	r := strings.Fields(str)
	openindex := -1
	for i := 0; i < len(r); i++ {
		if r[i] == "'" {
			if openindex == -1 {
				openindex = i
			} else {
				if i-openindex > 1 {
					r[openindex+1] = "'" + r[openindex+1]
					r[i-1] = r[i-1] + "'"
				}
				r = append(r[:openindex], append(r[openindex+1:i], r[i+1:]...)...)
				i = openindex - 1
				openindex = -1
			}
		}
	}
	return strings.Join(r, " ")
}
