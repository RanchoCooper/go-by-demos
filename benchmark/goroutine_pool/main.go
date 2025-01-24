package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/panjf2000/ants/v2"
)

const PoolSize = 20

func BenchmarkAntsPool(b *testing.B) {
	pool, _ := ants.NewPool(PoolSize)
	defer pool.Release()

	task := func() {
		time.Sleep(time.Millisecond) // 模拟工作
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := pool.Submit(task); err != nil {
			b.Fatalf("提交任务失败: %v", err)
		}
	}
}

func BenchmarkWorkerpool(b *testing.B) {
	pool := workerpool.New(PoolSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Submit(func() {
			time.Sleep(time.Millisecond) // 模拟工作
		})
	}
	pool.StopWait() // 等待所有任务完成
}

func main() {
	fmt.Println("BenchmarkAntsPool", testing.Benchmark(BenchmarkAntsPool))
	fmt.Println("BenchmarkWorkerpool", testing.Benchmark(BenchmarkWorkerpool))
}
