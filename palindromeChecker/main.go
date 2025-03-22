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
	//what we do is campare the rune of input and its reversed version
	//rune is the integer representation of unicode point
	var b string
	for l := len(uInput)-1; l>-1; l-- {
		b += string(uInput[l]) //when we convert using string the unicode make into letters
	}

	if uInput == b {
		fmt.Println("its a palindrome")
	} else {
		fmt.Println("its not a palindrome")
	}

	// for _, r := range uInput{ 									// r get the rune of character
	// 	fmt.Println(string(r))
	// }
}