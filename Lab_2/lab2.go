package main

import "fmt"

func main() {
	// Assign variable a and b 
	var a, b int
	// take the first input for variable a
	fmt.Print("input your first number: ")
	// use Scan to write into the memory 
	fmt.Scan(&a)

	// take the first input for variable b 
	fmt.Print("Input your second number: ")
	// write into the memory 
	fmt.Scan(&b)
	
	// && operator 
	// && operator only true when both condition are true
	fmt.Println("The answer of a > 0 && b > 0 is :",a > 0 && b > 0)

	// || operator 
	// || operator only false when both condition are flase 
	fmt.Println("The answer of a > b || b > a :",a > b || b > a )

	// ! operator 
	fmt.Println("The answer of !(a==b)",!(a==b))
}