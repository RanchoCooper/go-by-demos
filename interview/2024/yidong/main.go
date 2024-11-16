package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y

	return math.Sqrt(dx*dx + dy*dy)
}

func IsSquare(p1, p2, p3, p4 Point) bool {
	d1 := distance(p1, p2)
	d2 := distance(p2, p3)
	d3 := distance(p1, p4)

	if d1 == 0 || d2 == 0 || d3 == 0 {
		return false
	}

	// 判断四个边长是否相等
	if d1 == d2 && d2 == d3 {
		d4 := distance(p1, p3)
		d5 := distance(p2, p4)

		// 判断对角线是否相等
		if d4 == d5 {
			if d4 == math.Sqrt(2)*d1 {
				return true
			}
		}
	}

	return false
}

func main() {
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 0, Y: 1}
	p3 := Point{X: 1, Y: 1}
	p4 := Point{X: 1, Y: 0}

	fmt.Println(IsSquare(p1, p2, p3, p4))
}
