package main

import "fmt"

func main() {
	var age = 30

	fmt.Println("Age: ", age)

	var name = "Dan"
	_ = name

	s := "Test"

	fmt.Println(s)

	car, cost := "Audi", 3000
	fmt.Println(car, cost)

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int = 10, 20, 30
	c, a = a, c
	fmt.Println(a, b, c)

}
