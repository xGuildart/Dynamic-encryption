package main

import (
	"fmt"
	"dawm.org/encryption"
)

func main() {
	str := "Hello from ADMFactory.comasm"
	str = "hello From Drarga"
	fmt.Println("encrypt " + str)
	enc := encryption.Encrypt(str, "default", "hello hjobhjdgki")
	fmt.Println("=> " + enc)
	fmt.Println("decrypt " + enc)
	fmt.Println("=> " + encryption.Decrypt(enc, "default", "hello hjobhjdgki"))

	fmt.Println("encrypt " + str)
	enc2 := encryption.Encrypt(str, "special", "hello hjobhjdgki")
	fmt.Println("=> " + enc2)
	fmt.Println("decrypt " + enc2)
	fmt.Println("=> " + encryption.Decrypt(enc2, "special","hello hjobhjdgki"))

	str = "hello From Drarga"
	fmt.Println("encrypt " + str)
	enc3 := encryption.Encrypt(str, "dynamic", "super password K1991#")
	fmt.Println("=> " + enc3)
	fmt.Println("decrypt " + enc3)
	fmt.Println("=> " + encryption.Decrypt(enc3, "dynamic","super password K1991#"))
	r,rr,_ := encryption.GetMatrixWith(43,116)
	fmt.Printf("%d => %d", r, rr)
	// //fmt.Println(encryption.PrintHexTable(encryption.GenerateTable_x(0x16)))

	// // 6 1    ==> det = -16
	// // 4 -2
	// //m := make([][]int, 4)
	// m := []int{
	// 	6, 1, 4, -2,
	// }
	// fmt.Println(encryption.Determinant_n(m))
	// // 6 1 1
	// // 4 -2 5 ==> det = -306
	// // 2 8 7

	// n := []int{
	// 	6, 1, 1, 4, -2, 5, 2, 8, 7,
	// }
	// fmt.Println(encryption.Determinant_n(n))

	// // 2 3 1 1
	// // 1 2 3 1
	// // 1 1 2 3
	// // 3 1 1 2
	// mm := []int{
	// 	2, 3, 1, 1, 1, 2, 3, 1, 1, 1, 2, 3, 3, 1, 1, 2,
	// }
	// fmt.Println(encryption.Determinant_n(mm))

	// // 1 1 4 5
	// // 5 1 1 4
	// // 4 5 1 1
	// // 1 4 5 1

	// mm2 := []int{
	// 	1, 1, 4, 5, 5, 1, 1, 4, 4, 5, 1, 1, 1, 4, 5, 1,
	// }
	// fmt.Println(encryption.Determinant_n(mm2))

	// // 1 2 3 4         1 13 7 7
	// // 4 1 2 3    ==>  7 1 13 7
	// // 3 4 1 2    <==  7 7 1 13
	// // 2 3 4 1         13 7 7 1

	// o := []int{
	// 	1, 2, 3, 4, 4, 1, 2, 3, 3, 4, 1, 2, 2, 3, 4, 1,
	// }
	// fmt.Println(encryption.Determinant_n(o))
	// //    1  2  3
	// // B= 0  1  2
	// //    -1 -4 -1

	// // B := []int{
	// // 	1, 2, 3, 0, 1, 2, -1, -4, -1,
	// // }

	//fmt.Printf("%d", encryption.MatrixInverse(mm2, "fg"))
	// //fmt.Printf("\n%x", 0x20^0x33)
	// //fmt.Printf("%s", encryption.PrintHexTable(encryption.GenerateTable_x(13)))
	// encryption.GenerateTable_X(84)
	// fmt.Printf("%x", ((0x57 * 0x83) % encryption.PolyMin))

	// gf := encryption.DivGF(0x88, 0x06)
	// fmt.Printf("\n%s", gf.Hex_value)
	// fmt.Printf("\n%d", gf.Literal_value)
	// fmt.Printf("\n%d", gf.Fields)
	//encryption.FindMatrixWithDet1()
	//encryption.GetMatrixWith(1, 19)
	//fmt.Printf("\n%s",encryption.KeySchedule("hello hjobhjdgki"))
}
