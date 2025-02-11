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
	fmt.Printf("%v is hello and %[2]v is %[1]v\n", s, b)
	//so we tell that [1]v is 1st argument which is s
	//and [2]v refer to the 2nd argument which is b

	//we can get the 32 bit unicode code point using rune

	fmt.Printf("%8T %[1]v\n", []rune(s))
	//out put ->  []int32 [78 464 32 104 462 111] for each characters in "Nǐ hǎo" including the space
	//there is values like 464 that are > than 127 exist becuase this method expand using utf-8

	//if want it in sequance of byte

	fmt.Printf("%8T %[1]v\n", []byte(s))
	//output ->  []uint8 [78 199 144 32 104 199 142 111]
	//so in this case instead of large number for over sized char it used two integers for that
	//instead 464(ǐ) this use 99 144 both

	a := []byte(s)

	fmt.Printf("%d\n", len(a))
	fmt.Printf("%d\n", len(s))

	//we can only see 6 charac in "Nǐ hǎo" with the space but len return the length of the byte string that need to encode this in utf-8
	//"N"    → 1 byte
	// "ǐ"   → 2 bytes
	// " "   → 1 byte
	// "h"   → 1 byte
	// "ǎ"   → 2 bytes
	// "o"   → 1 byte
}
