package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := removeVowels(input)
	fmt.Println("Resultado:", result)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}
	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
