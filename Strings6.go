package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	words := strings.Fields(input)
	wordCount := len(words)
	fmt.Printf("A string contém %d palavra(s).\n", wordCount)
}
