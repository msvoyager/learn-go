package main

import (
	"fmt"
)
//provide a platform(interface) to call methods
type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}