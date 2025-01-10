package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

var (
	bytesSlice1 = []byte("a")
	bytesSlice2 = []byte("b")
	bytesSlice3 = []byte("c")
)

type A struct {
	a1 [64]byte
}

func BenchmarkStringConversion(b *testing.B) {
	// 使用string转换
	m := map[string]int{"abc": 1}
	for i := 0; i < b.N; i++ {
		a := A{}
		for i := 0; i < 64; i++ {
			a.a1[i] = 'a'
		}
		key := string(a.a1[:])
		m[key] = 2
	}
}

func BenchmarkByteConversion(b *testing.B) {
	// 使用unsafe.Pointer转换
	m := map[string]int{"abc": 1}
	for i := 0; i < b.N; i++ {
		a := A{}
		for i := 0; i < 64; i++ {
			a.a1[i] = 'a'
		}
		tt := a.a1[:]
		key := *(*string)(unsafe.Pointer(&tt))
		m[key] = 2
	}
}

func BenchmarkSprintf(b *testing.B) {
	// 使用Sprintf拼接字符串
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", string(bytesSlice1), string(bytesSlice2), string(bytesSlice3))
	}
}

func BenchmarkBuilder(b *testing.B) {
	// 使用Builder拼接字符串
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Write(bytesSlice1)
		builder.Write(bytesSlice2)
		builder.Write(bytesSlice3)
		_ = builder.String()
	}
}

func BenchmarkBuffer(b *testing.B) {
	// 使用Buffer拼接字符串
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.Write(bytesSlice1)
		buffer.Write(bytesSlice2)
		buffer.Write(bytesSlice3)
		_ = buffer.String()
	}
}

// 结论
// 1. Sprint形式最耗时
// 2. Buffer次之
// 3. Builder形式最快, 但和+相差不大
// 4. 但是在实际使用中，如果只是简单的字符串拼接，使用+是最简单的方式，性能也不错
func main() {
	result := testing.Benchmark(BenchmarkStringConversion)
	fmt.Println("BenchmarkStringConversion:", result)

	result = testing.Benchmark(BenchmarkByteConversion)
	fmt.Println("BenchmarkByteConversion", result)

	result = testing.Benchmark(BenchmarkSprintf)
	fmt.Println("BenchmarkSprintf:", result)

	result = testing.Benchmark(BenchmarkBuilder)
	fmt.Println("BenchmarkBuilder:", result)

	result = testing.Benchmark(BenchmarkBuffer)
	fmt.Println("BenchmarkBuffer:", result)
}
