package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	// str1 = "Hello"
	// str2 = " "
	// str3 = "World"

	// len = 360 + 360
	str1 = "HelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHello"
	str2 = " "
	str3 = "WorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorldWorld!"
)

func BenchmarkPlusOperator(b *testing.B) {
	// 直接使用+拼接字符串
	for i := 0; i < b.N; i++ {
		_ = str1 + str2 + str3
	}
}

func BenchmarkSprintf(b *testing.B) {
	// 使用Sprintf拼接字符串
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", str1, str2, str3)
	}
}

func BenchmarkJoin(b *testing.B) {
	// 使用Join拼接字符串
	strs := []string{str1, str2, str3}
	for i := 0; i < b.N; i++ {
		_ = strings.Join(strs, " ")
	}
}

func BenchmarkBuffer(b *testing.B) {
	// 使用Buffer拼接字符串
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.Reset()
		buffer.WriteString(str1)
		buffer.WriteString(str2)
		buffer.WriteString(str3)
		_ = buffer.String()
	}
}

// 结论
// 1. Sprint形式最耗时
// 2. Join形式次之
// 3. Buffer形式最快, 但和+相差不大
// 4. 但是在实际使用中，如果只是简单的字符串拼接，使用+是最简单的方式，性能也不错
func main() {
	result := testing.Benchmark(BenchmarkPlusOperator)
	fmt.Printf("BenchmarkPlusOperator: %s\n", result)

	result = testing.Benchmark(BenchmarkSprintf)
	fmt.Printf("BenchmarkSprintf: %s\n", result)

	result = testing.Benchmark(BenchmarkJoin)
	fmt.Printf("BenchmarkJoin: %s\n", result)

	result = testing.Benchmark(BenchmarkBuffer)
	fmt.Printf("BenchmarkBuffer: %s\n", result)
}
