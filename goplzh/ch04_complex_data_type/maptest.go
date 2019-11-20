package main

import "fmt"

func main()  {
	//ages := make(map[string]int)
	//ages["alice"] = 31
	//ages["charlie"] = 34
	//fmt.Println(ages)

	//ages := map[string]int{
	//	"alice": 31,
	//	"charlie": 34,
	//}
	//delete(ages, "alice")
	////ages["bob"] = ages["bob"] + 1
	//ages["bob"]++
	//fmt.Println(ages)
	//
	//for name, age := range ages{
	//	fmt.Printf("%s\t%d\n", name, age)
	//}

	//ages := map[string]int{
	//	"b": 2,
	//	"a": 1,
	//}
	//
	//keys := []string{}
	//for k, _ := range ages{
	//	keys = append(keys, k)
	//}
	//
	//sort.Strings(keys)
	//
	//for _, k:= range keys{
	//	fmt.Println(k, ages[k])
	//}

	var ages map[string]int
	fmt.Println(ages==nil)
	fmt.Println(len(ages))
	ages = map[string]int{}
	ages["a"] = 1
	fmt.Println(ages)

	if age, ok := ages["kk"]; !ok{
		fmt.Println(age)
	}
 }