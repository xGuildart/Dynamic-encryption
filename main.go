package main

import (
	"fmt"
	"dawm.org/encryption"
)

func Windcrypt(str string, kind string, key string) string{
	return encryption.Encrypt(str, kind, key)
}

func Winddcrypt(str string, kind string, key string) string{
	return encryption.Decrypt(str, kind, key)
}

func main() {
	str := "Hello from ADMFactory.comasm"
	fmt.Println("encrypt " + str)
	enc := Windcrypt(str, "dynamic", "hello hjobhjdgki")
	fmt.Println("=> " + enc)
	fmt.Println("decrypt " + enc)
	fmt.Println("=> " + Winddcrypt(enc, "dynamic", "hello hjobhjdgki"))
}