package encryption

import (
	"fmt"
	"math"
)

var polyMinFields []int = []int{8, 4, 3, 1, 0}

type GField struct {
	Fields        []int
	Hex_value     string
	Literal_value int
	Byte_value byte
	HexString     string
}

func (gf *GField) GField_new(li int) {
	gf.Literal_value = li
	gf.Fields = GetFields(li)
	gf.HexString = getHexString(li)
	gf.Hex_value = getHex(li)
	gf.Byte_value = byte(li)
}

func getHex(i int) string {
	r := fmt.Sprintf("%x", i)
	if len(r) == 1 {
		r = "0x0" + r
	} else {
		r = "0x" + r
	}
	return r
}

func getHexString(i int) string {
	r := fmt.Sprintf("%x", i)
	if len(r) == 1 {
		r = "0" + r
	}
	return r
}

func AddGF(a int, b int) *GField {
	var gf *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf.GField_new(a)
	gf.addGField(b)
	gf.computeLiteralValues()
	return gf
}

func DivGF(a int, b int) *GField {
	var gf *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf.GField_new(a)
	gf.divGField(b)

	return gf
}

func HexToInt(str string) int{
	if(len(str)==1){
		return strToInt(str)
	} else {
		bx := strToInt(str[:1])
		by := strToInt(str[1:2])
		return bx * 16 + by
	}
}

func strToInt(hex string) int{
	s := 0
	switch hex {
	case "0":
		{
			s = 0
			break
		}
	case "1":
		{
			s = 1
			break
		}
	case "2":
		{
			s = 2
			break
		}
	case "3":
		{
			s = 3
			break
		}
	case "4":
		{
			s = 4
			break
		}
	case "5":
		{
			s = 5
			break
		}
	case "6":
		{
			s = 6
			break
		}
	case "7":
		{
			s = 7
			break
		}
	case "8":
		{
			s = 8
			break
		}
	case "9":
		{
			s = 9
			break
		}
	case "a":
		{
			s = 10
			break
		}
	case "b":
		{
			s = 11
			break
		}
	case "c":
		{
			s = 12
			break
		}
	case "d":
		{
			s = 13
			break
		}
	case "e":
		{
			s = 14
			break
		}
	case "f":
		{
			s = 15
			break
		}
	}
  return s
}

func (gf *GField) addGField(b int) {
	var gf_i *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf_i.GField_new(b)
	if(gf_i.Literal_value == 0){
		return
	} else {
		gf.add(gf_i)
	}
}

func (gf *GField) add(gf_i *GField){
	nj := len(gf_i.Fields)
	b := gf_i.Fields
	for i := 0 ; i< nj;  i++{
		if gf.haveField(b[i]) {
			gf.removeField(b[i])
		} else {
			gf.Fields = append(gf.Fields, b[i])
		}
	}
	gf.computeLiteralValues()
	gf.reduceResultbyPolyMin()
}

func (gf *GField) divGField(b int) {
	var gf_i *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf_i.GField_new(b)

	if gf.Literal_value == 0 {
		return
	} else if gf.maxField() < gf_i.maxField() {
		fmt.Printf("The number %s is greater than %s, can't be devided\n", gf_i.Hex_value, gf.Hex_value)
	} else {
		gf.DevideBy(gf_i)
	}
}

func (gf *GField) DevideBy(gf_i *GField) {
	res := make([]int, 0)
	r := gf.maxField() - gf_i.maxField()
	var s *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	s.GField_new(1)
	s.Fields = gf.Fields
	j := 1
	for r >= 0 && len(s.Fields) > 0 {
		b := gf_i.multiX(r)
		n := len(b)
		res = append(res, r)
		for i := 0; i < n; i++ {
			if s.haveField(b[i]) {
				s.removeField(b[i])
			} else {
				s.Fields = append(s.Fields, b[i])
			}
		}
		//fmt.Printf("\niteration %d: res = %d, b =%d, s=%d, r=%d", j, res, b, s.Fields, r)
		r = s.maxField() - gf_i.maxField()
		j++
	}

	if len(s.Fields) != 0 {
		fmt.Printf("\nthe number %s can't be devided by %s\n", gf.Hex_value, gf_i.Hex_value)
	} else {
		gf.Fields = res
		gf.computeLiteralValues()
	}
}

