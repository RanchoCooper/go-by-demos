package main

import (
	"fmt"
	"strconv"
)

func main() {
	tmp := 0
	arr := make([]int, 0)
	for {
		n, _ := fmt.Scanln(&tmp)
		if n == 0 {
			break
		} else {
			arr = append(arr, tmp)
		}
	}
	result := 10001
	if len(arr) == 1 {
		if arr[0] >= 215 {
			fmt.Println(arr[0])
		} else {
			fmt.Println(215)
		}
		return
	} else if len(arr) == 2 {
		tmp, _ := strconv.Atoi(strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]))
		if tmp < result {
			result = tmp
		}
		tmp, _ = strconv.Atoi(strconv.Itoa(arr[1]) + strconv.Itoa(arr[0]))
		if tmp < result {
			result = tmp
		}
		if result >= 215 {
			fmt.Println(result)
		} else {
			fmt.Println(215)
		}
		return
	}
	for i := 0; i < len(arr)-2; i++ {
		for j := 1; j < len(arr)-1; j++ {
			for k := 2; k < len(arr); k++ {
				if i == j || i == k || j == k {
					continue
				}
				tmp, _ := strconv.Atoi(strconv.Itoa(arr[i]) + strconv.Itoa(arr[j]) + strconv.Itoa(k))
				if tmp < result {
					result = tmp
				}
				tmp, _ = strconv.Atoi(strconv.Itoa(arr[i]) + strconv.Itoa(arr[k]) + strconv.Itoa(j))
				if tmp < result {
					result = tmp
				}
				tmp, _ = strconv.Atoi(strconv.Itoa(arr[j]) + strconv.Itoa(arr[k]) + strconv.Itoa(i))
				if tmp < result {
					result = tmp
				}
				tmp, _ = strconv.Atoi(strconv.Itoa(arr[j]) + strconv.Itoa(arr[i]) + strconv.Itoa(k))
				if tmp < result {
					result = tmp
				}
				tmp, _ = strconv.Atoi(strconv.Itoa(arr[k]) + strconv.Itoa(arr[i]) + strconv.Itoa(j))
				if tmp < result {
					result = tmp
				}
				tmp, _ = strconv.Atoi(strconv.Itoa(arr[k]) + strconv.Itoa(arr[j]) + strconv.Itoa(i))
				if tmp < result {
					result = tmp
				}
			}
		}
	}
	fmt.Println(result)
	return
}
