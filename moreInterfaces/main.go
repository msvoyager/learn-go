package main

import (
	"fmt"
)

//resource i used in this topic -> https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go

//provide a platform(interface) to call methods
type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}
//so no we have two structs and and same like method to do two things
type Stringer interface {
	String() string

}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	

	b := Book{
		Title: "learn go by practice",
		Author: "henry max",
		Pages: 1000,
	}

	Print(a,b)
}

//without using interfaces we can create a function to call the method
// func Print(a Article) {
// 	fmt.Println(a.String())
// }


func Print(s,  b  Stringer) {
	fmt.Println(s.String())
	fmt.Println(b.String())

}