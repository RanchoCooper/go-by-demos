package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找到链表中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转链表的后半部分
	secondHalf := reverseList(slow.Next)
	slow.Next = nil

	// 比较链表的前半部分和反转后的后半部分
	left, right := head, secondHalf
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func main() {
	head := createList([]int{1, 2, 3, 2, 1})

	isPal := isPalindrome(head)
	fmt.Println(isPal)
}
