package main

type ListNode struct {
	val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		tmpNxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNxt
	}
	return prev
}

func reverseListRec(head *ListNode) *ListNode {
}
