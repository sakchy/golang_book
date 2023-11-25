package main

import (
	"fmt"
	"os"
)


func main() {
	s, sep := "", ""
	// if Args[1:] is not selected 
	// it prints path 
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}