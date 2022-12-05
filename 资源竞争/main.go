package main

import (
	"fmt"
	"sync"
)

var (
	myMap map[int]int = make(map[int]int, 10)
	lock  sync.Mutex
)

func test(a int, ch chan struct{}) {
	res := 1
	for i := a; i > 0; i-- {
		res *= i
	}
	myMap[a] = res
	ch <- struct{}{}
}

func writeData(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("w:", i)
		ch <- i
	}
	close(ch)
}
func readData(ch chan int, flag chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("r:", <-ch)
	}
	flag <- true
}

func main() {
	intCh := make(chan int, 1)
	flagCh := make(chan bool, 1)
	go writeData(intCh)
	go readData(intCh, flagCh)
	<-flagCh
}
