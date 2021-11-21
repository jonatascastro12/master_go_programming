package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil)

	m := []int{}

	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	// fmt.Println(a == b) // compile error - cannot compare two slices - slices can only be compared to nil

	var eq bool = true

	a = nil

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	
}
