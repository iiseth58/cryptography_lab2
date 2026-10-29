package main
import "fmt"
//Xor
func myXor(a, b int){
	fmt.Println("Xor operator: ",a ^ b)
}
//NOT
func myNOT(a, b int){
	fmt.Println("NOT a: ", ^a)
	fmt.Println("NOT b", ^b)
}
//OR
func myOR(a, b int){
	fmt.Println("OR operator: ", a | b)
}
//AND
func myAND(a, b int){
	fmt.Println("AND operator: ", a & b)
}
//Shift
func myShift(a uint8){
	fmt.Println("left shift", a << 1)
	fmt.Println("right shift", a >> 1)
}
// assignment operator
func myAssignment (a, b int){
	result := a 
	result += b
	fmt.Println("addition:", result)

	result = a 
	result &= b
	fmt.Println("bitwise AND assignment: ", result)

	result = a 
	result |= b
	fmt.Println("bitwise OR assignment:", result)

	result = a 
	result ^= b
	fmt.Println("bitwise XOR assigment:", result)

}

func main(){ 
	var a, b int
	fmt.Print("Enter your first number : ")
	fmt.Scan(&a)

	fmt.Print("Enter your second number : ")
	fmt.Scan(&b)
	
	fmt.Println("========Bitwise========")
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	fmt.Println("=========shift=========")
	if a >= 0 && a <= 127 {
		myShift(uint8(a))
	} else {
		println("Pls enter from 0 to 127")
	}
	fmt.Println("=========Assignment=========")
	myAssignment(a, b)
	
}