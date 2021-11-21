package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi("50"); err == nil {
		fmt.Println("No err here!", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) == 2 {
		fmt.Println("We are good to continue the progam b/c we have 1 arg")
	} else {
		fmt.Println("We are not good to continue. The programa doesn't have 1 arg. It will fail if continue")

	}
}
