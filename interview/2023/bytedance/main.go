package main

import "fmt"

// 输入数组：[1, 3, 6, 15, 15, 4, 4, 3]
// 输出数组元素去重后个数：5
// 要求：空间复杂度O(1)，不改变原数组
func getUniqueCount(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums {
		m[num] = struct{}{}
	}

	return len(m)
}

func main() {
	result := getUniqueCount([]int{1, 3, 6, 15, 15, 4, 4, 3})
	fmt.Println("result: ", result)
}
