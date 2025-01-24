package main

import (
	"fmt"
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	// 创建一个大小为10的协程池
	pool, err := ants.NewPool(10)
	if err != nil {
		log.Fatalf("创建协程池失败: %v", err)
	}
	defer pool.Release() // 确保程序退出时释放资源

	// 创建任务并提交到池中
	for i := 0; i < 20; i++ {
		// 任务函数
		task := func() {
			num := i
			fmt.Printf("正在处理任务 #%d\n", num)
			time.Sleep(1 * time.Second) // 模拟一些耗时操作
			fmt.Printf("任务 #%d 完成\n", num)
		}

		// 将任务提交到协程池
		if err := pool.Submit(task); err != nil {
			log.Printf("提交任务 #%d 失败: %v", i, err)
		}
	}

	// 等待所有任务完成
	time.Sleep(5 * time.Second) // 可以考虑用同步机制来代替，等所有任务完成后退出
}
