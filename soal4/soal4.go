package main

import "fmt"

// Node ...
type Node struct {
	left  *Node
	data  int
	right *Node
}

// New ...
func New(arr []int) *Node {
	if len(arr) < 1 {
		return &Node{}
	}

	n := &Node{nil, arr[0], nil}

	for i := 1; i < len(arr); i++ {
		n.insert(arr[i])
	}

	return n
}

func (n *Node) insert(val int) {
	if val <= n.data {
		if n.left == nil {
			n.left = &Node{nil, val, nil}
			return
		}

		n.left.insert(val)
	} else if val >= n.data {
		if n.right == nil {
			n.right = &Node{nil, val, nil}
			return
		}

		n.right.insert(val)
	}
}

func (n *Node) printnode() []int {
	array := []int{}

	if n.left != nil {
		array = append(array, n.left.printnode()...)
	}

	array = append(array, n.data)

	if n.right != nil {
		array = append(array, n.right.printnode()...)
	}

	return array
}

const INT_MIN int = -32999
const INT_MAX int = 32999

func (n *Node) isBSTHelper(max int, min int) bool {
	if n == nil {
		return true
	}

	fmt.Println("n.data", n.data, "min", min, "max", max)

	if n.data < min || n.data > max {
		return false
	}

	// return n.left.isBSTHelper(n.data-1, min) && n.right.isBSTHelper(max, n.data+1)
	return n.left.isBSTHelper(n.data, min) && n.right.isBSTHelper(max, n.data)
}

func main() {
	// arrInt := []int{1, 2, 4, 3, 7, 8}
	// arrInt := []int{3, 2, 1, 5, 4}
	arrInt := []int{3, 2, 5, 1, 4}
	fmt.Println("arrInt=", arrInt)

	bst := New(arrInt)

	fmt.Println("printnode=", bst.printnode())

	if bst.isBSTHelper(INT_MAX, INT_MIN) == true {
		fmt.Println("IS_BST")
	} else {
		fmt.Println("WRONG BST")
	}

}
