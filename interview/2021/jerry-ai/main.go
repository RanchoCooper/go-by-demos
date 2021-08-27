package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	start, end int
}

type RangeList struct {
	pairs []Pair
}

func NewRangeList() *RangeList {
	return &RangeList{pairs: make([]Pair, 0)}
}

func (r *RangeList) Len() int {
	return len(r.pairs)
}

func (r *RangeList) Swap(i, j int) {
	r.pairs[i], r.pairs[j] = r.pairs[j], r.pairs[i]
}

func (r *RangeList) Less(i, j int) bool {
	return r.pairs[i].start < r.pairs[j].start && r.pairs[i].end < r.pairs[j].end
}

func (r *RangeList) Add(start, end int) {
	if start > end {
		// panic("invalid input")
		start, end = end, start
	}

	if len(r.pairs) == 0 {
		r.pairs = append(r.pairs, Pair{start: start, end: end})
		return
	}

	for _, pair := range r.pairs {
		if start == pair.start && end == pair.end {
			return
		}
		if end < pair.start || start > pair.end {
			r.pairs = append(r.pairs, Pair{start: start, end: end})
		} else if start > pair.start && end > pair.end {
			continue
		}

		// FIXME maybe something wrong
		if start == pair.start {
			if end < pair.end {
				return
			} else {
				pair.end = end
			}
		}

		if end == pair.end {
			if start < pair.start {
				pair.start = start
			} else {
				return
			}
		}
	}

	// sort pairs
	sort.Sort(r)

	// remove useless range
}

func (r *RangeList) Remove(start, end int) {
	if len(r.pairs) == 0 {
		panic("RangeList is empty")
	}
}

func (r *RangeList) Print() {
	if len(r.pairs) == 0 {
		fmt.Println("RangeList is empty")
	}
	for _, pair := range r.pairs {
		fmt.Printf("[%d, %d)\n", pair.start, pair.end)
	}
}

func main() {
	rl := NewRangeList()
	rl.Add(1, 5)
	rl.Print()
	// Should display: [1, 5)
	rl.Add(10, 20)
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add(20, 20)
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add(20, 21)
	rl.Print();
	// Should display: [1, 5) [10, 21)
	rl.Add(2, 4)
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add(3, 8)
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove(10, 10)
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove(10, 11)
	rl.Print()
	// Should display: [1, 8) [11, 21)
	rl.Remove(15, 17)
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove(3, 19)
	rl.Print()
	// Should display: [1, 3) [19, 21)
}
