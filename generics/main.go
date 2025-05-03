package main

//when using package main it means tell go to compile this to binary
//all other codes(packages attached to import) attached to this to make the binary file
import (
	"fmt"
	"slices"
)
func Print[T any](s []T) {
	//this is how a normal generic function are written 
	//any means a empty interface - interface{}
	//we can also use [S ~[]T, T any] with that ~[]T we tell that S can have the type of underlying type of slice 

}

//generics are not  only limited to functions we can declare a generic type as well

type Pair[K,V any] struct {
	Key K
	Value V
}

//then we can create a generic method describing on the pair struct
func (p Pair[K, V]) Describe() string{
	return fmt.Sprintf("Key: %v, Value: %v", p.Key, p.Value)
}
func main() {
	fmt.Println("generics")

	seq := []int{0, 1, 1, 2, 3, 5, 8}
	seq = slices.DeleteFunc(seq, func(n int) bool { //DeleteFunc is a generic function that take any type as parameters and work with it
		return n%2 != 0 // delete the odd numbers
	})
	fmt.Println(seq)


	intStringPair := Pair[int, string]{Key: 1, Value: "one"}
	fmt.Println(intStringPair.Describe())

	floatBoolPair := Pair[float32, bool]{Key: 2.00, Value: true}
	fmt.Println(floatBoolPair.Describe())


	//as above we can now pass any type to the methods also
	//it means we dont have to create deferent function for each possible data type
}