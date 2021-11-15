package main

import (
    "fmt"
    "time"

    "go.uber.org/ratelimit"
)

/**
 * @author Rancho
 * @date 2021/11/15
 */

func main() {
    rl := ratelimit.New(100)
    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }
}
