package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) - 1
	left := 0

	for i := range arr {
		if arr[i] < arr[pivot] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[pivot] = arr[pivot], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

// Number ...
type Number struct {
	number  int
	counter int
}

func quickSortNumber(arr []Number) []Number {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) - 1
	left := 0

	for i := range arr {
		if arr[i].counter < arr[pivot].counter {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[pivot] = arr[pivot], arr[left]

	quickSortNumber(arr[:left])
	quickSortNumber(arr[left+1:])

	return arr
}

func SortQuantityAscending(array []int) []int {

	sortedArr := make([]int, len(array))

	copy(sortedArr, array)
	sortedArr = quickSort(sortedArr)

	// fmt.Println(array)
	// fmt.Println(sortedArr)

	sortedArray := make(map[int]int)

	for a, b := range sortedArr {
		// fmt.Println("a", a, "b", b)
		_ = a
		sortedArray[b]++
	}

	// fmt.Println(sortedArray)

	tmpArray := []Number{}

	for a, b := range sortedArray {
		n := Number{a, b}
		tmpArray = append(tmpArray, n)
	}

	// fmt.Println(tmpArray)

	quickSortNumber(tmpArray)

	// fmt.Println(tmpArray)

	for a := range tmpArray {
		for c := range tmpArray {
			if tmpArray[a].counter == tmpArray[c].counter && a != c {
				if tmpArray[a].number < tmpArray[c].number {
					tmpArray[a], tmpArray[c] = tmpArray[c], tmpArray[a]
				}
			}
		}
	}

	result := []int{}
	for i := 0; i < len(tmpArray); i++ {
		for j := 0; j < tmpArray[i].counter; j++ {
			// fmt.Println(tmpArray[i].number)
			result = append(result, tmpArray[i].number)
		}
	}

	return result
}

func main() {
	array := []int{1, 2, 4, 2, 2, 3, 6, 6, 4, 3, 7}

	fmt.Println(SortQuantityAscending(array))

}
