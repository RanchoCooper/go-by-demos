package main

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Value int
}

func reverseListNode(head *Node) (result *Node) {
	var prev *Node
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func merge(l1, l2 *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

func split(head *Node) (*Node, *Node) {
	slow, fast := head, head

	var prev *Node

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	return head, slow
}

func sort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	left = sort(left)
	right = sort(right)

	return merge(left, right)
}

func printListNode(head *Node) {
	for head != nil {
		fmt.Printf("%d\n", head.Value)
		head = head.Next
	}
}

func main() {
	head := &Node{Value: 4}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 1}

	sorted := sort(head)
	printListNode(sorted)
}
