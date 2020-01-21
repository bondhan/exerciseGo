package main

import (
	"fmt"
	"strings"
)


func Contains(mainStr string,pattern string) bool{

	if len(pattern) > len(mainStr) {
		return false
	}

	str:=strings.Split(mainStr, "")
	patt:=strings.Split(pattern, "")

	found:=false
	for i:=0;i<len(str);i++{
		found=false
		index:=i
		for j:=0;j<len(patt);j++{
			if str[index] == patt[j] {
				// fmt.Println(longStr[index], shortStr[j])
				found=true
				index++								
			} else {
				found=false
				break
			}
		}

		if found == true {
			break
		}
	}

	return found
}

func main() {

	str:="abbcde"
	pattern:="bc"


	fmt.Println(Contains(str, pattern))
}