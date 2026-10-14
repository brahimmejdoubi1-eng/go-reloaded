package main

import (
	"strconv"
	"strings"
)

func Bin(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(bin)" && i>0{
			l, err := strconv.ParseInt(r[i-1], 2, 64)
			if err == nil {
			r[i-1] = strconv.FormatInt(l, 10)
			r = append(r[:i], r[i+1:]...)
			i--
			}
		}

	}
	return strings.Join(r, " ")
}
