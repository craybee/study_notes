package main

import (
	"fmt"
	"unsafe"
)

/*
	1关闭channel后，再对该channel进行发送操作都将导致panic异常。对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据。
	2试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制
	3传统的方法遍历channel后不关闭，会阻塞导致deadlock
	4range会直到channel被close才会停止，停止前会阻塞后面的语句
*/

func unbuffered() {
	ch := make(chan int)
	//会导致死锁，因为我们无缓冲channel读写是同步的，赋值完成之后来不及读取channel，程序就已经阻塞了。对于有缓冲的channel，可以这样写，因为发送方会一直阻塞直到数据被拷贝到缓冲区；如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复。
	//ch <- 1
	//go func() {
	//	fmt.Println("", <-ch)
	//}()

	//正确写法
	go func() {
		fmt.Println("", <-ch)
	}()
	ch <- 1
}

func main() {
	//unbuffered()

	//空结构体的特点:1、不占用内存；2、地址不变
	var s struct{}
	var s1 struct{}
	fmt.Println("空结构体占用内存的情况：", unsafe.Sizeof(s))
	fmt.Printf("空结构体指针指向情况:s = %p, s1 = %p,两个指针的比较结果：%v\n", &s, &s1, &s == &s1)

	ch := make(chan int32)
	exit := make(chan struct{})
	defer close(ch)
	go func() {
		for i := 'a'; i < 'z'; i++ {
			fmt.Println("", string(<-ch))
			i++
			ch <- i
		}
	}()

	go func() {
		for i := 'a'; i < 'z'; i++ {
			ch <- i
			fmt.Println("", string(<-ch))
			i++
		}
		exit <- struct{}{}
	}()
	<-exit
	//ch := make(chan int)
	//exit := make(chan struct{})
	//
	//defer close(ch)
	////defer close(exit)
	//go func() {
	//	for i := 1; i <= 40; i++ {
	//		println("g1:", <-ch) // 执行步骤1， 执行步骤5
	//		i++                  //执行步骤6
	//		ch <- i              // 执行步骤7
	//	}
	//}()
	//
	//go func() {
	//	for i := 0; i < 40; i++ {
	//		i++                  // 执行步骤2
	//		ch <- i              //执行步骤3
	//		println("g2:", <-ch) //执行步骤4
	//	}
	//	exit <- struct{}{}
	//}()
	//
	//<-exit //直到上面两个协程完成退出
}
