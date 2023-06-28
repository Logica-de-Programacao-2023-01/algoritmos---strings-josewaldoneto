package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = sortString(str1)
	str2 = sortString(str2)

	if str1 == str2 {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
