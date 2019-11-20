package main

import "fmt"

type Employee struct{
	ID	int
	Name,Address string
	Salary int
}

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	//p := Point{1, 2}
	//p1 := Point{X:3, Y:4}
	//fmt.Println(p)
	//fmt.Println(p1)

	//pp := &Point{1, 2}
	//pp1 := new(Point)
	//fmt.Println(pp, pp1)

	//hits := make(map[address](int))
	//hits[address{"golang.org", 443}]++
	//fmt.Println(hits)

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
}




