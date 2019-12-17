package main

import "fmt"

func main() {
	str := "ababaa"
	total := 0
	for i := 0; i < len(str); i++ {
		s := str[i:]
		counter := 0

		for j := 0; j < len(s); j++ {
			if s[j] == str[j] {
				counter++
			} else {
				break
			}
		}

		total = total + counter

	}

	fmt.Println(total)

}
