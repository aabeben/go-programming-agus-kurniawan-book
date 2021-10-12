package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// define array
	fmt.Println("define array")
	var numbers [5]int
	var cities [5]string
	var matrix [3][3]int // array 2D

	// insert data
	fmt.Println(">>>>>insert array data")
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}
	// insert matrix data
	fmt.Println(">>>>>>insert matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}
	// display data
	fmt.Println(">>>>>>display array data")
	for j := 0; j < 5; j++ {
		fmt.Printf("number[%d] = %v\n", j, numbers[j])
		fmt.Printf("cities[%d] = %v\n", j, cities[j])
	}
	// display matrix data
	fmt.Println(">>>>>display matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d] = %v\n", i, j, matrix[i][j])
		}
	}
}
