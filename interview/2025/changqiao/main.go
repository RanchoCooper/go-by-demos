package main

import (
	"fmt"
	"sync"
)

func main() {
	numCh := make(chan int)
	letterCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// 数字 goroutine：接收索引，打印数字，然后把索引发给 letter
	go func() {
		defer wg.Done()
		for {
			n, ok := <-numCh
			if !ok {
				// numCh 被关闭，说明整个流程结束，关闭 letterCh 并退出
				close(letterCh)
				return
			}
			fmt.Print(n)
			// 把收到的索引传给 letter 去打印对应字母
			letterCh <- n
		}
	}()

	// 字母 goroutine：接收索引，打印字母，决定是否继续发送下一个索引或结束
	go func() {
		defer wg.Done()
		for {
			n, ok := <-letterCh
			if !ok {
				// letterCh 被关闭，流程结束
				return
			}
			// 打印对应大写字母
			fmt.Printf("%c", 'A'+n-1)
			if n < 26 {
				// 发送下一个索引回 num goroutine
				numCh <- n + 1
			} else {
				// 打到 26 (Z)，触发关闭 numCh 来让 num goroutine 退出并最终关闭 letterCh
				close(numCh)
				// 注意：不要这里关闭 letterCh —— 由 num goroutine 关闭 letterCh，以保证关闭顺序安全
			}
		}
	}()

	// 启动：从 1 开始
	numCh <- 1

	// 等待两个 goroutine 结束
	wg.Wait()
	// 换行以便输出整洁
	fmt.Println()
}
