package main

import (
	"strconv"
	"strings"
)

func ToLower(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(low)" && i >= 1 {
			r[i-1] = strings.ToLower(r[i-1])
			r = append(r[:i], r[i+1:]...)
			i--
		} else if r[i] == "(low," && i >= 1 {
			numstr, bool := strings.CutSuffix(r[i+1], ")")
			if bool {
				count, _ := strconv.Atoi(numstr)
				for j := i - count; j < i; j++ {
					r[j] = strings.ToLower(r[j])
				}
				r = append(r[:i], r[i+2:]...)
				i--
			}
		}
	}
	return strings.Join(r, " ")
}
