package main

import "fmt"

func counter(out chan<-int)  {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func main()  {
	//naturals := make(chan int)
	//go counter(naturals)

	ch := make(chan int, 10)
	fmt.Println(cap(ch))
	ch <- 10
	fmt.Println(len(ch))
}
