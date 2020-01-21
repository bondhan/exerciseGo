package main

import (
	"fmt"
	"math"
)

func printArr(grid [][]int32) {
	for row := 0; row < len(grid); row++ {
		fmt.Println(grid[row])
	}

	fmt.Printf("\n\n")
}

func getRegionSize(arr [][]int32, row int, column int) int32 {
	if row < 0 || column < 0 || row >= len(arr) || column >= len(arr[row]) {
		return 0
	}

	if arr[row][column] == 0 {
		return 0
	}

	arr[row][column] = 0
	var size int32
	size = 1

	printArr(arr)

	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if r != row || c != column {
				size += getRegionSize(arr, r, c)
			}
		}
	}

	return size

}

// Complete the maxRegion function below.
func maxRegion(grid [][]int32) int32 {
	var maxRegion int32
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == 1 {
				size := getRegionSize(grid, row, column)
				fmt.Println("size", size, "maxRegion", maxRegion)
				maxRegion = int32(math.Max(float64(size), float64(maxRegion)))
			}
		}
	}

	return maxRegion
}

func main() {
	a := [][]int32{
		{1, 0, 1, 0, 0, 1},
		{1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1}}

	printArr(a)
	fmt.Println(maxRegion(a))
}
