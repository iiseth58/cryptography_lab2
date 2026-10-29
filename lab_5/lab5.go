package main

import(
	"fmt"
	"encoding/base64"
	"encoding/hex"
)

func main(){
	var x string
	fmt.Print("Enter your text here: ")
	fmt.Scan(&x)

	// convert into byte 
	data := []byte(x)

	//binary 
	fmt.Print("Binary : ")
	for _, b := range data {
        fmt.Printf("%08b ", b)
    }
    fmt.Println()

	//hexadecimal 
	fmt.Println("Hexadecimal :", hex.EncodeToString(data))

	//based64
	fmt.Println("Base64 :", base64.StdEncoding.EncodeToString(data))
	
}