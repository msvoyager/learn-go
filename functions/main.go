package main

import (
	"fmt"
)

var cars = map[string]int{} //empty map
var fleet = map[string][2]int{} // map with string keys and [2]int values
func addToCars(bodyType string, count int) {
	cars[bodyType] += count
	fmt.Println(bodyType, "added amount", cars[bodyType])
}

//return from function

func getCount(bodyType string) int {
	fmt.Printf("look for %v \n", bodyType)
	count := cars[bodyType]
	return count
}


func addToFleet(bodyType string, avCount int, bkCount int) {
	fleet[bodyType] = [2]int{avCount, bkCount}
	fmt.Println(fleet)

}
//function can return multiple values

func getStatus(bodyType string) (availble int,booked int) {
	status := fleet[bodyType]
	// return status[0], status[1] OR we can use like
	availble = status[0]
	booked = status[1]
	
	return
}

func main() {
	fmt.Println("hello")
	addToCars("suv", 10)
	addToCars("sedan", 15)
	fmt.Println(cars)
	fmt.Println(getCount("suv"))
	addToFleet("suv", 10, 5)

	a, b := getStatus("suv")
	fmt.Println(" availble", a, "\n","booked", b)


}

