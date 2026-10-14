package main

import (
	"strconv"
	"strings"
)

func ToUpper(n string) string {
	r := strings.Fields(n)
	for i := 0; i < len(r); i++ {
		if r[i] == "(up)" && i >= 1 {
			r[i-1] = strings.ToUpper(r[i-1])
			r = append(r[:i], r[i+1:]...)
			i--
		} else if r[i] == "(up," && i >= 1 {
			numstr, bool := strings.CutSuffix(r[i+1], ")")
			if bool {
				count, err := strconv.Atoi(numstr)
				if err == nil {
					for j := i - count; j < i; j++ {
						r[j] = strings.ToUpper(r[j])
					}
					r = append(r[:i], r[i+2:]...)
					i--
				}
			}
		}
	}
	return strings.Join(r, " ")
}
