package main

import (
	"fmt"
	"strings"
	"unicode"
)

// strings 包学习
func Println(data ...interface{})  {
	fmt.Println(data)
}

func main()  {
	str1, str2 := "zzzou", "hao"

	println(strings.Compare(str1, str2))

	println(strings.Contains(str1, "sh"))

	println(strings.ContainsAny(str1, "ot"))

	println(strings.Count(str1, "zz"))

	strs := []string{"a", "b", "c"}
	println(strings.Join(strs, "#"))

	f := func(c rune) bool {
		fmt.Println(1)
		return unicode.Is(unicode.Han, c)
	}
	println(strings.IndexFunc("Hello, 世界", f))


}



