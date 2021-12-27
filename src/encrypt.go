package encryption

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func subBytes(bit16 string) string {
	if len(bit16) <= 1 {
		return bit16
	}

	if bit16[:1] == "g" || bit16[1:2] == "g" {
		return bit16
	}

	bx := HexToInt(bit16[:1])
	by := HexToInt(bit16[1:2])

	return s_box[bx][by]
}

func subBytes_i(bit16 string) string {
	if len(bit16) <= 1 {
		return bit16
	}

	if bit16[:1] == "g" || bit16[1:2] == "g" {
		return bit16
	}

	bx := HexToInt(bit16[:1])
	by := HexToInt(bit16[1:2])

	return s_box_i[bx][by]
}

func KeySchedule(str string) [][][]string{
	blocks := KeyToBlock(str)
	var r [][][]string
	r = make([][][]string, 0)
	r = append(r, blocks)
	
	dynkey[0] = HexToInt(r[0][0][0])
	dynkey[1] = HexToInt(r[0][3][3])

	for i:=1; i<11; i++{
		var s [][]string = [][]string{
			{"00","0","0","0"},
			{"00","0","0","0"},
			{"00","0","0","0"},
			{"00","0","0","0"},
		}

		for j:=0;j<4;j++ {
			if j==0 {	
				for k:=0;k<4;k++{
					//rotword
					if k<3 {
						s[k][j] = r[i-1][k+1][3-j] 
					} else {
						s[k][j] = r[i-1][0][3-j]
					}
					//s-box
					s[k][j] = subBytes(s[k][j])
				}
				//addition
				for k:=0;k<4;k++{
					s[k][j] = AddGF(AddGF(HexToInt(s[k][j]), int(math.Pow(float64(2),float64(i-1)))).Literal_value,
								 HexToInt(r[i-1][k][j])).HexString
				}
			} else{
				for k:=0;k<4;k++{
					s[k][j] = AddGF(HexToInt(s[k][j-1]), HexToInt(r[i-1][k][j])).HexString
				}
			}
		}
		r = append(r,s)
	}
	return r
} 

func KeyToBlock(str string) [][]string{
	hex := HexEncode(str)
	blocks := strTo4Blocks(hex)
	n := len(blocks) - 16
	b := make([]string, 16)
	var r [][]string = [][]string{
		{"0","0","0","0"},
		{"0","0","0","0"},
		{"0","0","0","0"},
		{"0","0","0","0"},
	}

	for i:=0; i < 16 ;i++{
		b[i] = blocks[i]
	}
	for i:=0; i < n ;i++{
		b[i] =   AddGF(HexToInt(b[i]), HexToInt(blocks[16+i])).HexString
	}

	for i:= 0;i<4;i++{
		for j:=0;j<4;j++{
			r[i][j] = b[j+i*4]
		}
	}

	return r
}

func InvertSubMatrix() [16][16]string {
	var r [16][16]string
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			s := s_box[i][j]
			bx := HexToInt(s[:1])
			by := HexToInt(s[1:2])

			x := fmt.Sprintf("%x", i*16+j)
			if len(x) == 1 {
				x = "0" + x
			}
			r[bx][by] = x
		}
	}
	return r
}


func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func StrSubBytes(blocks []string) []string {
	resB := make([]string, len(blocks))
	for i := 0; i < len(blocks); i++ {
		resB[i] = subBytes(blocks[i])
	}
	return resB
}

func StrSubBytes_i(blocks []string) []string {
	resB := make([]string, len(blocks))
	for i := 0; i < len(blocks); i++ {
		resB[i] = subBytes_i(blocks[i])
	}
	return resB
}

func shiftRow(row []string, nshift int) []string {
	if len(row) <= nshift {
		return row
	}

	ret := make([]string, len(row))
	for i := 0; i < len(row); i++ {
		ret[i] = row[rotate(len(row)-1, nshift, i)]
	}
	return ret
}

