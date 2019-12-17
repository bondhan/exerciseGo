package main

import (
	"fmt"
)

/*
Jika x>y maka x = x-y dan jika y>x maka y=y-x. Berhenti jika x=y
*/

func main() {

	x := 2437
	y := 875

	for true {
		if x > y {
			x = x - y
		} else if y > x {
			y = y - x
		}

		fmt.Println("x=", x, "y=", y)
		if y == x {
			fmt.Println("x=", x)
			break
		}
	}

}
