package main

import (
	"strconv"
	"strings"
)


func ToLower(str string)string {
	r := strings.Fields(str)
    for i:= 0; i<len(r); i++ {
		if r[i] == "(low)" {
			r[i-1] = strings.ToLower(r[i-1])
			r = append(r[:i], r[i+1:]...)
		}else if r[i] == "(low," {
			numstr := strings.TrimSuffix(r[i+1], ")")
			count,_ := strconv.Atoi(numstr)
			for j := i-count; j<i;j++ {
				r[j]=strings.ToLower(r[j])
			}
			r= append(r[:i],r[i+2:]...)
		}
	}
	return strings.Join(r," ")
}