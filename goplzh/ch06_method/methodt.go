package main

import (
	"image/color"
	"math"
	"sync"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list==nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

//type ColoredPoint struct {
//	Point
//	Color color.RGBA
//}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string  {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	// 方法声明
	//var p= Point{X: 1, Y: 2}
	//fmt.Println(p)

	// 基于指针对象的方法
	//var p1= Point{X: 2, Y: 2}
	//fmt.Println(p.Distance(p1))
	//Point{1,2}.Distance(p1)

	// Nil也是合法的接收器类型

	// 通过嵌入结构体来扩展类型
	//var cp ColoredPoint
	//cp.X = 1
	//fmt.Println(cp.Point.X)
	//cp.Point.Y = 2
	//fmt.Println(cp.X)
	//fmt.Println(cp.Distance(Point{1, 2}))

	//p := ColoredPoint{&Point{1, 1}, color.RGBA{1, 1,1,1}}
	//q := ColoredPoint{&Point{5, 3}, color.RGBA{1,1,1,1}}
	//fmt.Println(p.Distance(Point{2, 2}))
	//q.Point = p.Point
	//p.ScaleBy(2)
	//fmt.Println(*p.Point, *q.Point)


	// 方法值和方法表达式
	//p := Point{1,2}
	//q := Point{4, 6}
	//distanceFromP := p.Distance
	//fmt.Println(distanceFromP(q))
	//var orgin Point
	//fmt.Println(distanceFromP(orgin))
	//
	//scaleP := p.ScaleBy
	//scaleP(2)
	//scaleP(3)
	//scaleP(10)

	//p := Point{1, 2}
	//q := Point{4, 6}
	//distance := (*Point).Distance
	//fmt.Println(distance(&p, q))
	//fmt.Printf("%T\n", distance)
	//
	// 封装
}
