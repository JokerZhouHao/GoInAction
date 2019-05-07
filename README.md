Go语言实战学习
====

## 目录


### 第3章 打包和工具链

> go build/run 带有main方法的go文件

> 格式化代码: go fmt 包名

> 简单错误检查: go vet 包名
``` golang
package chapter3

import "fmt"

func main() {
	if true {
		fmt.Printf("The", 3.14)
	}
}

code\chapter3\fmttet.go:7: Printf call has arguments but no formatting directives
```

> 文档: godoc -http=:6060 （访问localhost:6060）

> 依赖管理
*   godep
*   gb和插件vender
*   gopkg.in


### 第4章 数组、切片和映射

> `数组和切片`：数组赋值时要求两数组类型和长度必须相等，同时数组每个元素依次赋值，同样，函数或方法参数使用数组时，会复制整个数组；切片赋值只是改变下地址、len和cap，for index, value := range slice 的value不是引用的slice的值，而是复制slice的元素
``` golang
package main

import "fmt"

func main() {
	/* ------------- 数组 ---------------- */
	// 声明一个包含5个元素的整型数组
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [5]int{1: 10, 4: 40}
	// 多维数组
	arr4 := [4][2]int{1: {20, 21}, 3: {1: 41}}
	fmt.Printf("%v\n%v\n%v\n%v\n", arr1, arr2, arr3, arr4)

	/* ------------- 切片 ---------------- */
	slice1 := make([]string, 5)
	slice2 := make([]string, 5, 10)
	slice3 := []int{10, 20, 30}
	// 声明空切片
	var slice4 []int
	slice5 := make([]int, 0)
	slice6 := []int{}
	// 多维切片
	slice7 := [][]int{{10}, {100, 200}}
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", slice1, slice2, slice3, slice4, slice5, slice6, slice7)
}
```

> `dict`: 如果不存在key，返回value对应类型的零值
``` golang
package main

import "fmt"

func main() {
	// 创建dict
	var dict1 map[string]int // nil的dict在映射赋值时会报错
	dict2 := make(map[string]int)
	dict3 := map[string]int{}

	// 获取映射值
	value, exists := dict2["Blue"]
	value = dict2["Blue"] // 注意：这种情况下，如果不存在，返回value对于类型的零值
	for key, value := range dict2 {
		fmt.Printf("%v\n%v\n", key, value)
	}

	// 删除键
	delete(dict2, "Blue")

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", dict1, dict2, dict3, value, exists, value)
}
```




