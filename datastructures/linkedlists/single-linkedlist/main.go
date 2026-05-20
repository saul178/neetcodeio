package main

type Node struct {
	val  int
	next *Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, nextNode *Node) *Node {
	return &Node{
		val:  val,
		next: nextNode,
	}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	head := NewNode(-1, nil)
	return &LinkedList{
		head: head,
		tail: head,
	}
}

func (ll *LinkedList) Get(index int) int {
	curr := ll.head.next
	i := 0
	for curr != nil {
		if i == index {
			return curr.val
		}
		i++
		curr = curr.next
	}
	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newNode := NewNode(val, ll.head.next)
	ll.head.next = newNode
	if newNode.next == nil {
		ll.tail = newNode
	}
}

func (ll *LinkedList) InsertTail(val int) {
	ll.tail.next = NewNode(val, nil)
	ll.tail = ll.head.next
}

func (ll *LinkedList) Remove(index int) bool {
	i := 0
	curr := ll.head
	for i < index && curr != nil {
		i++
		curr = curr.next
	}

	if curr != nil && curr.next != nil {
		if curr.next == ll.tail {
			ll.tail = curr
		}
		curr.next = curr.next.next
		return true
	}
	return false
}

func (ll *LinkedList) GetValues() []int {
	curr := ll.head.next
	res := []int{}
	for curr != nil {
		res = append(res, curr.val)
		curr = curr.next
	}
	return res
}

func reverseList(head *Node) *Node {
	var prev *Node = nil
	curr := head

	for curr != nil {
		tmpNxt := curr.next
		curr.next = prev
		prev = curr
		curr = tmpNxt
	}
	return prev
}

func reverseListRec(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := head
	if head.next != nil {
		newHead = reverseList(head.next)
		head.next.next = head
	}
	head.next = nil
	return newHead
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var dummyNode *ListNode = nil
	var tail *ListNode = dummyNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return dummyNode.Next
}

func main() {
}
