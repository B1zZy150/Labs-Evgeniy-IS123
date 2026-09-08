package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(os.Args[1:])
	fmt.Printf("  Время: %v\n", time.Since(start))
	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("  Время: %v\n", time.Since(start))
	start = time.Now()

	fmt.Println(strings.Join(os.Args[0:], " ")) //1.1
	fmt.Printf("  Время: %v\n", time.Since(start))

	start = time.Now()
	var s, sep string //1.2
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(s, i)
	}
	fmt.Printf("  Время: %v\n", time.Since(start)) //1.3
}