func unshiftRows(block []string) []string {
	y := int(float64(len(block)) / float64(4))
	b := make([]string, len(block))
	ii := 0
	for j := 0; j < y; j++ {
		s := (4 - j%4) % 4
		c := make([]string, 4)
		for i := 0; i < 4; i++ {
			c[i] = block[i+j*4]
		}
		c = shiftRow(c, s)
		for i := 0; i < 4; i++ {
			b[i+j*4] = c[i]
			ii = i+j*4
		}
	}

	for i:=1 ;i<len(block) - ii;i++{
		b[i + ii] = block[i + ii]
	}

	return b
}

func shiftRows(block []string) []string {
	y := int(float64(len(block)) / float64(4))
	b := make([]string, len(block))
	ii :=0
	for j := 0; j < y; j++ {
		s := j % 4
		c := make([]string, 4)
		for i := 0; i < 4; i++ {
			c[i] = block[i+j*4]
		}
		c = shiftRow(c, s)
		for i := 0; i < 4; i++ {
			b[i+j*4] = c[i]
			ii = i+j*4
		}
	}

	for i:=1 ;i<len(block) - ii;i++{
		b[i + ii] = block[i + ii]
	}

	return b
}

func strTo4Blocks(hexString string) []string {
	ny := int(math.Ceil(float64(len(hexString)) / float64(8)))
	block := make([]string, len(hexString)/2)
	for j := 0; j < ny; j++ {
		for i := 0; i < 4 && 8*j+i*2 < len(hexString); i++ {
			block[i+j*4] = hexString[8*j+i*2 : 8*j+i*2+2]
		}
	}
	return block
}

func HexEncode(str string) string {
	hx := hex.EncodeToString([]byte(str))
	return fmt.Sprintf("%s", hx)
}

func HexDecode_(strb []string) string {
	for i, _ := range strb {
		x, _ := hex.DecodeString(strb[i])
		strb[i] = fmt.Sprintf("%s", x)
	}
	str := strings.ReplaceAll(strings.Join(strb, ""), "g", "")
	return str
}

func rotate(max int, step int, val int) int {
	return (val + step) % (max + 1)
}


func mixColumns(block []string) []string {
	ny := int(float64(len(block)) / float64(16))
	for i := 0; i < ny; i++ {
		for j := 0; j < 4; j++ {
			r := make([]string, 4)
			for k := 0; k < 4; k++ {
				r[k] = block[k*4+j+i*16]
			}

			if kind == "special" {
				r = mixColumn2(r)
			} else if kind == "dynamic" {
				r = dynamicMix(r)
			} else {
				// default
				r = mixColumn(r)
			}

			for k := 0; k < 4; k++ {
				block[k*4+j+i*16] = r[k]
			}
		}

	}
	return block
}

func unmixColumns(block []string) []string {
	ny := int(float64(len(block)) / float64(16))
	for i := 0; i < ny; i++ {
		for j := 0; j < 4; j++ {
			r := make([]string, 4)
			for k := 0; k < 4; k++ {
				r[k] = block[k*4+j+i*16]
			}

			if kind == "special" {
				r = unmixColumn2(r)
			} else if kind == "dynamic" {
				r = dynamicUnmix(r)
			} else { // default
				r = unmixColumn(r)
			}

			for k := 0; k < 4; k++ {
				block[k*4+j+i*16] = r[k]
			}
		}

	}
	return block
}

// 2 3 1 1
// 1 2 3 1
// 1 1 2 3
// 3 1 1 2

func mixColumn(col []string) []string {
	a := make([]byte, 4)
	for c := 0; c < 4; c++ {
		t, _ := hex.DecodeString(col[c])
		a[c] = t[0]
	}

	col[0] = hex.EncodeToString([]byte{table_2[a[0]] ^ a[3] ^ a[2] ^ table_3[a[1]]})
	col[1] = hex.EncodeToString([]byte{table_2[a[1]] ^ a[0] ^ a[3] ^ table_3[a[2]]})
	col[2] = hex.EncodeToString([]byte{table_2[a[2]] ^ a[1] ^ a[0] ^ table_3[a[3]]})
	col[3] = hex.EncodeToString([]byte{table_2[a[3]] ^ a[2] ^ a[1] ^ table_3[a[0]]})

	return col
}

