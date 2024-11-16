package main

import "fmt"

func countWays(n int) int {
	if n <= 1 {
		return 1
	}

	prev, curr := 1, 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

func main() {
	n := 5
	ways := countWays(n)
	fmt.Printf("对于%d阶台阶，有%d种方案", n, ways)
}
