package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type a language:")
	language, _ := reader.ReadString('\n')

	switch strings.Replace(language, "\n", "", 1) {
	case "python":
		fmt.Println("Good! Python for it")
	case "go", "golang", "Golang":
		fmt.Println("Excelent! Go for it")
	default:
		fmt.Println("Never mind")
	}
}
