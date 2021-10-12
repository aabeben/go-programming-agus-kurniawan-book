package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// define slice
	fmt.Println("define slices")
	var numbers []int
	numbers = make([]int, 5)
	matrix := make([][]int, 3*3)

	// insert data
	fmt.Println(">>>>>insert slice data")
	for i := 0; i < 5; i++ {
		     numbers[i] =rand.Intn(100) 
	}
	 
	
	
	for i:=0;i<3;i++{
		matrix[i] = make([]int,3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}	
	// display data
	fmt.Println(">>>>>display slice data")
		
	
	// display matrix data
	fmt.Println(">>>>>display slice 2D data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i = %d, j = %d, matrix[%d][%d] = %v\n",i,j,i,j,matrix[i][j])
		}
	}
}