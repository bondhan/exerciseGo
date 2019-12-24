package main

import (
	"fmt"
	"reflect"
	"testing"
)

type dataTest struct {
	input  []int
	output []int
}

func TestSoal2(t *testing.T) {

	test1 := dataTest{
		[]int{9, 7, 2, 2, 8, 777, 23, 31, 31},
		[]int{7, 8, 9, 23, 777, 2, 2, 31, 31},
	}

	res := SortQuantityAscending(test1.input)

	if reflect.DeepEqual(res, test1.output) == false {
		fmt.Println("res:", res)
		t.Errorf("Fail on basic sort")
	}

	test2 := dataTest{
		[]int{9999, 7777, 22222, 22222, 88, 777, 2233, 2331, 3231},
		[]int{88, 777, 2233, 2331, 3231, 7777, 9999, 22222, 22222},
	}

	res = SortQuantityAscending(test2.input)

	if reflect.DeepEqual(res, test2.output) == false {
		fmt.Println("res:", res)
		t.Errorf("Fail on basic sort")
	}
}
