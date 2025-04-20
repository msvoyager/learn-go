package main

import "fmt"

//function that has no name

func processNum(num int, operation func(int) int) int {
	return operation(num)
}



func main() {
	result := processNum(5, func(i int) int {
		return i *2
	})
	//in here we dont create a seperate function to double the value we just write an anonymous function inside another function call
	
	fmt.Println(result)
}