package queue

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

type Queue struct {
	Start, End *Node
	Height     int
}

// Enqueue adds a new element to the end of the queue.
// It creates a new Node with the given value and sets it as the new end of the queue.
// If the queue is empty, the new node is set as both the start and end of the queue.
// The height of the queue is incremented by one.
//
// Parameters:
//
//	value (any): The value to be added to the queue.
func (queue *Queue) Enqueue(value any) {
	newNode := &Node{value, nil}
	if queue.Height == 0 {
		queue.Start = newNode
		queue.End = newNode
	} else {
		queue.End.Next = newNode
		queue.End = newNode
	}
	queue.Height++
}

// Dequeue removes and returns the front element of the queue.
// If the queue is empty, it returns nil.
// The height of the queue is decremented by one.
//
// Returns:
//
//	any: The value of the front element of the queue, or nil if the queue is empty.
func (queue *Queue) Dequeue() any {
	if queue.Height == 0 {
		return nil
	}
	temp := queue.Start.Value
	if queue.Height == 1 {
		queue.Start = nil
		queue.End = nil
	} else {
		queue.Start = queue.Start.Next
	}
	queue.Height--
	return temp
}

// PrintQueue prints all elements in the queue from front to back.
// It iterates through the nodes starting from the front and prints each node's value.
func (queue *Queue) PrintQueue() {
	temp := queue.Start
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}
