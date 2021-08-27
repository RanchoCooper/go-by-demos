package main

import "math"
import "fmt"

// abbacceebaa => acceeb? bacce,ceeba 5

func deepCopy(m map[byte]bool) map[byte]bool {
    result := make(map[byte]bool)
    for k, v := range m {
        result[k] = v
    }
    return result
}


func findMinSubString(str string) int {
    // generate dict contains each element
    mapping := make(map[byte]bool)
    for _, s := range str {
        mapping[byte(s)] = true
    }

    ans := math.MaxInt32
    for i := 0; i < len(str); i++ {
        m := deepCopy(mapping)
        j := i

        for ; j < j + ans; j++ {
            if _, ok := m[str[j]]; ok {
                delete(m, str[j])
            }
            if len(m) == 0 {
                tmp := j - i + 1
                if tmp < ans {
                    ans = tmp
                }
                break
            }
            if j == len(str) -1 {
                break
            }
        }
    }
    return ans
}

func main() {
    fmt.Println(findMinSubString("abbacceebaa"))
}
