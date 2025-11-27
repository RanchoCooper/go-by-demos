package main

import (
	"fmt"
	"sync"
)

// printA 打印字符'a'，然后通知printB执行
func printA(a <-chan struct{}, b chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-a // 等待a通道信号
		fmt.Print("a")
		b <- struct{}{} // 发送信号给b通道
	}
}

// printB 打印字符'b'，然后通知printC执行
func printB(b <-chan struct{}, c chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-b // 等待b通道信号
		fmt.Print("b")
		c <- struct{}{} // 发送信号给c通道
	}
}

// printC 打印字符'c'，然后在不是最后一次的情况下通知printA执行
func printC(c <-chan struct{}, a chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-c // 等待c通道信号
		fmt.Print("c")
		// 如果不是最后一次，则发送信号给a通道
		if i < 9 {
			a <- struct{}{}
		}
	}
}

func main() {
	// 创建三个通道用于协程间通信
	// a通道使用缓冲区1是为了避免第一次启动时阻塞
	a := make(chan struct{}, 1)
	b := make(chan struct{})
	c := make(chan struct{})

	// 创建等待组以等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(3)

	// 启动三个协程分别打印a、b、c
	go printA(a, b, &wg)
	go printB(b, c, &wg)
	go printC(c, a, &wg)

	// 启动第一个协程
	a <- struct{}{}

	// 等待所有协程完成
	wg.Wait()

	// 输出换行符使结果更清晰
	fmt.Println()
}
