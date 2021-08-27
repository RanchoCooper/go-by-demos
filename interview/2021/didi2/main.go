package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ADf1SDF 1
// ADf1DEE 20
// dF 30
// Ce 10
// Ce 50
//
// =>
//
// ADf1DEE 20
// ADf1SDF 1
// Ce 50
// Ce 10
// dF 30

type Strings []string

func (d Strings) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Strings) Len() int {
	return len(d)
}

func (d Strings) Less(i, j int) bool {
	s1s := strings.Split(d[i], " ")
	s2s := strings.Split(d[j], " ")

	if s1s[0] == s2s[0] {
		n1, _ := strconv.Atoi(s1s[1])
		n2, _ := strconv.Atoi(s2s[1])
		return n1 > n2
	}
	return s1s[0] < s2s[0]
}

func sortStrings(strs []string) []string {
	var ss Strings
	ss = strs
	if len(ss) == 0 || ss == nil {
		return []string{}
	}

	sort.Sort(ss)
	return ss
}

func main() {
	ss := []string {
		"ADf1SDF 1",
		"ADf1DEE 20",
		"dF 30",
		"Ce 10",
		"Ce 50",
	}
	r := sortStrings(ss)
	for _, s := range r {
		fmt.Println(s)
	}

}
