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
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0

	subSlice := nums[1 : len(nums)-2]

	for _, num := range subSlice {
		sum += num
	}

	fmt.Println("Elements:", subSlice, "Sum:", sum)

}

func ex6() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	newFriends := make([]string, len(friends))

	copy(newFriends, friends)

	newFriends[0] = "Another"

	fmt.Println(friends, newFriends)
}

func ex7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	newFriends := []string{}

	newFriends = append(newFriends, friends...)

	newFriends[0] = "Another"

	fmt.Println(friends, newFriends)
}

func ex8() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := []int{}

	newYears = append(newYears, years[0:3]...)
	newYears = append(newYears, years[len(years)-4:]...)

	fmt.Println(newYears)
}

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
	ex8()
}
