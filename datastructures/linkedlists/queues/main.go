package main

import "fmt"

// https://medium.com/@danielabatibabatunde1/mastering-queues-in-golang-be77414abe9e
// queues are similar to stacks and their operations are
// they support fifo: first in, first out
type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{queue: []int{}}
}

// main operations:
// Enqueue push elements to the end of the queue
// Dequeue removes elements from the beginning of the stack
func (q *Queue) Enqueue(val int) {
	q.queue = append(q.queue, val)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue is empty.")
	}
	topElem := q.queue[0]
	q.queue = q.queue[1:]
	return topElem, nil
}

// sub operations:
// Peak shows the first element of the queue
func (q *Queue) Peak() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue is empty.")
	}
	return q.queue[0], nil
}

// isEmpty checks if the queue has elements or not
func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

// isFull checks if the queue is at max capacity

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/description/
// circular sandwiches = 0, square sandwiches = 1
// all students stand in a queue, each prefer either a square OR circular sandwiches
// len(students) == len(sandwiches)

// if student[i] prefers sandwich[i] then they will take the sandwich and leave the queue
// else the student goes to the back of the queue
// return the number of students that didnt eat because there perfered sandwiches wasnt there
func countStudents(students []int, sandwiches []int) int {
	failCount := 0

	for len(students) > 0 && failCount < len(students) {
		if students[0] == sandwiches[0] {
			// remove student and sandwich from the front
			students = students[1:]
			sandwiches = sandwiches[1:]
			// reset fail count since someone ate a sandwich
			failCount = 0
		} else {
			// we move the student to the end and update failcount since no sandwich was taken
			students = append(students[1:], students[0])
			failCount++
		}
	}
	return len(students)
}

// https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	stack []int
}

func NewMyStack() MyStack {
	return MyStack{stack: []int{}}
}

func (s *MyStack) Push(x int) {
	s.stack = append(s.stack, x)
	for i := 0; i < len(s.stack)-1; i++ {
		s.stack = append(s.stack, s.stack[0])
		s.stack = s.stack[1:]
	}
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		fmt.Println("stack is empty.")
		return -1
	}
	val := s.stack[0]
	s.stack = s.stack[1:]
	return val
}

func (s *MyStack) Top() int {
	if s.Empty() {
		fmt.Println("stack is empty.")
		return -1
	}
	return s.stack[0]
}

func (s *MyStack) Empty() bool {
	return len(s.stack) == 0
}
