package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	fmt.Println("Resultado:", reversed)
}
