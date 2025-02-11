package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)
func main(){
	if len(os.Args) < 3{
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(-1)
	}

	old, new := os.Args[1],os.Args[2]
	fmt.Println(os.Args[0],os.Args[1],os.Args[2])
	scan := bufio.NewScanner(os.Stdin)

	

	for scan.Scan() { //this act like a while loop
		
		s := strings.Split(scan.Text(), old,) //split the read line into slice removing old-word
		fmt.Println(s)
		//for example s = ["", " are in the"]
		t := strings.Join(s, new) // join the slice using new word

		fmt.Println(t)

	}

}