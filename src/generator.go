package encryption

import (
	"fmt"
	"math"
	"io/ioutil"
	"io/fs"
	"strings"
	"math/rand"
	"time"
)

func matrice(m []int, i int, j int) []int {
	n := math.Sqrt(float64(len(m)))
	var r []int = make([]int, 0)
	r = make([]int, int((n-1)*(n-1)))
	for k := 0; k < int(n-1); k++ {
		for s := 0; s < int(n-1); s++ {
			if s < j && k < i {
				r[k*int(n-1)+s] = m[k*int(n)+s]
			} else if s < j && k >= i {
				r[k*int(n-1)+s] = m[(k+1)*int(n)+s]
			} else if s >= j && k < i {
				r[k*int(n-1)+s] = m[k*int(n)+s+1]
			} else {
				r[k*int(n-1)+s] = m[(k+1)*int(n)+s+1]
			}
		}
	}
	return r
}

func determinant_2(m []int) int {
	r := 0
	if ev == "decimal" {
		r = m[0]*m[3] - m[1]*m[2]
	} else {
		r = MulGF(m[0], m[3]).Literal_value ^ MulGF(m[1], m[2]).Literal_value
	}
	return r
}

func Determinant_n(m []int) int {
	n := math.Sqrt(float64(len(m)))
	var d int = 0
	if n == 2 {
		d = determinant_2(m)
	} else {
		for i := 0; i < int(n); i++ {
			var new_m []int = matrice(m, 0, i)
			var r int
			if ev == "decimal" {
				r = int(math.Pow(-1, float64(i))) * m[i] * Determinant_n(new_m)
				d += r
			} else {
				r = MulGF(m[i], Determinant_n(new_m)).Literal_value
				d ^= r
			}

		}
	}
	return d
}

func BringTcomAndDet(m []int, ev_ string) {
	ev = ev_
	fmt.Printf("%d, %d", Transpose_Matrice(Com_matrice(m)),
		Determinant_n(m))
}

func Com_matrice(m []int) []int {
	n := math.Sqrt(float64(len(m)))
	var r []int = make([]int, 0)
	r = make([]int, int(n*n))
	if n == 2 {
		if ev == "decimal" {
			r[0] = m[3]
			r[1] = -1 * m[2]
			r[2] = -1 * m[1]
			r[3] = m[0]
		} else {
			r[0] = m[3]
			r[1] = m[2]
			r[2] = m[1]
			r[3] = m[0]
		}
	} else {

		for i := 0; i < int(n); i++ {
			for j := 0; j < int(n); j++ {
				if ev == "decimal" {
					r[i*int(n)+j] = int(math.Pow(-1, float64(i+j))) * Determinant_n(matrice(m, i, j))
				} else {
					r[i*int(n)+j] = Determinant_n(matrice(m, i, j))
				}

			}
		}
	}

	return r
}

func Transpose_Matrice(m []int) []int {
	n := math.Sqrt(float64(len(m)))
	var r []int = make([]int, 0)
	r = make([]int, int(n*n))
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			r[i*int(n)+j] = m[j*int(n)+i]
		}
	}
	return r
}

func MatrixInverse(m []int, ev_ string) []int {
	ev = ev_
	det := Determinant_n(m)
	n := math.Sqrt(float64(len(m)))
	var r []int = make([]int, 0)
	//fmt.Printf("det=%d\n", det)
	if det != 0 {
		r = make([]int, int(n*n))
		tcom := Transpose_Matrice(Com_matrice(m))
		//fmt.Printf("tcom=%d\n", tcom)
		for i := 0; i < int(n); i++ {
			for j := 0; j < int(n); j++ {
				r[i*int(n)+j] = DivG(tcom[i*int(n)+j], det)
			}
		}
	} else {
		fmt.Printf("this matrix doesn't have inverse")
	}
	return r
}

func DivG(a int, b int) int {
	var r int = 0
	if b != 0 && ev == "decimal" {
		r = a / b
	} else {
		r = DivGF(a, b).Literal_value
	}

	return r
}


func FindMatrixWithDet1() {
	ev = "fg"
	var r []int = make([]int, 0)
	r = make([]int, 16)
	for i := 1; i < 256; i++ {
		for j := 1; j < 256; j++ {
			for k := 1; k < 256; k++ {
				for l := 1; l < 256; l++ {
					r = getCurculantMatrix(i, j, k, l)
					if Determinant_n(r) == 1 {
						s := fmt.Sprintf("M=%d ==> M^-1=%d\n", r, MatrixInverse(r, "ev"))
						fmt.Printf("matrix with det == 1 ==>\n%s", s)
						ioutil.WriteFile("matrixGaloisWithDet1.txt", []byte(s), fs.ModeAppend.Type())
					}
				}
			}
		}
	}
}

func getCurculantMatrix(i int, j int, k int, l int) []int {
	var r []int = make([]int, 0)
	r = make([]int, 16)
	r = []int{i, j, k, l, l, i, j, k, k, l, i, j, j, k, l, i}
	return r
}

func GetMatrixWith(i int, j int) ([]int, []int, bool) {
	ev = "fg"
	var r []int = make([]int, 0)
	r = make([]int, 16)

	for j-i>25 {
		i++
		j--
	}

	min := 255 * 2
	kmin := i
	lmin := i
	found := false


	for k := i; k < j; k++ {
		for l := i; l < j; l++ {
			r = getCurculantMatrix(i, j, k, l)
			if Determinant_n(r) == 1 {
				if min > k+l {
					min = k + l
					kmin = k
					lmin = l
					found = true
				}
			}
		}
	}

	if found {
		r = getCurculantMatrix(i, j, kmin, lmin)
		//s := fmt.Sprintf("M=%d ==> M^-1=%d\n", r, MatrixInverse(r, "ev"))
		return r,MatrixInverse(r,"ev"),false
	} else {
		return nil,nil,true
	}
}

func PrintHexTable(b []string) string {
	a := fmt.Sprintf("%s", b)
	a = strings.Replace(a, " ", ",", -1)
	a = strings.Replace(a, "] ", "}", -1)
	a = strings.Replace(a, "[", "{\n", -1)
	return a
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
			x := fmt.Sprintf("%x", c)
			if len(x) == 1 {
				x = "0" + x
			}
			a[i][j] = x
		}
	}
	return a
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

func GenerateTable_X(x int) {
	r := make([]int, 0)
	s := make([]string, 0)
	for i := 0; i < 256; i++ {
		gf := MulGF(x, i)
		r = append(r, gf.Literal_value)
		s = append(s, gf.Hex_value)
	}
	fmt.Printf("\nGenerated table for %d=0x%x =>\n %s", x, x, PrintHexTable(s))
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
