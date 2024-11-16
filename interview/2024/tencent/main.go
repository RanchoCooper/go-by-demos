package main

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseListNode(head *ListNode) *ListNode {
	prev := &ListNode{}
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	// head := &ListNode{Value: 1}
	// head.Next = &ListNode{Value: 2}
	// head.Next.Next = &ListNode{Value: 3}
	// head.Next.Next.Next = &ListNode{Value: 4}
	// head.Next.Next.Next.Next = &ListNode{Value: 5}
	//
	// newHead := reverseListNode(head)
	//
	// for newHead != nil {
	// 	fmt.Println(newHead.Value)
	// 	newHead = newHead.Next
	// }

	wg := sync.WaitGroup{}
	n := 3
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanA
			fmt.Println("A")
			chanB <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanB
			fmt.Println("B")
			chanC <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanC
			fmt.Println("C")
			if i != n-1 {
				chanA <- struct{}{}
			}
		}
	}()

	chanA <- struct{}{}

	wg.Wait()
}
