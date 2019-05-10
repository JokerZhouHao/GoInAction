Go语言实战学习
====

## 目录

[第3章 打包和工具链](#ch3)

[第4章 数组、切片和映射](#ch4)

[第8章 标准库](#ch8)

[第9章 测试和性能](#ch9)

<span id="ch3"> </span>
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

<span id="ch4"> </span>
### 第4章 数组、切片和映射

> **数组和切片**: 数组赋值时要求两数组类型和长度必须相等，同时数组每个元素依次赋值，同样，函数或方法参数使用数组时，会复制整个数组；切片赋值只是改变下地址、len和cap，for index, value := range slice 的value不是引用的slice的值，而是复制slice的元素
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

> **map**: 如果不存在key，返回value对应类型的零值
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

<span id="ch8"> </span>
### 第8章 标准库

> **log**
``` golang
// This sample program demonstrates how to create customized loggers.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

```

> **json**
``` golang
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	InsertKeyHere string `json:"insert-key-here"`
	Key           string `json:"key"`
}

func main() {
	uri := "http://echo.jsontest.com/insert-key-here/insert-value-here/key/value"

	// 请求
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()

	// 解析json
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(result)

	// 将结构体转化为json二进制串
	pretty, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(pretty)
	fmt.Println(string(pretty))

	// 将json串转化为map
	var c map[string]interface{}
	err = json.Unmarshal(pretty, &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c["key"])

}
```

<span id="ch9"> </span>
### 第9章 测试和性能

> golang测试工具要求测试的go文件必须以`_test`结尾，对于基础单元测试的函数必须是`func TestXXX...(t *testing.T){ . . . }`结构；对于基准测试的函数必须是`func BenchmarkXXX...(b *testing.B) { . . . for i := 0; i < b.N; i++ { . . . } . . . }`结构，在测试点前要调用`b.resetTimer()`重置计时器。

**基础单元测试**: 测试当前路径下的所有Test函数：go test -v -test；测试指定函数：go test -v -test="要测试的函数名，支持正则表达式"
``` golang
// Sample test to show how to write a basic unit test.
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload validates the http Get function can download content.
func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
```

**基准测试**: go test -v -run="None" -bench="测试的函数名，支持正则"   [-benchtime="3s"默认1s]

``` golang
// Sample benchmarks to test which function is better for converting
// an integer into a string. First using the fmt.Sprintf function,
// then the strconv.FormatInt function and then strconv.Itoa.
package listing01_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf provides performance numbers for the
// fmt.Sprintf function.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
```



