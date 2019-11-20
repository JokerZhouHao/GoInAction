package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64  {
	return math.Sqrt(x*x + y*y)
}

func f(i, j, k int, s, t string)  {

}

func add(x int, y int) int  {
	return x + y
}

func split(x int) (int, int) {
	return x-1, x+1
}

func bareTest(x int) (a int, b int)  {
	a = x + 1
	b = x - 1
	return
}

func square(n int) int {
	return n * n
}
func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func add1(r rune) rune  {
	return r + 1
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals{
		total += val
	}
	return total
}

func ff(...int)  {

}

func g([]int)  {

}

func df()  {
	fmt.Println("deferred execute !")
}
func main()  {
	// 函数声明
	//fmt.Println(hypot(3, 4))

	//fmt.Printf("%T\n", add)

	// 多返回值
	//fmt.Println(split(3))
	// bare返回值
	//fmt.Println(bareTest(3))

	//var EOF = errors.New("EOF")
	//fmt.Println(EOF)

	// 函数字面值
	//f := square
	//fmt.Println(f(3))


	//var f func(int) int
	//fmt.Println(f)

	// 函数即行为
	//fmt.Println(strings.Map(add1, "Admix"))

	// 匿名函数，可以保存代码块中的局部变量，同时也说明，变量的生命周期并非完全在所在代码块中
	//fmt.Println(strings.Map(func(r rune) rune {
	//	return r + 1
	//}, "HAL-9000"))

	//f := squares()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	// 警告：捕获迭代变量
	//var funcs []func()
	//var its = [...]int{1, 2, 3}
	//for i, _ := range its {
	//	funcs = append(funcs, func(){
	//		fmt.Println(i)
	//	})
	//}
	//for _, f := range funcs{
	//	f()
	//}

	//可变参数
	//fmt.Println(sum(1, 2, 3, 4, 5))
	//valus := []int{1, 2, 3, 4}
	//fmt.Println(sum(valus...))
	//fmt.Printf("%T\n", ff)
	//fmt.Printf("%T\n", g)

	// deferred函数
	defer func() {
		if p:=recover(); p!=nil{
			fmt.Printf("recover execute ! ")
		}
	}()
	i := 3
	arr := [...]int{1, 2, 3}
	fmt.Println(arr[i])
	fmt.Println("after panic")
}

















