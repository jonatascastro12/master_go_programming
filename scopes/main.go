package main

import (
	"fmt"

	f "fmt"
)

// permitted

const done = false //package scoped

func main() {
	var tark = "Running" // local (block) scoped

	fmt.Println(tark, done)

	const done = true

	f.Println("done inside main()", done)

	f1()

	f.Println("Thif fmt is renamed")
}

func f1() {
	fmt.Printf("done inside f1(): %v \n", done) // this is done from the package scope
}