func dynamicMix(col []string) []string{
	
	r,_,err := GetMatrixWith(dynkey[0],dynkey[1])
	if(err){
		col= mixColumn(col)
	} else {
		a := make([]int, 4)
		for c := 0; c < 4; c++ {
			t, _ := hex.DecodeString(col[c])
			a[c] = int(t[0])
		}
		
		col[0] = hex.EncodeToString([]byte{MulGF(r[0], a[0]).Byte_value ^ MulGF(r[1], a[1]).Byte_value ^ MulGF(r[2], a[2]).Byte_value ^ MulGF(r[3], a[3]).Byte_value})
		col[1] = hex.EncodeToString([]byte{MulGF(r[3], a[0]).Byte_value ^ MulGF(r[0], a[1]).Byte_value ^ MulGF(r[1], a[2]).Byte_value ^ MulGF(r[2], a[3]).Byte_value})
		col[2] = hex.EncodeToString([]byte{MulGF(r[2], a[0]).Byte_value ^ MulGF(r[3], a[1]).Byte_value ^ MulGF(r[0], a[2]).Byte_value ^ MulGF(r[1], a[3]).Byte_value})
		col[3] = hex.EncodeToString([]byte{MulGF(r[1], a[0]).Byte_value ^ MulGF(r[2], a[1]).Byte_value ^ MulGF(r[3], a[2]).Byte_value ^ MulGF(r[0], a[3]).Byte_value})
	}

	return col
}

func dynamicUnmix(col []string) []string{
	
	_,r,err := GetMatrixWith(dynkey[0],dynkey[1])
	if(err){
		col = unmixColumn(col)
	} else {
		a := make([]int, 4)
		for c := 0; c < 4; c++ {
			t, _ := hex.DecodeString(col[c])
			a[c] = int(t[0])
		}

		col[0] = hex.EncodeToString([]byte{MulGF(r[0], a[0]).Byte_value ^ MulGF(r[1], a[1]).Byte_value ^ MulGF(r[2], a[2]).Byte_value ^ MulGF(r[3], a[3]).Byte_value})
		col[1] = hex.EncodeToString([]byte{MulGF(r[3], a[0]).Byte_value ^ MulGF(r[0], a[1]).Byte_value ^ MulGF(r[1], a[2]).Byte_value ^ MulGF(r[2], a[3]).Byte_value})
		col[2] = hex.EncodeToString([]byte{MulGF(r[2], a[0]).Byte_value ^ MulGF(r[3], a[1]).Byte_value ^ MulGF(r[0], a[2]).Byte_value ^ MulGF(r[1], a[3]).Byte_value})
		col[3] = hex.EncodeToString([]byte{MulGF(r[1], a[0]).Byte_value ^ MulGF(r[2], a[1]).Byte_value ^ MulGF(r[3], a[2]).Byte_value ^ MulGF(r[0], a[3]).Byte_value})
	}

	return col
}

// 14 11 13 9
// 9 14 11 13
// 13 9 14 11
// 11 13 9 14

func unmixColumn(col []string) []string {
	a := make([]byte, 4)
	for c := 0; c < 4; c++ {
		t, _ := hex.DecodeString(col[c])
		a[c] = t[0]
	}
	col[0] = hex.EncodeToString([]byte{table_14[a[0]] ^ table_11[a[1]] ^ table_13[a[2]] ^ table_9[a[3]]})
	col[1] = hex.EncodeToString([]byte{table_9[a[0]] ^ table_14[a[1]] ^ table_11[a[2]] ^ table_13[a[3]]})
	col[2] = hex.EncodeToString([]byte{table_13[a[0]] ^ table_9[a[1]] ^ table_14[a[2]] ^ table_11[a[3]]})
	col[3] = hex.EncodeToString([]byte{table_11[a[0]] ^ table_13[a[1]] ^ table_9[a[2]] ^ table_14[a[3]]})

	return col
}

