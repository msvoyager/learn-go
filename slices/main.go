package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	var b []int
	a = append(a, 3)

	fmt.Println(len(a))

	b = a // a and b are point to the same underlying array

	fmt.Println(b, len(b))

	fmt.Printf("%p \n", a)
	fmt.Printf("%p \n", b)

	b = append(b, 4) // still a and b point to the same array that has data

	fmt.Println(b, len(b))
	
	fmt.Println(a, len(a))
	fmt.Printf("%p \n", b)
	fmt.Println(a[2])
}