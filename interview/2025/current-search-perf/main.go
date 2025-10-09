package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func ContainsConcurrentlyFast(arr []int, target int, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	n := len(arr)
	if n == 0 {
		return false
	}

	numWorkers := runtime.GOMAXPROCS(0)
	if numWorkers <= 0 {
		numWorkers = 4
	}
	chunk := (n + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	var foundFlag uint32 // 0 未找到，1 已找到

	// 每个 worker 每隔 checkInterval 次检查 ctx.Done()/foundFlag
	const checkInterval = 64

	for i := 0; i < n; i += chunk {
		lo := i
		hi := i + chunk
		if hi > n {
			hi = n
		}
		wg.Add(1)
		go func(lo, hi int) {
			defer wg.Done()
			counter := 0
			for j := lo; j < hi; j++ {
				// 先快速检查原子标志（非常便宜）
				if atomic.LoadUint32(&foundFlag) == 1 {
					return
				}

				if arr[j] == target {
					// 标记并取消上下文
					if atomic.CompareAndSwapUint32(&foundFlag, 0, 1) {
						cancel()
					}
					return
				}

				counter++
				if counter >= checkInterval {
					counter = 0
					// 间隔性检查 ctx，这里做非阻塞检查
					select {
					case <-ctx.Done():
						return
					default:
					}
				}
			}
		}(lo, hi)
	}

	// 等待结果或超时
	<-ctx.Done() // ctx 被 cancel（找到或超时）时返回
	// 如果是找到，foundFlag == 1；如果超时则可能为 0
	wg.Wait()
	return atomic.LoadUint32(&foundFlag) == 1
}

func main() {
	// 示例（小数组）
	arr := []int{5, 3, 7, 9, 2, 11, 8}
	fmt.Println(ContainsConcurrentlyFast(arr, 11, 5*time.Second)) // true
	fmt.Println(ContainsConcurrentlyFast(arr, 42, 1*time.Second)) // false
}
