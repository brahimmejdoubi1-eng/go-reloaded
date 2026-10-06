package main

import (
	// "fmt"
	"strconv"
	"strings"
)


func Bin(str string)string {
	r := strings.Fields(str)
	for i := 0; i<len(r);i++ {
		if r[i] == "(bin)" {
			l,_ := strconv.ParseInt(r[i-1],2,64)
			r[i-1]=strconv.FormatInt(l,10)
			r=append(r[:i],r[i+1:]...)
		}
		
	}
	return strings.Join(r, " ")
}

// func main() {
// 	fmt.Println(Bin("10 (bin) years"))
// }