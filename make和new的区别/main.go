package main

import (
	"fmt"
	"reflect"
)

//make和new都是用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
//make返回的还是引用类型本身；而new返回的是指向类型的指针。
//make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据。

func main() {
	a := []int{0, 1}
	a0 := new([]int)
	a0 = &a
	fmt.Println("==========a0 := new([]int)=========")
	fmt.Printf("type:%+v \n", reflect.TypeOf(a0))
	fmt.Printf("len:%+v \n", len(*a0))
	fmt.Printf("cap:%+v \n", cap(*a0))
	fmt.Printf("%+v \n", a0)

	var a1 []int
	fmt.Println("==========var a1 []int=========")
	fmt.Printf("type:%+v \n", reflect.TypeOf(a1))
	fmt.Printf("len:%+v \n", len(a1))
	fmt.Printf("cap:%+v \n", cap(a1))
	fmt.Printf("%+v \n", a1)

	var a2 [10]int
	fmt.Println("==========var a2 [10]int=========")
	fmt.Printf("type:%+v \n", reflect.TypeOf(a2))
	fmt.Printf("len:%+v \n", len(a2))
	fmt.Printf("cap:%+v \n", cap(a2))
	fmt.Printf("%+v \n", a2)

	a3 := make([]int, 10)
	fmt.Println("==========a3 := make([]int, 10)=========")
	fmt.Printf("type:%+v \n", reflect.TypeOf(a3))
	fmt.Printf("len:%+v \n", len(a3))
	fmt.Printf("cap:%+v \n", cap(a3))
	fmt.Printf("%+v \n", a3)

	a4 := make([]int, 10, 50)
	fmt.Println("==========a4 := make([]int, 10, 50)=========")
	fmt.Printf("type:%+v \n", reflect.TypeOf(a4))
	fmt.Printf("len:%+v \n", len(a4))
	fmt.Printf("cap:%+v \n", cap(a4))
	fmt.Printf("%+v \n", a4)
}
