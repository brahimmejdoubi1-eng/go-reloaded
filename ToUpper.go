package main

import (
	"strconv"
	"strings"
)

func ToUpper(n string) string {
	r := strings.Fields(n)
	for i := 0; i < len(r); i++ {
		if r[i] == "(up)" && i>=1 {
			r[i-1] = strings.ToUpper(r[i-1])
			r = append(r[:i], r[i+1:]...)
		} else if r[i] == "(up," && i>=1 {
			numstr := strings.TrimSuffix(r[i+1], ")")
			count, _ := strconv.Atoi(numstr)
			for j := i - count; j < i; j++ {
				r[j] = strings.ToUpper(r[j])
			}
			r = append(r[:i], r[i+2:]...)
		}
	}
	return strings.Join(r, " ")
}
