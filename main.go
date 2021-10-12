package main

import "fmt"

func add(numbers ...int) int{
	result:=0
	for _, number := range numbers{
		result+=number
	}
	return result
}
func main(){
	fmt.Printf("result = %d\n", add(1, 2, 3, 4, 5))
}