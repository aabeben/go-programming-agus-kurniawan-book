package main 

import (
	"fmt"
	"math"
)


func circle_area(r float64){
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle area (r=%.2f) = %.2f \n",r, area)
}
func calculate(a, b, c float64){
	result:=a*b*c
	fmt.Printf("a=%.2f, b=%.2f, c=%.2f result=%.2f\n", a, b, c, result)
}
func main(){
	var r float64 = 2.5
	circle_area(r)
	calculate(1.0, 2.0, 3.0)
}