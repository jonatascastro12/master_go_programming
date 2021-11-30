package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ex6()
}

func ex1() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

func ex2() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func ex3() {
	count := 0
	for i := 1; i <= 50; i++ {
		if count == 3 {
			break
		}

		if i%7 == 0 {
			fmt.Println(i)
			count++
		}

	}
}

func ex4() {
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}

func ex5() {
	i := 1989

	fmt.Println("All years since I birth")

	for i <= 2021 {
		fmt.Println(i)
		i++
	}
}

func ex6() {
	rawAge := os.Args[1]

	if rawAge == "" {
		fmt.Println("No age is provided!")
		break
	}

	age, _ := strconv.Atoi(rawAge)

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

}
