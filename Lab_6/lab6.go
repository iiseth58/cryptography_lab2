package main

import "fmt"

// function for encrypt and decrypt 
func xorEncrypt(text string, key byte) string {
	data := []byte(text)
	for i := 0; i < len(data); i++{
		data[i] = data[i] ^ key
	}
	return string(data)
}

func main(){
	var text string 
	var key uint8

	fmt.Print("Enter plaintext: ")
    fmt.Scan(&text)

    fmt.Print("Enter key (0-255): ")
    fmt.Scan(&key)

    // Encrypt 
    encrypted := xorEncrypt(text, key)

    // Decrypt 
    decrypted := xorEncrypt(encrypted, key)

    fmt.Println("Original:", text)
    fmt.Printf("Encrypted: %x\n", encrypted)
    fmt.Println("Decrypted:", decrypted)
}