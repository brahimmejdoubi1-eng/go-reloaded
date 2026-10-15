package main

import (
	"strconv"
	"strings"
)

// helper to uppercase the first character and lowercase teh rest
func Capitalize(str string) string {
	first := strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
	return first
}

// Cap handles "(cap)" or "(cap, N)": capitalize the previous
// word or (the previous N words) and remove the marker
func Cap(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(cap)" {
			if i >= 1 {
				r[i-1] = Capitalize(r[i-1])
			}
			r = append(r[:i], r[i+1:]...)
			i--
		} else if r[i] == "(cap," && i+1 < len(r) {
			numstr, ok := strings.CutSuffix(r[i+1], ")")
			if !ok {
				continue
			}
			count, err := strconv.Atoi(numstr)
			if err != nil {
				continue
			}
			if count > i {
				count = i
			}
			for j := i - count; j < i; j++ {
				r[j] = Capitalize(r[j])
			}
			r = append(r[:i], r[i+2:]...)
			i--
		}
	}
	return strings.Join(r, " ")
}
