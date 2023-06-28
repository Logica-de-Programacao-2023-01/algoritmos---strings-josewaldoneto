package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite uma sequência de números: ")
	fmt.Scanln(&input)

	if isConsecutive(input) {
		fmt.Println("A sequência é crescente.")
	} else {
		fmt.Println("A sequência não é crescente.")
	}
}

func isConsecutive(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		current, _ := strconv.Atoi(string(s[i]))
		next, _ := strconv.Atoi(string(s[i+1]))
		if current+1 != next {
			return false
		}
	}
	return true
}
