package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a +b 
		return b
	}
}

func main() {
	fmt.Println("fibonacci Sequence")

	f := fib() //f and g get defferent a & b variables

	// fmt.Println(f(),g(), f(), g(), f())
	// fmt.Println(g(), g(), g(), g(), g())


	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
	
}
// a b 
// 1 1
// 1 2
// 2 3
// 3 5
// 5 8
// 8 13
