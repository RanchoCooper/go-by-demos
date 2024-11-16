package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	resultCh := make(chan int, 5)
	doneCh := make(chan struct{})

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(100)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("add %d to channel\n", num)
			resultCh <- num
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)

		sum := 0
		for i := 0; i < 5; i++ {
			num, _ := <-resultCh
			sum += num
		}

		fmt.Println("sum of random numbers: ", sum)
		doneCh <- struct{}{}
	}()

	<-doneCh
}
