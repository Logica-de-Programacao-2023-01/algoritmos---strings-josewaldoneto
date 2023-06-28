package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isNumeric(input) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
