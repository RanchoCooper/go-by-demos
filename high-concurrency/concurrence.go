package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

type Task interface {
	Do()
}

type Worker struct {
	JobQueue chan Task
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Task)}
}

func (w Worker) Run(wq chan chan Task) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
				wg.Done()
			}
		}
	}()
}

type WorkerPool struct {
	workerlen   int
	JobQueue    chan Task
	WorkerQueue chan chan Task
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Task),
		WorkerQueue: make(chan chan Task, workerlen),
	}
}
func (wp *WorkerPool) Run() {
	fmt.Println("初始化worker")
	//初始化worker
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	// 循环获取可用的worker,往worker中写job
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

type Score struct {
	Num int
}

func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	num := 100 * 100 * 20
	// debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := NewWorkerPool(num)
	p.Run()
	
	//写入一亿条数据
	datanum := 100 * 100 * 100 * 100
	go func() {
		for i := 1; i <= datanum; i++ {
			sc := &Score{Num: i}
			p.JobQueue <- sc //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
			wg.Add(1)
		}
	}()
	wg.Wait()
	
	//循环打印输出当前进程的Goroutine 个数
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
	
}
