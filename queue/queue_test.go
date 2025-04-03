package queue

import (
	"testing"
)

// TestEnqueue tests the Enqueue method of the Queue struct.
// It enqueues three elements into the queue and checks if the start,
// end, and height of the queue are as expected.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestEnqueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	if queue.Start.Value != 30 {
		t.Errorf("Enqueue() failed, expected start value to be 30, got %v", queue.Start.Value)
	}
	if queue.End.Value != 50 {
		t.Errorf("Enqueue() failed, expected end value to be 50, got %v", queue.End.Value)
	}
	if queue.Height != 3 {
		t.Errorf("Enqueue() failed, expected height to be 3, got %v", queue.Height)
	}
}

// TestDequeue tests the Dequeue method of the Queue struct.
// It enqueues three elements into the queue, dequeues one element,
// and checks if the value, start, end, and height of the queue are as expected.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestDequeue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	value := queue.Dequeue()
	if value != 30 {
		t.Errorf("Dequeue() failed, expected value to be 30, got %v", value)
	}
	if queue.Start.Value != 40 {
		t.Errorf("Dequeue() failed, expected start value to be 40, got %v", queue.Start.Value)
	}
	if queue.End.Value != 50 {
		t.Errorf("Dequeue() failed, expected end value to be 50, got %v", queue.End.Value)
	}
	if queue.Height != 2 {
		t.Errorf("Dequeue() failed, expected height to be 2, got %v", queue.Height)
	}
}

// TestEmptyQueue tests the Dequeue method on an empty Queue.
// It checks if the Dequeue method returns nil and if the height of the queue is 0.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestEmptyQueue(t *testing.T) {
	queue := Queue{}

	value := queue.Dequeue()
	if value != nil {
		t.Errorf("Dequeue() on empty queue failed, expected default value nil, got %v", value)
	}

	if queue.Height != 0 {
		t.Errorf("Empty queue should have height 0, got %v", queue.Height)
	}
}

// TestSingleItemQueue tests the behavior of the Queue when a single item is enqueued and then dequeued.
// It enqueues one element into the queue, checks if the start and end pointers are the same,
// dequeues the element, and verifies if the queue is empty and the height is zero.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestSingleItemQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)

	if queue.Start != queue.End {
		t.Errorf("In a single item queue, Start and End should point to the same node")
	}

	value := queue.Dequeue()
	if value != 10 {
		t.Errorf("Dequeue() failed, expected 10, got %v", value)
	}

	if queue.Start != nil || queue.End != nil {
		t.Errorf("Queue should be empty after dequeueing the only item")
	}

	if queue.Height != 0 {
		t.Errorf("Queue height should be 0 after dequeueing the only item, got %v", queue.Height)
	}
}

// TestEnqueueAfterEmptyingQueue tests the behavior of the Queue when
// an item is enqueued after the queue has been emptied.
// It enqueues one element, dequeues it to empty the queue, then enqueues
// another element and checks if the start and end pointers
// are correct and if the height of the queue is as expected.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestEnqueueAfterEmptyingQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Dequeue()

	queue.Enqueue(20)
	if queue.Start.Value != 20 || queue.End.Value != 20 {
		t.Errorf("Failed to enqueue after emptying queue")
	}
	if queue.Height != 1 {
		t.Errorf("Queue height should be 1 after enqueueing to empty queue, got %v", queue.Height)
	}
}

// TestMultipleDequeues tests the Dequeue method of the Queue struct when
// multiple elements are enqueued and then dequeued.
// It enqueues five elements into the queue, dequeues each element, and
// checks if the dequeued values are as expected.
// Finally, it verifies if the queue is empty and the height is zero.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestMultipleDequeues(t *testing.T) {
	queue := Queue{}
	values := []int{10, 20, 30, 40, 50}

	for _, v := range values {
		queue.Enqueue(v)
	}

	for i, expected := range values {
		value := queue.Dequeue()
		if value != expected {
			t.Errorf("Dequeue #%d failed, expected %v, got %v", i+1, expected, value)
		}
	}

	if queue.Height != 0 || queue.Start != nil || queue.End != nil {
		t.Errorf("Queue should be empty after dequeueing all items")
	}
}
