package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("pallindrome checking programme enter your word below:")
	var uInput string
	fmt.Scanln(&uInput)
	uInput = strings.ToUpper(uInput)
	var b string
	for l := len(uInput)-1; l>-1; l-- {
		b += string(uInput[l])
	}

	if uInput == b {
		fmt.Println("its a palindrome")
	} else {
		fmt.Println("its not a palindrome")
	}
}