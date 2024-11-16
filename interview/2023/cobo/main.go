package main

import "fmt"

// 实现单向链表的删除方法，要求实现时间复杂度为0(n)和0(1)的两种算法，空间复杂度也要求为0(1)，代码需要最终可运行

type ListNode struct {
	Value int
	Next  *ListNode
}

func DeleteNode1(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	next := node.Next
	node.Value = next.Value
	node.Next = next.Next
}

func DeleteNodeN(head *ListNode, value int) *ListNode {
	if head != nil && head.Value == value {
		// head is target
		return head.Next
	}

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}
	node4 := &ListNode{Value: 4}
	node5 := &ListNode{Value: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	// delete value equals 3
	head = DeleteNodeN(head, 3)
	printList(head)

	DeleteNode1(node4)
	printList(head)
}
