package main

import (
	"fmt"
	"strings"
)

// "github.com/paddie/gokmp"

//adapted from https://www.tutorialspoint.com/Knuth-Morris-Pratt-Algorithm

func findPrefix(pattern string) []int {
	length := 0
	prefArray := make([]int, len(pattern))

	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[length] {
			length++
			prefArray[i] = length
			// prefArray = append(prefArray, length)
		} else {
			if length != 0 {
				length = prefArray[length-1]
				i--
			} else {
				prefArray[i] = 0
			}
		}
	}

	return prefArray
}

func kmpPattSearch(mainString string, pattern string) []int {
	n, m, i, j := 0, 0, 0, 0
	n = len(mainString)
	m = len(pattern)
	// locArray := make([]int, (len(mainString) / len(pattern)))
	locArray := []int{}
	loc := 0

	prefixArray := findPrefix(pattern)

	for i < n {
		if mainString[i] == pattern[j] {
			i++
			j++
		}

		if j == m {
			// locArray[loc] = i - j
			locArray = append(locArray, i-j)
			loc++
			j = prefixArray[j-1]
		} else if i < n && pattern[j] != mainString[i] {
			if j != 0 {
				j = prefixArray[j-1]
			} else {
				i++
			}
		}
	}

	return locArray

}

/*
1. "" in ababaa = 6
2. babaa in ababaa = 0
3. abaa in ababaa = 3
4. baa in ababaa = 0
5. aa in ababaa = 1
6. a in ababaa = 1

*/

func getLengthFromPattern(s string, p string) int {
	// if pattern is empty string we return length of whole main string
	if p == "" {
		return len(s)
	}

	for i := 0; i < len(p); i++ {
		if strings.HasPrefix(s, p[:len(p)-i]) {
			// fmt.Println("p[:len(p)-i]=", p[:len(p)-i], "len=", len(p[:len(p)-i]))
			return len(p[:len(p)-i])
		}
	}

	return 0
}

func main() {
	str := "ababaa"

	total := 0
	for i := 0; i < len(str); i++ {
		s := str[i:]

		counter := getLengthFromPattern(str, s)

		// fmt.Printf("str=%s s=%s k=%v\n", str, s, counter)
		total = total + counter
	}

	fmt.Println(total)

	// str := "AAAABAAAAABBBAAAAB"
	// patt := ""
	// ret := kmpPattSearch(str, patt)

	// // for a := range ret {
	// fmt.Println(ret)
	// // }

}
