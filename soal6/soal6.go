package main

import (
	"fmt"
	"strings"
)


func main() {

	string1:="abcde"
	string2:="ab"

	str1:=strings.Split(string1, "")
	str2:=strings.Split(string2, "")

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
				fmt.Println(longStr[index], shortStr[j])
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

	fmt.Println(string1,string2)
	fmt.Println(found)
}