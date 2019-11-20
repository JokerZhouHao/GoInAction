package main

import (
	"bytes"
	"fmt"
)

func main()  {

	//s := "hello, world"
	//fmt.Println(len(s))
	//fmt.Println(s[0], s[7])
	//
	//fmt.Println(s[0:5])
	//fmt.Println(s[:5])
	//fmt.Println(s[7:])
	//fmt.Println(s[:])
	//
	//fmt.Println("goodbyt" + s[5:])

	//s := "Hello, 世界"
	//fmt.Println(len(s))
	//fmt.Println(utf8.RuneCountInString(s))

	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:])
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}

	//for i, r := range "Hello, 世界"{
	//	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//}

	//n := 0
	//for range s{
	//	n++
	//}
	//fmt.Printf("%d", n)

	//s := "プログラム"
	//fmt.Printf("% x\n", s)
	//r := []rune(s)
	//fmt.Printf("%x\n", r)
	//fmt.Println(string(r))
	//fmt.Println(string(65))
	//fmt.Println(string(0x4eac))
	//fmt.Println(string(1234567))

	//fmt.Println(comma("123456789"))

	//s := "abc"
	//b := []byte(s)
	//s2 :=string(b)
	//fmt.Println(s2)

	fmt.Println(intsToString([]int{1, 2, 3}))
}

func intsToString(values []int) string  {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range(values) {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comma(s string) string  {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func funcName() string {
	return "プログラム"
}
