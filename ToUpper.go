package main

import (
	"strconv"
	"strings"
)

// ToUpper handles "(up)" or "(up, N)": it uppercases the previous
// word or (previous N words) and removes the marker
func ToUpper(n string) string {
	r := strings.Fields(n)
	for i := 0; i < len(r); i++ {
		if r[i] == "(up)" {
			if i >= 1 {
				r[i-1] = strings.ToUpper(r[i-1])
			}
			r = append(r[:i], r[i+1:]...)
			i--
		} else if r[i] == "(up," && i+1 < len(r) {
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
				r[j] = strings.ToUpper(r[j])
			}
			r = append(r[:i], r[i+2:]...)
			i--
		}
	}
	return strings.Join(r, " ")
}
