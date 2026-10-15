package main

import (
	"strings"
)

// IsVowel checks if the 1st character of a string is a vowel or an h
func IsVowel(str string) bool {
	n := str[0]

	if n == 'a' || n == 'e' || n == 'i' || n == 'o' || n == 'u' || n == 'h' || n == 'A' || n == 'E' || n == 'I' || n == 'O' || n == 'U' || n == 'H' {
		return true
	}
	return false
}

// Articles converts "a" or "A" into "an" or "An" if the next word
// starts with a vowel or an h
func Articles(str string) string {
	r := strings.Fields(str)
	for i := 0; i < len(r); i++ {
		if (r[i] == "a" || r[i] == "A") && i+1 < len(r) && IsVowel(r[i+1]) {
			r[i] += "n"
		}
	}
	return strings.Join(r, " ")
}
