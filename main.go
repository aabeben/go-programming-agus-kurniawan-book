package main

import (
	"fmt"
	"math/rand"
)
func foo(){
	fmt.Println("foo() was called!")
	fmt.Println(rand.Intn(100))
}
func main(){
	foo()
}