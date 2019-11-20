package main

import "fmt"

var a = 10

func init()  {
	a = 20
}

func main()  {
	//a := 100
	//fmt.Println(a)

	x := "hello!"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}

	biof
}
