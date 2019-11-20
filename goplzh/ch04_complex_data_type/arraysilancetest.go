package  main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main()  {
	//var a [3]int
	//fmt.Println(a[0])
	//fmt.Println(a[len(a)-1])
	//
	//for i, v := range a{
	//	fmt.Printf("%d %d\n", i, v)
	//}
	//
	//for _, v := range a{
	//	fmt.Printf("%d\n", v)
	//}

	//var q [3]int = [3]int{1, 2, 3}
	//var r [3]int = [3]int{1, 2}
	//q := [...]int{1, 2, 3, 4}
	//fmt.Println(q, r)

	//symbol := [...]string{USD: "$", EUR: "#", GBP: "E", RMB: "EE"}
	//fmt.Println(symbol)
	//
	//r := [...]int{99: -1}
	//fmt.Println(len(r))

	//a := [2]int{1,2}
	//b := [...]int{1,2}
	//c := [2]int{1,3}
	//fmt.Println(a==b, a==c, b==c)
	//d := [3]int{1, 2}
	//fmt.Println(d)

	//a := [3]int{1, 2}
	//siA := a[:]
	//siA[0] = 2
	//fmt.Println(a)

	//var s []int
	//fmt.Println(s==nil)
	//s = []int(nil)
	//fmt.Println(s==nil)
	//s = []int{}
	//fmt.Println(s==nil)

	//s := make([]int, 10)
	//fmt.Println(s)
	//
	//var runes []rune
	//for _, r := range "Hello, 世界"{
	//	runes = append(runes, r)
	//}
	//fmt.Println("%q\n", runes)

	//var x, y []int
	//for i := 0; i < 10; i++{
	//	y = append(x, i)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(x), x)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	//	x = y
	//}

	//data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data))
	//fmt.Printf("%q\n", data)

	data := []int{1, 2, 3}
	data1 := []int{4, 5}
	copy(data[:], data1[:])
	fmt.Println(data)
}

func zero(ptr *[32]byte)  {
	*ptr = [32]byte{}
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings{
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
