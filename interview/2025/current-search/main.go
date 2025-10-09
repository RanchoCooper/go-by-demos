package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ContainsConcurrently 在 timeout 时间内并发查找 target，找到返回 true，超时或未找到返回 false。
// 不会泄漏 goroutine（会等待所有 worker 退出）。
func ContainsConcurrently(arr []int, target int, timeout time.Duration) bool {
	// 上下文：用于超时和“找到后取消”
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	n := len(arr)
	if n == 0 {
		return false
	}

	// 决定 worker 数量：以 CPU 数为基准，可根据场景调整
	numWorkers := runtime.GOMAXPROCS(0)
	if numWorkers <= 0 {
		numWorkers = 4
	}
	// 切片分块大小
	chunk := (n + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	found := make(chan struct{}, 1) // 发现时通知

	for i := 0; i < n; i += chunk {
		lo := i
		hi := i + chunk
		if hi > n {
			hi = n
		}
		wg.Add(1)
		go func(lo, hi int) {
			defer wg.Done()
			for j := lo; j < hi; j++ {
				// 优先检查取消信号，尽快退出
				select {
				case <-ctx.Done():
					return
				default:
				}
				if arr[j] == target {
					// 尝试通知（非阻塞）
					select {
					case found <- struct{}{}:
					default:
					}
					// 立即取消其他 goroutine
					cancel()
					return
				}
			}
		}(lo, hi)
	}

	// 等待结果或者超时
	select {
	case <-found:
		// 已经 cancel() 了，等待其他 worker 退出以防泄漏
		wg.Wait()
		return true
	case <-ctx.Done():
		// 可能是超时或被 cancel（若超时则为超时）
		wg.Wait()
		return false
	}
}

func main() {
	// 示例：较小数组演示
	arr := []int{5, 3, 7, 9, 2, 11, 8}
	fmt.Println(ContainsConcurrently(arr, 11, 5*time.Second)) // true
	fmt.Println(ContainsConcurrently(arr, 42, 1*time.Second)) // false
}
