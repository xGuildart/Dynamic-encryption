package main

import (
	"fmt"
	"dawm.org/encryption"
)

import "C"

//export Windcrypt
func Windcrypt(str *C.char, kind *C.char, key *C.char) *C.char{
	return C.CString(encryption.Encrypt(C.GoString(str), C.GoString(kind), C.GoString(key)))
}

//export Winddcrypt
func Winddcrypt(str *C.char, kind *C.char, key *C.char) *C.char{
	return C.CString(encryption.Decrypt(C.GoString(str), C.GoString(kind), C.GoString(key)))
}

func main() {
	str := "Hello from ADMFactory.comasm"
	fmt.Println("encrypt " + str)
	enc := Windcrypt(C.CString(str), C.CString("dynamic"), C.CString("hello hjobhjdgki"))
	fmt.Println("=> " + C.GoString(enc))
	fmt.Println("decrypt " + C.GoString(enc))
	fmt.Println("=> " + C.GoString(Winddcrypt(enc, C.CString("dynamic"), C.CString("hello hjobhjdgki"))))
}