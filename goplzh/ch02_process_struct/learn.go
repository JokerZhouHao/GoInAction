package main

/*
第二章内容：
包：
一级一级下来，当前go文件调用使用的包名为所在文件夹名
可以简单的理解为一个文件夹就是一个包，这个包的功能由该文件下的所有go文件组成
包初始化使用func init(){ . . . }

变量和方法名命名方法：和java一样，采用驼峰命名法

声明关键词：var、const、type、func

声明变量：
var 变量名 类型 = 表达式（类型和表达式至少一个）
x1, x2, .... , xn := y1, y2, . . ., yn x1---xn中至少有一个未声明过

元组赋值：
i, j = j, i
_, j = j, i  左边忽略j

指针:
&或new(int)

 */


import "fmt"

var hao int64 = 2000
const zhou  = 100
type INT int64

func init()  {
	hao = 10000
}

func add(a int64, b int64) int64 {
	return a + b
}

func sub(a int32, b int32) int32 {
	return a - b

}

func main()  {
	//pHao := &hao
	//pp := new(int)
	//*pp = 3000
	//fmt.Println(*pp)
	//fmt.Println(*pHao)
	//fmt.Println(add(hao, zhou))
	fmt.Println(sub(10,
			20,
		))
}






