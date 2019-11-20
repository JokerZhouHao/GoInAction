package main

import "fmt"

func add(x int, y int) int  {
	return x + y
}

func sub(x, y int) (z int)  {
	z = x - y
	return
}

func first(x int, _ int) int  {
	return x
}

func zero(int, int) int  {
	return 0
}

func squares() func() int {
	var x int
	return func() int{
		x++
		return x * x
	}
}

func sum(values ...int) int {
	total := 0
	for _, val := range values{
		total += val
	}
	return total
}

func recoverTest()  {
	defer func(){
		if p := recover(); p != nil{
			fmt.Println("recover")
		}
	}()
	panic(12)
}

func main()  {
	//fmt.Printf("%T\n", add)
	//fmt.Printf("%T\n", sub)
	//fmt.Printf("%T\n", first)
	//fmt.Printf("%T\n", zero)

	//f1 := zero
	//fmt.Printf("%T\n", f1)
	//fmt.Println(f1(1, 1))

	// 匿名函数
	//f := squares()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	// 可变参数
	//fmt.Println(sum(1, 2, 3, 4))
	//arr1 := [...]int{1, 2, 3, 4}
	//slice1 := arr1[:]
	//fmt.Println(sum(slice1...))

	// recover测试
	recoverTest()
}