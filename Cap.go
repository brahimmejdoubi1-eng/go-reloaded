package main

import (
	"strconv"
	"strings"
)

func Capitilize(str string) string {
	first := strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
	return first
}

func Cap(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(cap)" && i >= 1 {
			r[i-1] = Capitilize(r[i-1])
			r = append(r[:i], r[i+1:]...)
		} else if r[i] == "(cap," && i >= 1 {
			numstr, bool := strings.CutSuffix(r[i+1], ")")
			if bool {
				count, _ := strconv.Atoi(numstr)
				for j := i - count; j < i; j++ {
					r[j] = Capitilize(r[j])
				}
				r = append(r[:i], r[i+2:]...)
			}
		}
	}
	return strings.Join(r, " ")
}
