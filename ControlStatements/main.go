package main

import (
	"fmt"
)
	


func main() {
	fmt.Println("hello")

	var myMap =  map[string]string {
		"name" : "minuda",
		"age"  : "20",
		"education" : "A/L",
	} 

	for k,v := range myMap {
		fmt.Println(k, v)
	}

	//blank identifier to get the value only

	for _,v := range myMap {
		fmt.Println(v)
		
	}



	//infinite loop
	i, j := 0, 3
	for {
		i, j = i + 50, j * j
		fmt.Println(i,j)
		if j > i {
			break //stop the infinite loop on condition
		}
	}



}