package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @author Rancho
 * @date 2022/4/4
 */

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex
	g1Counter, g2Counter := 0, 0

	// g1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				g1Counter++
				mu.Unlock()
			}
		}
	}()

	// g2
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(100 * time.Microsecond)
				mu.Lock()
				g2Counter++
				mu.Unlock()
			}
		}
	}()

	time.Sleep(2 * time.Second)
	done <- true

	fmt.Printf("g1: %d\ng2: %d\n", g1Counter, g2Counter)
}
