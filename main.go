package main


import (
	"fmt"
	"math/rand"
)

func main(){
	// define map
	fmt.Println("define map")
	products:=make(map[string]int)
	customers:=make(map[string]int)


	// insert data
	fmt.Println(">>>>>insert into products map...")
	products["product1"] = rand.Intn(100)
	products["product2"] = rand.Intn(100)



	// insert into customers map
	fmt.Println(">>>>>insert into customers map...")
	customers["customer1"] = rand.Intn(100)
	customers["customer2"] = rand.Intn(100)

	// display map data
	fmt.Println(">>>>>display product map data...")
	for name, value := range products {
             fmt.Printf("%s => %d\n", name, value)
	}
	fmt.Println(">>>>display customers map data...")
	for name, value:=range customers {
		fmt.Printf("%s => %d\n", name, value)
	}
}