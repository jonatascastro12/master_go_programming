package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}

	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}

	fmt.Println(names, len(names))

	cities := [...]string{
		5: "Paris",
		"London",
		1: "NYC",
	}

	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
}
