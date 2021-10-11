package main

import "fmt"

func main(){
	// declare a variable
	var str string
	var n, m int
	var mn float32

	str = "Hello World!"
	n = 10
	m = 50
	mn = 2.45
	fmt.Printf("%v\n", str)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", mn)


	// Declare and assign values to variables
	var city string = "Jakarta"
	var x int = 100
	fmt.Printf("value of city: %v\n", city)
	fmt.Printf("value of x: %v\n", x)

	country:="ID"
	countryCode:=62
	fmt.Printf("value of country: %v\n", country)
	fmt.Printf("value of countryCode: %v\n", countryCode)

	// define multiple variables
	var (
		name string
		email string
		age int
	)
	name="John"
	email="john@mail.com"
	age=42
	fmt.Printf("value of name: %v\n", name)
	fmt.Printf("value of email: %v\n", email)
	fmt.Printf("value of age: %v\n", age)

}