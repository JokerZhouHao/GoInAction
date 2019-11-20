package main

import (
	"fmt"
	"time"
)

func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++{
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// squarter
	go func() {
		for x := range naturals{
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutines)
	for x := range squares{
		fmt.Println(x)
	}
}
