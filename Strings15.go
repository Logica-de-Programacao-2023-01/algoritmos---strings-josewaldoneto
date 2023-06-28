package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	output := replaceVowels(input)
	fmt.Println(output)
}

func replaceVowels(s string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, vowel := range vowels {
		s = strings.ReplaceAll(s, vowel, "*")
	}
	return s
}
