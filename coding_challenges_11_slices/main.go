package main

import (
	"fmt"
	"os"
	"strconv"
)

func ex1() {
	s := []string{"el1", "el2", "el3"}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func ex2() {
	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, 6)

	a := 10.0
	mySlice[0] = a

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}

func ex3() {
	nums := []float64{1.0, 3.4, 5.6}

	nums = append(nums, 10.1)

	nums = append(nums, 4, 1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{2.2, 3.3}

	nums = append(nums, n...)

	fmt.Println(nums)
}

func ex4() {
	args := os.Args[1:]

	sum, product := 0., 1.

	if len(args) < 2 {
		fmt.Printf("You should input at least 2 args. Found %d arg(s)", len(args))
		return
	}

	if len(args) > 10 {
		fmt.Printf("You should input max of 10 args. Found %d arg(s)", len(args))
		return
	}

	for i := 0; i < len(args); i++ {
		if s, err := strconv.ParseFloat(args[i], 64); err != nil {
			fmt.Println("Non float argument found: ", args[i])
			_ = s
			return
		} else {
			sum += s
			product *= s
		}
	}

	fmt.Println("SUM:", sum, "PRODUCT:", product)
}

func ex5() {
	
}

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	ex5()
}
