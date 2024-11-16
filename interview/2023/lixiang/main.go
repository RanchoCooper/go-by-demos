package main

import "fmt"

// 有一个自然数序列a[n]，求最大的a[k]，满足 a[k] = a[i] + a[j]

func getMaxK(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	maxSum := 0
	m := make(map[int]struct{})
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = struct{}{}
	}

	for i := 0; i < len(arr); i++ {
		sum := arr[i]

		for j := i; j < len(arr); j++ {
			target := sum - arr[j]
			if _, ok := m[target]; ok {
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}

func main() {
	result := getMaxK([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(result)
}
