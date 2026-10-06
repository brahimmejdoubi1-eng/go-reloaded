package main

import (
	"fmt"
	"strings"
)

func IsVowel(str string)bool {
	n := str[0]

	if n=='a' ||n=='e' ||n=='i' ||n=='o' ||n=='u' ||n=='h' ||n=='A' ||n=='E' ||n=='I' ||n=='O' ||n=='U' ||n=='H' {
		return true
	}
	return false 
}


func Article(str string)string {
	r := strings.Fields(str)
	for i:=0;i<len(r);i++ {
		if r[i]=="a" || r[i]=="A" {
			if IsVowel(r[i+1]) {
				r[i]+="n"
			}
		}
	}
	return strings.Join(r, " ")
}
func main() {
	fmt.Println(Article("I saw a elephant."))
}