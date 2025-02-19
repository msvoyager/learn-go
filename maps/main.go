package main

import "fmt"

func main() {
	var m map[string]int      //nil.but no storage - nil map there is no descriptor to a behind hash table to store keys
	p := make(map[string]int) // not-nil, but empty - empty map
	fmt.Println(m["the"])     //because of there is not key as "the" this return the empty value of the value type which is "" for string and 0 for int
	fmt.Println(p["the"])
	fmt.Printf("%T\n", m["the"])

	// m["the"] = 1 //we cant set key-value this in nil map becaise there is no storage assign to hold them program will panic
	p["the"] = 2 //this is okay because there is a place for keys

	m = p

	//this will copy the p descriptor to m so m and p both are pointing to the same hash table similar to what happen in slices
	fmt.Println(p, "\n", m)

	m["the"]++ //this will only work if the value type is numeric-increase the value by one in both m and p because it update the same hash table
	fmt.Println(p, "\n", m)

}
