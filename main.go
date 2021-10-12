package main 

import (
	"fmt"
	"math"
)


func circle_area(r float64){
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle area (r=%.2f) = %.2f \n",r, area)
}
func main(){
	var r float64 = 2.5
	circle_area(r)
}