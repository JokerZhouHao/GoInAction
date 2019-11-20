package main

import "fmt"

func f() *int {
	v := 1
	return &v
}
func main(){
//	x := 1
//	p := &x
//	fmt.Println(*p)
//	fmt.Println(f())
	p := new (int)
	fmt.Println(*p)
	*p = 2
	*p *= 2
	x := 1
	y := 2
	x, y = y, x
	fmt.Println()
	fmt.Println(x, y)
	fmt.Println(*p)
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)
}
