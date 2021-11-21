package main

import "fmt"

func main() {
	// INT TYPE
	var i1 int8 = -128

	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535

	fmt.Printf("%T\n", i2)

	// FLOAT TYPE
	var f1, f2, f3 = 1.1, -.2, 5.

	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// RUNE TYPE = alias of int32
	var r rune = 'a'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	//BOOL TYPE
	var b bool = true
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	// ARRAY TYPE
	var numbers = [4]string{"4", "5", "-9", "100"}
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)

	// SLICE TYPE (similar to Python list)
	var cities = []string{"London", "Tokyo", "Washignton"}
	fmt.Printf("%T\n", cities)
	fmt.Println(cities)

	// MAP TYPE (similar to Python dict)
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)
	fmt.Println(balances)

	// STRUCT TYPE

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T", you)

	// POINTER TYPE

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// FUNCTION TYPE
	fmt.Printf("%T\n", f)

}

func f() {

}
