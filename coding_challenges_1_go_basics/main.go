// CONSTANTS

// CONSTANTS - ex5
//
// package main

// import "math"

// func main() {
// 	const a int = 7
// 	const b float64 = 5.6
// 	const c = float64(a) * b

// 	x := 8
// 	var xc int = 8

// 	var noIPv6 = math.Pow(2, 128)

// 	_, _, _ = x, xc, noIPv6

// }

// CONSTANTS - ex6

// package main

// import "fmt"

// func main() {
// 	const (
// 		Jun = iota + 6
// 		Jul
// 		Aug
// 	)

// 	fmt.Println(Jun, Jul, Aug)

// }

// FMT PACKAGE - ex1

// package main

// import "fmt"

// func main() {
// 	x, y, z := 10, 15.5, "Gophers"
// 	score := []int{10, 20, 30}

// 	fmt.Printf("%d %f %s %v\n", x, y, z, score)

// 	fmt.Printf("%q\n", z)

// 	fmt.Printf("%v %v %v %v\n", x, y, z, score)

// 	fmt.Printf("%T %T\n", y, score)

// }

// FMT PACKAGE - ex2

// package main

// import "fmt"

// func main() {
// 	const x float64 = 1.422349587101

// 	fmt.Printf("%.4f", x)

// }

// FMT PACKAGE - ex2

// package main

// import "fmt"

// func main() {
// 	shape := "circle"
// 	radius := 3.2
// 	const pi float64 = 3.14159
// 	circumference := 2 * pi * radius

// 	fmt.Printf("Shape: %q\n", shape)
// 	fmt.Printf("Circle's circumference with radius %.1f is %v\n", radius, circumference)
// 	_ = shape
// }

// OPERATORS AND CONVERTIONS - ex1

// package main

// import "fmt"

// func main() {
// 	var i int = 3
// 	var f float64 = 3.2

// 	x := float64(i)
// 	y := int(f)

// 	fmt.Printf("x type is %T, while y's type is %T", x, y)
// }

// OPERATORS AND CONVERTIONS - ex2

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var i = 3
// 	var f = 3.2
// 	var s1, s2 = "3.14", "5"

// 	c_i := strconv.Itoa(i)
// 	i_s2, _ := strconv.Atoi(s2)
// 	s_f := fmt.Sprint(f)
// 	f_s1, _ := strconv.ParseFloat(s1, 64)

// 	fmt.Printf("%T %v\n", c_i, c_i)
// 	fmt.Printf("%T %v\n", i_s2, i_s2)
// 	fmt.Printf("%T %v\n", s_f, s_f)
// 	fmt.Printf("%T %v\n", f_s1, f_s1)
// }

// OPERATORS AND CONVERTIONS - ex3

// package main

// import "fmt"

// func main() {
// 	x, y := 4, 5.1

// 	z := float64(x) * y
// 	fmt.Println(z)

// 	a := 5
// 	b := 6.2 * float64(a)
// 	fmt.Println(b)
// }

// OPERATORS AND CONVERTIONS - ex4

// package main

// import "fmt"

// func main() {
// 	const SunEarthDistanceInMeters = 149.6 * 1_000_000_000
// 	const LightSpeedInMPS = 299_792_458

// 	fmt.Printf("Sunlight's time to reach the Earth: %.1f minutes", SunEarthDistanceInMeters/LightSpeedInMPS/60)
// }

// OPERATORS AND CONVERTIONS - ex5

// package main

// import "fmt"

// func main() {
// 	x, y := 0.1, 5
// 	var z float64

// 	// Write the correct logical operator (&&, || , !)
// 	// inside the expression so that result1 will be false and result2 will be true.

// 	result1 := !(x < z) && int(x) != int(z)
// 	fmt.Println(result1)

// 	result2 := y == 1*5 || int(z) == 0
// 	fmt.Println(result2)
// }

// OPERATORS AND CONVERTIONS - ex1

// package main

// import "fmt"

// func main() {
// 	type duration int

// 	var hour duration

// 	fmt.Println(hour)
// 	fmt.Printf("%T\n", hour)

// 	hour = 3600
// 	fmt.Println(hour)

// }

// OPERATORS AND CONVERTIONS - ex2

// package main

// import "fmt"

// type duration int

// func main() {
// 	type duration = int

// 	var hour duration = 3600
// 	minute := 60

// 	fmt.Println(hour != minute)
// }

// OPERATORS AND CONVERTIONS - ex3

package main

import "fmt"

type duration int

func main() {
	type mile float64
	type kilometer float64

	const m2km kilometer = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis) * m2km

	fmt.Printf("Distance from Berlin to Paris in KM: %.1fkm", kmBerlinToParis)

}
