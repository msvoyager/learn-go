package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	} 
}

func main() {
	pos, neg := adder(), adder() //pos and neg become pointers to the anonymous function in adder()
	fmt.Println(pos(1), neg(1))
	fmt.Println(pos(1), neg(1))
	fmt.Println(pos(1), neg(1))

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }
}
