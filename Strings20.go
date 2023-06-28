package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isCamelCase(input) {
		wordCount := countWords(input)
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func isCamelCase(s string) bool {
	if len(s) == 0 || unicode.IsUpper(rune(s[0])) {
		return false
	}
	for _, r := range s {
		if unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func countWords(s string) int {
	count := 1
	for _, r := range s {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count
}
