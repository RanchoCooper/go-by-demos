package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		tmp := target - num
		if idx, ok := m[tmp]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
