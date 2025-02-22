package main

import "fmt"

func main() {

	//the type used for the key must have ==(equality) and !=(inequality) operational
	//string numbers arrays are good but slices or maps not good
	//we can only check equality of map is to nil or not nil

	var m map[string]int      //nil.but no storage - nil map there is no descriptor to a behind hash table to store keys
	p := make(map[string]int) // not-nil, but empty - empty map
	fmt.Println(p, "\n", m)
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


	//create a map with key  value pair

	var c = map[string]int {
		"first" : 1,
		"second" : 2,
	}
	fmt.Println(c)
	var d map[string]int

	// b := c == d //we expected to answer to be false but we cant compare because comparison operators not work in maps
	//we can only compare to nil
	e := d == nil //the return is true
	
	fmt.Println(e)
	fmt.Println(len(c)) //-> 2


	//there is another way to make a empty map 

	f := map[string]int{}

	// a := f["the"] //this return 0 but how can we sure is there is no key as "the" or there is a key as "the" and its value is 0

	//we can check it by

	a, ok := f["the"] // if ok = "false" it means there is no key as given
	fmt.Println(a, ok) //0, false
	//this is using in conditional statements to make sure return was not the default value

	if w, ok := f["the"]; ok {
		fmt.Println("there was a key called 'the'", w)
	}

	


}
