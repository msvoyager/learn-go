package main

import "fmt"

//variadic function can have many parameters as we like

func sum(numbers ...int) (total int) {

	for _,n := range numbers {
		total += n
		
	}

	return
}


//function as a parameter

func processNum(num int, operation func(int) int) int {
	return operation(num)
}

func double(num int) int {
	return num * 2
}

func half(num int) int {
	return num /2
}



func main() {
	
	t := sum(1,1,1,1,1,1,1,1,1)

	fmt.Println(t)

	//we can pass the parameters of a slice using spreading operator -> ...

	odds := []int{1,1,1,1,1,1,1}
	fmt.Println("sum", sum(odds...))

	result := processNum(10, half) //double is a function

	fmt.Println(result)


}