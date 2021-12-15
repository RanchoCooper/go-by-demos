package util

import (
    "math/rand"
    "regexp"
    "sort"
)

/**
 * @author Rancho
 * @date 2021/12/15
 */

const (
    AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
    AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    AsciiLetters   = AsciiLowercase + AsciiUppercase
    Digits         = "0123456789"
    // HexDigits = Digits + "abcdef" + "ABCDEF"
    // OctDigits = "01234567"
    // Punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func IsDigits(s string) bool {
    if len(s) == 0 {
        return false
    }
    if s[0] == '-' || s[0] == '+' {
        s = s[1:]
    }
    reg := regexp.MustCompile(`^[0-9]+$`)
    return reg.MatchString(s)
}

func IsAllowedString(s string) bool {
    reg := regexp.MustCompile(`^[0-9a-zA-Z_-]+$`)
    return reg.MatchString(s)
}

// RandString returns random string.
func RandString(length int, isDigital bool) string {
    var char string
    if isDigital {
        char = Digits
    } else {
        char = AsciiLetters + Digits
    }
    bytes := make([]byte, length)
    for i := 0; i < length; i++ {
        bytes[i] = char[rand.Intn(len(char)-1)]
    }
    return string(bytes)
}

// RandPickWithExcept random pick n number in range range_ return inorder
func RandPickWithExcept(range_ int, n int, m []int) *[]int {
    if range_ < n || range_-len(m) <= 0 {
        return nil
    }

    result := make([]int, 0)
    except := make(map[int]bool, len(m))

    for _, v := range m {
        except[v] = true
    }
    candidate := make([]int, 0)
    for i := 0; i < range_; i++ {
        if _, ok := except[i]; !ok {
            candidate = append(candidate, i)
        }
    }

    for ; len(result) != n; {
        index := rand.Intn(len(candidate))
        result = append(result, candidate[index])
        candidate[index] = candidate[len(candidate)-1]
        candidate = candidate[:len(candidate)-1]
    }
    var intSlice sort.IntSlice
    intSlice = result
    intSlice.Sort()
    return (*[]int)(&intSlice)
}

// HasOverlap will check is there are overlap between slice n and m
func HasOverlap(n, m []int) bool {
    if n == nil || m == nil {
        return false
    }
    if len(n) == 0 || len(m) == 0 {
        return false
    }
    map_ := make(map[int]bool, len(n))
    for _, v := range n {
        map_[v] = true
    }
    for _, v := range m {
        if map_[v] {
            return true
        }
    }

    return false
}
