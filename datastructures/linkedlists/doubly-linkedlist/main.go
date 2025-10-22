package main

type Node struct {
	Next *Node
	Prev *Node
	Val  any
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() *MyLinkedList {
	head := &Node{nil, nil, -1}
	tail := &Node{nil, nil, -1}
	head.Next = tail
	tail.Prev = head
	return &MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
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
