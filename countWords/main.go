package main

import (
	"fmt"
	"bufio"
	"os"
	// "sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "no of unique words")
	fmt.Println(words)

}