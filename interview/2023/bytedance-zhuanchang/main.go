package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 求小于n的最大数
func findMaxNumber(n int, A []int) int {
	// 将数字数组A按照从大到小排序
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	// 将n转换为字符串
	nStr := strconv.Itoa(n)

	// 将nStr转换为字节数组
	bytes := []byte(nStr)

	// 从高位开始逐位替换
	for i := 0; i < len(bytes); i++ {
		// 将当前位置的数字转换为整数
		digit, _ := strconv.Atoi(string(bytes[i]))
		// 查找A中小于当前数字的最大元素
		for j := 0; j < len(A); j++ {
			if A[j] < digit {
				// 替换当前位置的数字
				bytes[i] = byte(A[j] + '0')
				break
			}
		}
	}

	// 将字节数组转换为整数
	maxNumber, _ := strconv.Atoi(string(bytes))

	return maxNumber
}

func main() {
	// 示例输入
	n := 23121
	A := []int{2, 4, 9}

	// 求小于n的最大数
	maxNumber := findMaxNumber(n, A)

	// 输出结果
	fmt.Println("Max Number:", maxNumber)
}
