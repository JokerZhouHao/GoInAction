package main

import "fmt"

/*
第四章：数组、切片、map、结构体
 */

// 数组
func array()  {
	fmt.Println("数组 : ")
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]int;
	arr3 := [...]int{1, 2, 3}
	arr4 := [...]int{1:2, 4:5}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}

// 切片
func slice()  {
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("切片 : ")
	slice1 := arr1[0:1]
	str1 := "zhouhao"
	slice2 := str1[1:4]
	bytes := [...]byte{0, 1, 2, 3}
	slice3 := bytes[1:3]
	slice3[1] = 9
	str2 := "tolive"
	slice4 := str2[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(bytes)
	fmt.Println(slice4)
}

// mapp
func mapp()  {
	fmt.Println("map")
	var m1 map[string]int
	fmt.Println(m1 == nil)
	fmt.Println(len(m1))

	m2 := map[string]int{
		"a" : 1,
		"b" : 2,
	}
	for k, v := range m2{
		fmt.Println(k, v)
	}
	fmt.Println(m2)
	
}

// 结构体
type Student struct {
	name	string
	age		int
}

func getStu(stu Student) Student {
	return stu
}

func structt()  {

	//var stu = new Student{"zhou", 25}
	stu1 := Student{
		"zhou",
		25,
	}
	var stu2 = Student{"hao", 29}
	fmt.Println(stu1)
	stu2.age = 50
	fmt.Println(stu2)

	stu3 := getStu(stu2)
	stu3.age = 1000
	fmt.Println(stu3)

	var stu4 = new(Student)
	fmt.Println(stu4 == nil)

}

func main()  {
	//slice()
	//mapp()
	structt()
}
