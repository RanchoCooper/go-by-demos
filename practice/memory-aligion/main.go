package main

import (
	"fmt"
	"unsafe"
)

// bool 1 byte
// string 16 byte

type A struct {
	a bool		// 8byte, 补7byte
	b string	// 16byte
	c bool		// 8byte, 补7byte
}
type B struct {
	a bool
	c bool		// 8byte, 补8-1-1=6byte
	b string	// 16byte
}

func main() {
	fmt.Printf("Size A:%d\n", unsafe.Sizeof(A{}))
	fmt.Printf("Size B:%d\n", unsafe.Sizeof(B{}))
}
