// package main
//
// import "fmt"
//
// func main() {
// 	s1 := []int{0, 1, 2, 3}
// 	s2 := s1[1:3]
// 	fmt.Printf("s2: %v, cap(s2): %d\n", s2, cap(s2)) // 1
// 	// s2: [1,2], cap: 4
//
// 	s1[1] = 10
// 	fmt.Printf("s1: %v, s2: %v\n", s1, s2) // 2
// 	// s1: [0, 10, 2, 3], s2: [10, 2]
//
// 	s2 = append(s2, 100)
// 	s2 = append(s2, 200)
// 	fmt.Printf("s1: %v, s2: %v\n", s1, s2) // 3
// 	// s1:[0, 10, 2], s2: [1, 2, 100, 200]
//
// 	s3 := make([]int, 0, 4)
// 	s3 = append(s3, 0, 1, 2)
// 	s4 := append100(s3)
// 	fmt.Printf("s3: %v, s4: %v\n", s3, s4) // 4
// 	// s3: [0, 1, 2], s4: [0, 1, 2, 100]
// }
//
// func append100(s []int) []int {
// 	s = append(s, 100)
// 	return s
// }

// package main
//
// import "fmt"
//
// func main() {
// 	var m map[int]int
// 	fmt.Println(m[0]) // 1
// 	m[0] = 1
// 	fmt.Println(m[0]) // 2
// }

// 编写一个 Go 程序来完成以下任务:
// 1. 启动5个goroutine，每个goroutine执行以下操作:
// - 生成一个随机整数，发送到 channel 中
// - goroutine 休眠 200ms 后结束
// 2. 启动一个单独的 goroutine，从 channel 中接收这些随机整数，将它们求和后，打印最终的和。
// 3. main 中等待所有 goroutine 完成。
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			num := rand.Intn(100)
			ch <- num
			fmt.Println("send ", num)
			time.Sleep(200 * time.Microsecond)
		}()
	}

	sum := 0

	go func() {
		for num := range ch {
			sum += num
		}
	}()

	wg.Wait()
	close(ch)

	fmt.Println("sum of rand numbers is ", sum)
}
