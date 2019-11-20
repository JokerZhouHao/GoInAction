package main

import (
	"bytes"
	"fmt"
)

func replaceSpace(str string) string {
	var buf bytes.Buffer
	for _, c := range str {
		if c != ' ' {
			buf.WriteString(string(c))
		} else {
			buf.WriteString("%20")
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(replaceSpace("k s"))
}
