package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func sum(n1, n2 int) {
	defer fmt.Println("n1:", n1)
	defer fmt.Println("n2:", n2)
	res := n1 + n2
	n1++
	fmt.Println("res:", res, "n1:", n1)

}

//当执行到defer时，会把defer后面的语句压入到独立的栈中，也会讲相关的值拷贝同时入栈
//当函数执行完毕后，再从栈按先入后出的方式出栈执行语句
func main() {
	sum(10, 20)
	defer fmt.Println("main1")
	defer fmt.Println("main2")
	fmt.Println("main")
}
