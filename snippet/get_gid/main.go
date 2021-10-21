package main

import (
    "bytes"
    "fmt"
    "runtime"
    "strconv"
)

/**
 * @author Rancho
 * @date 2021/10/21
 */

func main() {
    fmt.Println(GetGID())
}

func GetGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}