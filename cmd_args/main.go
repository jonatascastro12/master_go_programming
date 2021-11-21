package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	var l int = len(os.Args)
	fmt.Printf("%d %T\n", l, l)
}
