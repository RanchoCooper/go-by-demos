package main

import (
	"fmt"
	"math"
)

// [1,  5, 9]
// [10,11,13]
// [12,13,15]
// k = 8
// 13

func findK(nums [][]int64, k int) int64 {
	if len(nums) == 0 || nums == nil {
		return 0
	}

	ans := make([]int64, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, math.MaxInt64)
	}
	m := make(map[int64]bool, 0)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			if _, ok := m[nums[i][j]]; ok {
				continue
			} else {
				m[nums[i][j]] = true
			}
		}
	}

	for num := range m {
		if num < ans[len(ans) - 1] {
			insert(&ans, num)
		}
	}

	return ans[len(ans) - 1]
}

func insert(ans *[]int64, target int64) {
	for i := len(*ans) - 1; i >= 0; i-- {
		if (*ans)[i] > target && i != 0 {
			continue
		} else {
			(*ans)[i] = target
			break
		}
	}
}

// [1,  5, 9]
// [10,11,13]
// [12,13,15]
func main1() {
	nums := [][]int64{[]int64{1, 5, 9}, []int64{10, 11, 13}, []int64{12, 13, 15}}
	fmt.Println(findK(nums, 8))
}

// /////////////////////////////////////////////////////////////////////////////////

// [ 0, 1, 2, 3, 4]
// [24,23,22,21, 5]
// [12,13,14,15,16]
// [11,17,18,19,20]
// [10, 9, 8, 7, 6]

// FIXME
func getMaxRout(nums [][]int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	x, y := 0, 0
	max := math.MaxInt32
	for i := 0; i < 4; i++ {
		// current := nums[x][y]
		next := nums[x + dx[i]][y + dy[i]]
		if next <= max {
			max = next
		}
	}

	return max
}

// /////////////////////////////////////////////////////////////////////////////////

func getSubArray(nums1, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return []int{}
	}

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	ans := make([]int, 0)
	found := false
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] != nums2[j] {
				continue
			}

			found = true
			for nums1[i] == nums2[j] && j < len(nums2){
				ans = append(ans, nums2[j])
				j++
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(getSubArray([]int{1, 2, 3, 4, 5, 6, 8, 9}, []int{2, 8, 9}))
}
