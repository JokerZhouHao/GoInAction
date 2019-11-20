package main

import "fmt"

/*
第3章内容
基本数据类型
数组：int8, int16, int32, int64, uint8 ... ; float64, float64
bool
string
遍历string：
for i, r := range str{ . . . }

常量：
常量其占用空间看作至少有256bit
const x = XX
const(
	a = 12
	b	// 默认位20
	c = 20
	d
)
}

 */

const (
	AA = 12
	AAA
	BB = 33
	BBB
)


func main()  {
	var str = "周豪"
	for i, r := range str{
		fmt.Printf("%d %c\n", i, r)
	}
	fmt.Printf("%d %d %d %d\n", AA, AAA, BB, BBB)

}
