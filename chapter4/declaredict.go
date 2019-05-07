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
