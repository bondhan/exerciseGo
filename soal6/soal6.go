package main

import (
	"fmt"
	"strings"
)


func Contains(mainStr string,pattern string) bool{

	if len(pattern) > len(mainStr) {
		return false
	}

	str1:=strings.Split(mainStr, "")
	str2:=strings.Split(pattern, "")

	var longStr,shortStr []string

	if len(str1) > len(str2) {
		longStr = str1
		shortStr = str2
	} else {
		longStr = str2
		shortStr = str2
	}

	found:=false
	for i:=0;i<len(longStr);i++{
		found=false
		index:=i
		for j:=0;j<len(shortStr);j++{
			if longStr[index] == shortStr[j] {
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

	str:="abcde"
	pattern:="abcd"


	fmt.Println(Contains(str, pattern))
}