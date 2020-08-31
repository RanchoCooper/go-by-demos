package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// NOTE. 只记录第一个出错的协程错误
func main() {
	group, _ := errgroup.WithContext(context.Background())

	for i:= 0; i < 5; i++ {
		index := i + 1
		group.Go(func() error {
			fmt.Printf("start execute the %d goroutine\n", index)
			time.Sleep(time.Duration(index) * time.Second)
			if index % 2 == 0 {
				return fmt.Errorf("something has failed on goroutine: %d", index)
			}
			fmt.Printf("goroutine: %d end\n", index)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
