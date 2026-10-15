package main

import (
	"strconv"
	"strings"
)

// ToLower handles "(low)" and "(low, N)": it lowercases the previous
// word (or the previous N words) and removes the marker.
func ToLower(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if r[i] == "(low)" {
			if i >= 1 {
				r[i-1] = strings.ToLower(r[i-1])
			}
			r = append(r[:i], r[i+1:]...)
			i--
		} else if r[i] == "(low," && i+1 < len(r) {
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
				r[j] = strings.ToLower(r[j])
			}

			r = append(r[:i], r[i+2:]...)
			i--
		}
	}
	return strings.Join(r, " ")
}
