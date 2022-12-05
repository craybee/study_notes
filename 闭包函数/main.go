package main

import "fmt"

func add() func(x int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func create() func() int {
	c := 2
	return func() int {
		c += 1
		return c
	}
}

//返回闭包时并不是单纯返回一个函数，而是返回了一个结构体，记录下函数返回地址和引用的环境中的变量地址。
//type Closure struct {
//	F func()()
//	i *int
//}

func main() {
	a := add()
	b := add()
	fmt.Println("", a(1))
	fmt.Println("", a(2))
	fmt.Println("", b(1))
	fmt.Println("", b(3))

	f1 := create()
	f2 := create()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f2())

	//a := new(int)
	//f := squares()
	//fmt.Println(f(1)) // "1"
	//fmt.Println(f(3)) // "4"
	//fmt.Println(f(3)) // "9"
	//fmt.Println(f(4)) // "16"
}
