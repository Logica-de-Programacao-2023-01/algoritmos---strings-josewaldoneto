package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isDescending := checkDescending(input)
	if isDescending {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}

func checkDescending(str string) bool {
	if len(str) < 2 {
		return false
	}

	previous, _ := strconv.Atoi(string(str[0]))
	for i := 1; i < len(str); i++ {
		current, _ := strconv.Atoi(string(str[i]))
		if current >= previous {
			return false
		}
		previous = current
	}
	return true
}
