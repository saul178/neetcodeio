package main

type ListNode struct {
	Val  int
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
	if head == nil {
		return nil
	}

	newHead := head
	if head.Next != nil {
		newHead = reverseListRec(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummyNode := &ListNode{}
	node := dummyNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	node.Next = list1
	if list1 == nil {
		node.Next = list2
	}
	return dummyNode.Next
}

// design a single linked list
type Node struct {
	Val  any
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() *MyLinkedList {
	return &MyLinkedList{}
}

func (mll *MyLinkedList) Get(index int) int {
}

func (mll *MyLinkedList) AddAtHead(val int) {
}

func (mll *MyLinkedList) AddAtTail(val int) {
}

func (mll *MyLinkedList) AddAtIndex(val int) {
}

func (mll *MyLinkedList) DeleteAtIndex(val int) {
}
