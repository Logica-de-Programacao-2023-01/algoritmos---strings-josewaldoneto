package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	str = removeSpacesAndPunctuation(str)
	reversed := reverseString(str)
	return str == reversed
}

func removeSpacesAndPunctuation(str string) string {
	punctuation := `.,;:?!()[]{}"'`
	for _, char := range punctuation {
		str = strings.ReplaceAll(str, string(char), "")
	}
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
