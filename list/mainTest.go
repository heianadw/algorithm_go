package list

import (
	"fmt"
	"math"
)

//import "fmt"

//const (
//	var1  string = "hello"
//	var2
//	var3
//)
//
func main() {
	//var m,k int
	//m,k = 20,m
	//
	//fmt.Println(m,k)
	//fmt.Println(var1,var2,var3)
	test1("1111", 2222, "33333", 1.21)
	s := make([]byte, 0)
	//s = append(s,8)
	s = append(s, 7)
	s = append(s, 65)
	s = append(s, 79)
	fmt.Println(string(s))
	fmt.Println(math.Pow10(0))
	fmt.Println(s[0], s[1])
	fmt.Printf("%T %T\n", s[0], s[1])
	var ss []byte
	if ss == nil {
		fmt.Println("1111")
	}
	fmt.Println(len(ss), ss)
	sss := make([]byte, 4)
	sss[0] = 'A'
	fmt.Println(sss)
	var a interface{}
	fmt.Println(a)
	//fmt.Println(theKOutOfBoundsErr)
}

func test1(str string, args ...interface{}) {
	fmt.Println(str, args)
	fmt.Println(args...)
}
