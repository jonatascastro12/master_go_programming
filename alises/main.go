package main

import "fmt"

func main() {
	var x byte = 2
	var y uint8 = 2
	fmt.Println(x == y)

	type second = uint
	var hour second = 3600

	fmt.Printf("Minutes in an hour: %d %T \n", hour/60, hour)

	var test int = 4.7

	_ = test
}
