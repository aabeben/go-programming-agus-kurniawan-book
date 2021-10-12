package main

import "fmt"


func main(){
	a,b,c := 1.0, 2.0, 3.0
	result:=advanced_calculate(a, b, c)
	fmt.Printf("a = %.2f, b = %.2f, c = %.2f , result = %.2f\n",a, b, c, result)
}
func advanced_calculate(a, b, c float64) float64 {
	result:=a*b*c
	return result
}