func mixColumn2(col []string) []string {
	a := make([]byte, 4)
	for c := 0; c < 4; c++ {
		t, _ := hex.DecodeString(col[c])
		a[c] = t[0]
	}
	col[0] = hex.EncodeToString([]byte{a[0] ^ a[1] ^ table_4[a[2]] ^ table_5[a[3]]})
	col[1] = hex.EncodeToString([]byte{table_5[a[0]] ^ a[1] ^ a[2] ^ table_4[a[3]]})
	col[2] = hex.EncodeToString([]byte{table_4[a[0]] ^ table_5[a[1]] ^ a[2] ^ a[3]})
	col[3] = hex.EncodeToString([]byte{a[0] ^ table_4[a[1]] ^ table_5[a[2]] ^ a[3]})

	return col
}

func unmixColumn2(col []string) []string {
	a := make([]byte, 4)
	for c := 0; c < 4; c++ {
		t, _ := hex.DecodeString(col[c])
		a[c] = t[0]
	}
	col[0] = hex.EncodeToString([]byte{table_x51[a[0]] ^ table_x41[a[1]] ^ table_x54[a[2]] ^ table_x45[a[3]]})
	col[1] = hex.EncodeToString([]byte{table_x45[a[0]] ^ table_x51[a[1]] ^ table_x41[a[2]] ^ table_x54[a[3]]})
	col[2] = hex.EncodeToString([]byte{table_x54[a[0]] ^ table_x45[a[1]] ^ table_x51[a[2]] ^ table_x41[a[3]]})
	col[3] = hex.EncodeToString([]byte{table_x41[a[0]] ^ table_x54[a[1]] ^ table_x45[a[2]] ^ table_x51[a[3]]})

	return col
}

func addRoundKey(block []string, keys [][][]string, round int) []string{
	ny := int(float64(len(block)) / float64(16))
	var s []string = make([]string, 0)
	for i:=0;i<len(block);i++{
		s = append(s,block[i])
	}

	for i := 0; i < ny; i++ {
		for j := 0; j < 4; j++ {
			r := make([]string, 4)
			for k := 0; k < 4; k++ {
				r[k] = s[k*4+j+i*16]	
				r[k] = AddGF(HexToInt(r[k]), HexToInt(keys[round][k][j])).HexString
			}

			for k := 0; k < 4; k++ {
				s[k*4+j+i*16] = r[k]
			}
		}
	}
	return s
}

func Decrypt(hexString string, kind_ string, key string) string {
	kind = kind_
	keys := KeySchedule(key)
	str := strTo4Blocks(hexString)

	strkey := addRoundKey(str, keys, 10)
	ush := unshiftRows(strkey)
	sub := StrSubBytes_i(ush)
	var unmix []string

	for i:=9; i>0;i--{
	strkey = addRoundKey(sub, keys, i)
	unmix = unmixColumns(strkey)
	ush = unshiftRows(unmix)
	sub = StrSubBytes_i(ush)
	}
	
	str = addRoundKey(sub, keys, 0)	
	
	return HexDecode_(str)
}

func Encrypt(str string, kind_ string, key string) string {
	kind = kind_
	hexString := HexEncode(str)
	block := strTo4Blocks(hexString)

	keys := KeySchedule(key)
	// add
	strkey := addRoundKey(block, keys, 0)
	var sub []string
	var shift []string
	var mix []string

	for i:=1;i<10;i++{
		sub = StrSubBytes(strkey)
		shift = shiftRows(sub)
		mix = mixColumns(shift)
		strkey= addRoundKey(mix, keys, i)
	}
	
	sub = StrSubBytes(strkey)
	shift = shiftRows(sub)
	mix = addRoundKey(shift, keys, 10)
	
	return strings.Join(mix, "")
}

