package main

import (
	"strings"
)

// IsPunc checks if the character is a punctuation
func IsPunc(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ';' || r == ':' {
		return true
	}
	return false
}

// Punct handles punctuations removes space before them
// an add space after them
func Punct(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		if IsPunc(rune(str[i])) {
			result = strings.TrimRight(result, " ")
			result += string(str[i])
			if i+1 < len(str) && rune(str[i+1]) != ' ' {
				result += " "
			}
		} else {
			result += string(str[i])
		}
	}
	return result
}
