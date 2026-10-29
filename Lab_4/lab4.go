package main

import "fmt"

func Add(a, b int){
	fmt.Print("Enter your first number:")
	fmt.Scan(&a)
	fmt.Print("Enter your second number:")
	fmt.Scan(&b)
	fmt.Println("The result of addition is :", a + b)
}
func Sub(a, b int){
	fmt.Print("Enter your first number:")
	fmt.Scan(&a)
	fmt.Print("Enter your second number:")
	fmt.Scan(&b)
	fmt.Println("The result of substract is :", a - b)
}
func Mul(a, b int){
	fmt.Print("Enter your first number:")
	fmt.Scan(&a)
	fmt.Print("Enter your second number:")
	fmt.Scan(&b)
	fmt.Println("The result Multiply is: ", a * b )
}
func Div(a, b int){
	fmt.Print("Enter your first number:")
	fmt.Scan(&a)
	fmt.Print("Enter your second number:")
	fmt.Scan(&b)
	if b == 0 {
		fmt.Println("!!!!Error division by zero!!!!")
	}else{
	fmt.Println("The result of Divide is: ",a / b)
	}
}
func Mod(a, b int){
	fmt.Print("Enter your first number:")
	fmt.Scan(&a)
	fmt.Print("Enter your second number:")
	fmt.Scan(&b)
	fmt.Println("The result of mod is: ", a % b)
}

func main(){
	var a,b int
	fmt.Println("======Mini Calculator=====")
	fmt.Println("1)Add 2)Sub 3)Mul 4)Div 5)Mod 6)Exit")
	for{
		var x int
		fmt.Print("Choose number (1 - 6): ")
		fmt.Scan(&x)
		if x == 1 {
			Add(a, b)
		}
		if x == 2 {
			Sub(a, b)
		}
		if x == 3 {
			Mul(a, b)
		}
		if x == 4 {
			Div(a, b)
		}
		if x == 5 {
			Mod(a, b)
		}
		if x == 6 {
			break
		}
	}
}