package main

import "fmt"

func main()  {
	ch := make(chan  int, 1)
	//fmt.Println(cap(ch))
	for i := 0; i < 11; i++{
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println("send ", i)
		default:
			fmt.Println("default")
		}
	}
}