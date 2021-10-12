package main

import "fmt"

func main() {
	// iteration -for
	var i int
	for i = 0; i < 5; i++ {
		if i == 3 {
			fmt.Printf("i = %v\n\n", i)
			break
		}
		fmt.Printf("i = %v\n", i)
	}
	for j := 5; j < 11; j++ {
		if j == 7 {
			continue
		}
		fmt.Printf("j = %v\n", j)
	}
}
