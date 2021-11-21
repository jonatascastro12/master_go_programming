package main

import "fmt"

func main() {
	const a float64 = 5.1 // typed constant

	const b = 6.7 //untyped constant

	const c float64 = a * b

	const str = "Hello " + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x

	var x = 5.5
	var y = 12 * x

	fmt.Printf("%T %T\n", x, y)

	const (
		i float64 = iota * 2
		j
		l
		m
	)

	fmt.Println(i, j, l, m)

	const xx = 7
	fmt.Printf("%T\n", xx)
	const yy float64 = xx * 3.1
	fmt.Printf("%T\n", xx)

}
