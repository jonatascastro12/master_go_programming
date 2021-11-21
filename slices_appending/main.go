package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	n := []int{100, 200}

	numbers = append(numbers, n...)
	fmt.Println(numbers)

	src := []int{10, 20, 30}

	dst := make([]int, 2)

	nn := copy(dst, src)

	fmt.Println(src, dst, nn)

	
}
