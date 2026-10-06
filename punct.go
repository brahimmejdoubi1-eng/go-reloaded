package main

import (
	"fmt"
	"strings"
)

func isPunc(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ';' || r == ':' {
		return true
	}
	return false
}

func Punct(str string) string {
	result := ""
	for i:=0;i<len(str);i++ {
		if isPunc(rune(str[i])) {
			result=strings.TrimRight(result," ")
			result += string(str[i])
			if i+1 < len(str) && rune(str[i+1])!=' '  {
				result += " "
			}
		} else {
			result += string(str[i])
		}
	}
	return result
}

func main() {
	fmt.Println(Punct("I was thinking...You were right"))
	fmt.Println(Punct("Wait what!? really"))
	fmt.Println(Punct("end of sentence."))
}