package encryption

// s_box = [16] [16] string {
// 	{"36","5c","a7","55","98","18","0a","24","2d","6f","9a","7e","cd","76","",""},
// }

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var s_box [16][16]string = [16][16]string{
	{"cc", "47", "48", "97", "6e", "8", "59", "36", "10", "e", "25", "87", "d8", "f3", "30", "91"},
	{"49", "7a", "1d", "73", "38", "e6", "b0", "fe", "3e", "8d", "fc", "e2", "d5", "19", "f5", "b"},
	{"5c", "d6", "f9", "40", "6b", "15", "dc", "82", "a5", "64", "ed", "65", "a3", "12", "1f", "c2"},
	{"9d", "4f", "61", "85", "e1", "34", "41", "d1", "eb", "46", "94", "c7", "27", "d3", "6", "3f"},
	{"53", "4e", "ce", "89", "78", "cb", "ad", "a8", "d4", "43", "e4", "e9", "2c", "58", "80", "e0"},
	{"c6", "de", "0", "1e", "77", "be", "56", "dd", "ef", "18", "31", "92", "ac", "21", "a9", "f"},
	{"a4", "13", "cf", "2", "bb", "72", "70", "1a", "3d", "26", "4", "5f", "2b", "bc", "6c", "4a"},
	{"55", "8e", "3c", "f1", "86", "fb", "b8", "4b", "84", "90", "5b", "aa", "62", "ae", "2d", "a6"},
	{"66", "99", "52", "c1", "c5", "35", "81", "ec", "37", "c3", "51", "83", "24", "bf", "3a", "8f"},
	{"17", "6a", "1", "ab", "8a", "b9", "57", "ff", "a2", "b2", "a1", "e5", "b3", "af", "95", "6d"},
	{"d7", "63", "c", "5", "8c", "5e", "2f", "76", "df", "5d", "a", "d9", "67", "7", "da", "14"},
	{"7e", "a0", "1b", "f4", "b7", "3b", "6f", "b1", "c8", "1c", "42", "ca", "39", "a7", "68", "4d"},
	{"c0", "20", "33", "75", "db", "9e", "c4", "ba", "28", "69", "8b", "d2", "b5", "f2", "29", "5a"},
	{"e3", "54", "f0", "b4", "2e", "c9", "fd", "74", "fa", "e8", "3", "98", "ee", "7b", "93", "11"},
	{"4c", "71", "22", "9f", "96", "44", "d0", "cd", "ea", "f7", "9c", "d", "f6", "b6", "f8", "32"},
	{"79", "50", "23", "7d", "9", "16", "7c", "bd", "2a", "7f", "e7", "45", "60", "9b", "88", "9a"},
}

func AESEncryption(plainText string, key string) {

}

func SubBytes(plainText string, key string) {

}

func generateRandomHex(maxValue ...int) int {
	var maxVal int

	if len(maxValue) == 0 {
		maxVal = 256
	} else {
		maxVal = maxValue[0]
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(maxVal)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GenerateMatrix() [16][16]string {
	var a [16][16]string
	var b []int
	var maxValue int = 256
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			c := generateRandomHex(maxValue)
			for intInSlice(c, b) {
				c = generateRandomHex(maxValue)
			}
			if c == maxValue {
				maxValue--
			}
			b = append(b, c)
			a[i][j] = fmt.Sprintf("%x", c)
		}
	}
	return a
}

func HexEncode(str string) string {
	hx := hex.EncodeToString([]byte(str))
	fmt.Printf("%b", []byte(str))
	return fmt.Sprintf("%s", hx)
}

func PrintHexMatrix(b [16][16]string) string {
	a := fmt.Sprintf("%s", b)
	a = strings.Replace(a, "[[", "{\n{\"", -1)
	a = strings.Replace(a, "]]", "\"}\n}", -1)
	a = strings.Replace(a, "] ", "\"},", -1)
	a = strings.Replace(a, "[", "\n{\"", -1)
	a = strings.Replace(a, " ", "\",\"", -1)
	return a
}
