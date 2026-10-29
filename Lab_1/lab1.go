package main

import "fmt"

func main(){
	// assign variable a and b as integer
	var a, b int
    // take the input for variable a 
	fmt.Print("Enter your first number: ")
	fmt.Scan(&a)
    // take the input for variable b
	fmt.Print("Enter your second number: ")
	fmt.Scan(&b)

	result := a 
	fmt.Println("Assignment (=): ", result)

	result = a
	result += b
	fmt.Println("Addition (+=): ", result)

	result = a
	result -= b 
	fmt.Println("Substrction (-=) ", result)

	result = a 
	result *= b
	fmt.Println("multiply (*=)", result)

	result = a 
	result /= b
	fmt.Println("multiply (/=)", result)

	result = a 
	result %= b
	fmt.Println("multiply (%=)", result)

}