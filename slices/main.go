package main

import (
	"fmt"
)

func main() {
	// a := []int{1, 2}
	// var b []int
	// a = append(a, 3)

	// fmt.Println(len(a))

	// b = a // a and b are point to the same underlying array

	// fmt.Println(b, len(b))

	// fmt.Printf("%p \n", a)
	// fmt.Printf("%p \n", b)

	// b = append(b, 4) // still a and b point to the same array that has data

	// fmt.Println(b, len(b))

	// fmt.Println(a, len(a))
	// fmt.Printf("%p \n", b)
	// fmt.Println(a[2])

	//slice example

	// t := []byte("string")

	// fmt.Println(len(t), t)
	// fmt.Println(t[3:5], len(t[3:5]))//5-3 = 2

	var w = [...]int{1, 2, 3} //this is and array three dots will replace to the length of elements
	var x = []int{0, 1, 2, 5, 8, 9}    //but this is a slice

	// fmt.Printf("%T", w)//[4]int
	// fmt.Printf("%T", x)//[]]int

	// c := make([]int, 5) // create a slice with a default variable length -> [0 0 0 0 0]

	// fmt.Println(c, len(c))

	y := doi(w, x)

	fmt.Println(w, x, y)

}

func doi(a [3]int, b []int) []int {
	a[0] = 4
	b[0] = 3

	c := make([]int, 5)
	c[4] = 42 //[0 0 0 0 42]
	fmt.Println(c)
	copy(c, b) //copy(dst, src []Type) int
	// c -> [0 0 0 0 42] and b -> [3 1 2] what copy do is get the items from b and replace with the existing elements that are in same index
	//if the b no of elements are not sufficent to replace all the values in c they dont change in this ex
	//c -> [3 1 2 0 42]

	//if the src has more values than the dst it will replace the elements that it has
	// c -> [0 0 0 0 42] b ->[3 1 2 5 8 9] ===> c -> [3 1 2 5 8]

	
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)

	return c
}
