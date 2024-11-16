package main

import (
	"fmt"
	"sort"
)

func findLongestWord(arr []string) string {
	// 排序
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j])
	})

	// 遍历排序后的数组
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		found := false

		// 分解当前单词
		for j := 1; j < len(word); j++ {
			part1 := word[:j]
			part2 := word[j:]

			if contains(arr, part1) && contains(arr, part2) {
				found = true
				break
			}
		}

		if found {
			return word
		}
	}
	return ""
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func main() {
	arr := []string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}
	longestWord := findLongestWord(arr)
	fmt.Println("Longest word: ", longestWord)
}
