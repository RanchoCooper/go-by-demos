package main

import (
	"fmt"
	"math"
)

func findShortestSubArray(longArr, shortArr []int) []int {
	shortSet := make(map[int]bool) // 记录短数组中的元素是否出现过
	shortCount := len(shortArr)    // 短数组中元素个数
	left, right := 0, 0            // 滑动窗口的左右边界
	result := []int{}              // 结果数组的起始/结束位置
	minLength := math.MaxInt

	for _, num := range shortArr {
		shortSet[num] = true
	}

	for right < len(longArr) {
		// 右指针移动，直到滑动窗口包含短数组中所有元素
		if shortSet[longArr[right]] {
			shortCount--
		}
		shortSet[longArr[right]] = false
		right++
		fmt.Printf("@debug  left: %d, right: %d\n", left, right)

		// 当滑动窗口包含短数组中所有元素时
		if shortCount == 0 {
			if right-left < minLength {
				//fmt.Printf("@debug left: %d, right: %d", left, right)
				minLength = right - left
				result = []int{left, right}
			}

			// 缩小窗口左边界，尝试找到更小的数组
			if shortSet[longArr[left]] {
				shortCount++
			}
			shortSet[longArr[left]] = true
			left += 1
		}
	}
	return longArr[result[0]:result[1]]
}

func main() {
	result := findShortestSubArray([]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, []int{1, 5, 9})
	fmt.Println("sub array is: ")
	for _, num := range result {
		fmt.Printf("%d, ", num)
	}
}
