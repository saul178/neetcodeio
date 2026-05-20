package main

type DoublyLinkedNode struct {
	Val  int
	Prev *DoublyLinkedNode
	Next *DoublyLinkedNode
}

func NewDoublyLinkedNode(val int, prev *DoublyLinkedNode, next *DoublyLinkedNode) *DoublyLinkedNode {
	return &DoublyLinkedNode{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

type MyLinkedList struct {
	Head *DoublyLinkedNode
	Tail *DoublyLinkedNode
}

func Constructor() *MyLinkedList {
	// create two separate dummy nodes, this will help making insertions and deletions easier
	dummyHead := NewDoublyLinkedNode(-1, nil, nil)
	dummyTail := NewDoublyLinkedNode(-1, nil, nil)

	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead

	return &MyLinkedList{
		Head: dummyHead,
		Tail: dummyTail,
	}
}

func (ll *MyLinkedList) Get(index int) int {
	curr := ll.Head.Next
	i := 0

	for i < index && curr != nil {
		i++
		curr = curr.Next
	}
	if curr != nil && curr != ll.Tail {
		return curr.Val
	}
	return -1
}

func (ll *MyLinkedList) AddAtHead(val int) {
	newNode := NewDoublyLinkedNode(val, ll.Head, ll.Head.Next)
	newNode.Next.Prev = newNode
	ll.Head.Next = newNode
}

func (ll *MyLinkedList) AddAtTail(val int) {
	newNode := NewDoublyLinkedNode(val, ll.Tail.Prev, ll.Tail)
	newNode.Prev.Next = newNode
	ll.Tail.Prev = newNode
}

func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	curr := ll.Head.Next

	i := 0
	for i < index && curr != nil {
		i++
		curr = curr.Next
	}

	if curr == nil {
		return
	}

	newNode := NewDoublyLinkedNode(val, curr.Prev, curr)
	newNode.Prev.Next = newNode
	newNode.Next.Prev = newNode
}

func (ll *MyLinkedList) DeleteAtIndex(index int) {
	curr := ll.Head.Next
	i := 0

	for i < index && curr != nil {
		i++
		curr = curr.Next
	}

	if curr != nil && curr != ll.Tail {
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}
}

func main() {
}
