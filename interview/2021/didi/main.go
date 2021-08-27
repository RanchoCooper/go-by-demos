package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	m := make(map[int]int, 0)

	for index, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{index, m[another]}
		}

		m[num] = index
	}
	return []int{}
}

func threeSum(nums []int) [][]int {
	m := make(map[int]bool, 0)
	ans := make([][]int, 0)
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			another := -(nums[i] + nums[j])
			if _, ok := m[another]; ok {
				ans = append(ans, []int{nums[i], nums[j], another})
			}
			m[another] = true
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, -2, -3, -4, 1, 2, 3, 4}))
}
