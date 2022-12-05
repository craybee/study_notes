package main

import (
	"fmt"
)

var ch chan int

func elegance() {
	v := <-ch
	fmt.Println("the ch value receive", v)
}

func main() {
	ch = make(chan int, 2)
	for i := 0; i < 5; i++ {
		ch <- i
		go elegance()
		fmt.Println("the result i", i)
	}
	close(ch)
}
