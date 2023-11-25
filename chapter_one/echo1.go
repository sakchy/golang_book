package main


import (
	"fmt"
	"os"
)

// comments begin with 

/*
FOR LOOP
for initialization, conditions, post {

}

WHILE LOOP
for condition {

}

INIFINITE LOOP 
for {

}
*/

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
