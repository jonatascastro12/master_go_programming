package main

import (
	"fmt"
)

func main() {
	var x = 3   //int type
	var y = 3.1 //float64 type

	x = x * int(y)

	fmt.Printf("%T\n", y)

	x = int(float64(x) * y)

}
