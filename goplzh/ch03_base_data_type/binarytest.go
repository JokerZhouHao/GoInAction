package main

import "fmt"

func main()  {
	//var x uint8 = 1<<1 | 1<<5
	//var y uint8 = 1<<1 | 1<<2
	//
	//fmt.Printf("%08b\n", x)
	//fmt.Printf("%08b\n", y)
	//
	//fmt.Printf("%08b\n", x&y)
	//fmt.Printf("%08b\n", x|y)
	//fmt.Printf("%08b\n", x^y)
	//fmt.Printf("%08b\n", x&^y)

	//var apples int32 = 1
	//var oranges int16 = 2
	//var compote = apples + int32(oranges)
	//fmt.Println(compote)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	var t = 1.2
	var t1 float64 = t
	fmt.Println(t1)

}
