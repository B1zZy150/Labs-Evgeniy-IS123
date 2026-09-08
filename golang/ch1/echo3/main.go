package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[0:], " ")) //1.1

	var s, sep string //1.2
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(s, i)
	}

}