func MulGF(a int, b int) *GField {
	var gf *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf.GField_new(a)
	gf.MulGField(b)

	return gf
}

func GetFields(li int) []int {
	var t int = li
	var r []int
	if t >= 128 {
		r = append(r, 7)
		t -= 128
	}
	if t >= 64 {
		r = append(r, 6)
		t -= 64
	}
	if t >= 32 {
		r = append(r, 5)
		t -= 32
	}
	if t >= 16 {
		r = append(r, 4)
		t -= 16
	}
	if t >= 8 {
		r = append(r, 3)
		t -= 8
	}
	if t >= 4 {
		r = append(r, 2)
		t -= 4
	}
	if t >= 2 {
		r = append(r, 1)
		t -= 2
	}
	if t == 1 {
		r = append(r, 0)
		t -= 1
	}

	return r
}

func (gf *GField) MulGField(i int) {
	var gf_i *GField = &GField{[]int{}, "0x00", 0, 0x00, "00"}
	gf_i.GField_new(i)
	r := &GField{[]int{}, "0x00", 0, 0x00, "00"}
	//fmt.Printf("r=%d\n", r.Fields)
	ni := len(gf_i.Fields)
	nj := len(gf.Fields)
	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			m := gf.Fields[j] + gf_i.Fields[i]
			//fmt.Printf("case :%d, r=%d, haveField=%t\n", m, r.Fields, r.haveField(m))
			if r.haveField(m) {
				r.removeField(m)
			} else {
				r.Fields = append(r.Fields, m)
			}
		}

	}
	r.reduceResultbyPolyMin()
	r.computeLiteralValues()
	gf.copyField(r)
}

func (gf *GField) reduceResultbyPolyMin() {
	if gf.maxField() >= 8 {
		r := gf.maxField() - 8
		b := getPowerPolyMin(r)
		for i := 0; i < 5; i++ {
			//fmt.Printf("case :%d, r=%d, haveField=%t\n", b[i], gf.Fields, gf.haveField(b[i]))
			if gf.haveField(b[i]) {
				gf.removeField(b[i])
			} else {
				gf.Fields = append(gf.Fields, b[i])
			}
		}
		gf.computeLiteralValues()
		gf.reduceResultbyPolyMin()
	}
}

func getPowerPolyMin(x int) []int {
	var r []int = make([]int, 0)
	for i := 0; i < 5; i++ {
		r = append(r, x+polyMinFields[i])
	}
	return r
}

func (gf *GField) multiX(x int) []int {
	var r []int = make([]int, 0)
	n := len(gf.Fields)
	for i := 0; i < n; i++ {
		r = append(r, x+gf.Fields[i])
	}
	return r
}

func (gf *GField) maxField() int {
	n := len(gf.Fields)
	var r int = 0
	for i := 0; i < n; i++ {
		if r < gf.Fields[i] {
			r = gf.Fields[i]
		}
	}
	return r
}

func (gf *GField) copyField(f *GField) {
	gf.Fields = f.Fields
	gf.Literal_value = f.Literal_value
	gf.Byte_value = f.Byte_value
	gf.Hex_value = f.Hex_value
	gf.HexString = f.HexString
}

func (gf *GField) computeLiteralValues() {
	nj := len(gf.Fields)
	var r int = 0
	for i := 0; i < nj; i++ {
		r += int(math.Pow(2, float64(gf.Fields[i])))
	}
	gf.Literal_value = r
	gf.Byte_value = byte(r)
	gf.Hex_value = getHex(r)
	gf.HexString = getHexString(r)
}

func (gf *GField) removeField(m int) {
	nj := len(gf.Fields)
	r := make([]int, 0)
	for i := 0; i < nj; i++ {
		if gf.Fields[i] != m {
			r = append(r, gf.Fields[i])
		}
	}
	gf.Fields = r
}

func (gf *GField) haveField(m int) bool {
	nj := len(gf.Fields)
	b := false
	for i := 0; i < nj; i++ {
		if m == gf.Fields[i] {
			b = true
			break
		}
	}
	return b
}
