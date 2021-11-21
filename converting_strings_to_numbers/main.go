package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)

	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	s1 := "3.123"

	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)

	fmt.Println(f1, err)

	if err == nil {
		fmt.Println("No Error")
	}

	fmt.Printf("%T\n", f1)

	i, err := strconv.Atoi("-50")

	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, but i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, but i value is %q\n", s2, s2)

}
