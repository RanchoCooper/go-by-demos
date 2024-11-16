package main

import "fmt"

func NearestGreater(a []int) []int {
	n := len(a)
	b := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && a[stack[len(stack)-1]] < a[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b[idx] = i
		}

		stack = append(stack, i)
	}

	for _, idx := range stack {
		b[idx] = -1
	}
	return b
}

func main() {
	a := []int{1, 4, 2, 1, 7, 6}
	b := NearestGreater(a)
	fmt.Println(b)
}
