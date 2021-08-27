package main

// x1, x2, x3....xn
//
// x1
// n1号线(i1-j1)
// n2号线(i2-j2)
// n3号线(i3-j3)
// ...
//
// time = t(wai-n1) t(i1-j1)+ t(j1-i2(buxing)) + t(wait-n2) +t(i2-j2) + t(j2-i3(buxing))
//
// type Node struct {
// 	Next []*Node
// 	Name string
// 	lineNumber int
// }
//
// func findRout(start, end) {
// 	queue := make([]*Node, 0)
// 	queue = append(queue, start)
// 	routs := [][]*Node
// 	m := make(map[*Node]bool)
//
// 	var bfs func(start, end, rout*)
// 	bfs = func(start, end, rout*) {
// 		for len(queue) != 0 {
// 			current := queue[0]
// 			queue = queue[1:]
// 			if m[current] {
// 				continue
// 			}
// 			rout = append(rout, current)
// 			m[current] = true
// 			if current == end {
// 				routs = append(routs, *rout)
// 			}
// 			for n := range current.Next {
// 				queue = append(queue, n)
// 			}
// 			bfs(current, end, rout)
// 			rout = make([]*Node, 0)
// 		}
// 	}
// 	rout := make([]*Node, 0)
// 	bfs(start, end, rout)
// 	return routs
// }

func main() {

}
