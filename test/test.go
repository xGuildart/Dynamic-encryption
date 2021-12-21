package main

import (
	"fmt"

	"dawm.org/encryption"
	//"math/rand"
)

func main() {
	str := "Hello from ADMFactory.com"
	//hx := hex.EncodeToString([]byte(str))
	//fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(str + " ==> " + encryption.HexEncode(str))
	str = ";"
	fmt.Println(str + " ==> " + encryption.HexEncode(str))
	//a := encryption.GenerateRandomHex()
	//fmt.Println(fmt.Sprintf("%d", a) + " ==> " + fmt.Sprintf("%b", encryption.IntToHex(a)))

	fmt.Println(encryption.PrintHexMatrix(encryption.GenerateMatrix()))
}
