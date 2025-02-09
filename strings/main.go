package main

import (
	"fmt"
)

func main() {
	s := "Nǐ hǎo"
	b := "hello"
	// so in this we need to pass s variable twice because we use two placeholders
	// fmt.Printf("%8T %v\n", s,s)

	// but we can refer a variable to use it for other place holder with the refering number

	fmt.Printf("%8T %[1]v\n", s) // telling that first argument wish is s to use 
	fmt.Printf("%v is hello and %[2]v is %[1]v",s,b)
	//so we tell that [1]v is 1st argument which is s 
	//and [2]v refer to the 2nd argument which is b

	

}