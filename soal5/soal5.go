package main

import "fmt"

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

func main() {
	// str := "ababaa"
	// total := 0
	// for i := 0; i < len(str); i++ {
	// 	s := str[i:]
	// 	counter := 0

	// 	for j := 0; j < len(s); j++ {
	// 		if s[j] == str[j] {
	// 			counter++
	// 		} else {
	// 			break
	// 		}
	// 	}

	// 	total = total + counter
	// }

	// fmt.Println(total)

	str := "AAAABAAAAABBBAAAAB"
	patt := "AAAB"
	ret := kmpPattSearch(str, patt)

	// for a := range ret {
	fmt.Println(ret)
	// }

}
