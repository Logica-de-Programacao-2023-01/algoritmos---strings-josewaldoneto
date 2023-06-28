package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	uniqueLetters := findUniqueLetters(input)
	fmt.Println("Letras únicas:", uniqueLetters)
}

func findUniqueLetters(str string) string {
	letterCount := make(map[rune]int)

	for _, char := range str {
		letterCount[char]++
	}

	uniqueLetters := ""
	for _, char := range str {
		if letterCount[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	return uniqueLetters
}
