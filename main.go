package main

import "fmt"

func main(){
	var str string = "Hello"
	var n, m int = 1, 2
	var mn , _ float32 = 1.5, 4.5

	fmt.Printf("value: %v, type: %T\n",str, str)
	fmt.Printf("value: %v, type: %T\n",n,n)
	fmt.Printf("value: %v, type: %T\n",m, m)
	fmt.Printf("value: %v, type: %T\n", mn, mn)
}