package stack

import "testing"

// TestPush tests the Push method of the Stack struct.
// It verifies that pushing elements onto the stack updates
// the top element and height correctly.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestPush(t *testing.T) {
	stack := Stack{}

	// Push the first element and check the top and height
	stack.Push(10)
	if stack.Top == nil {
		t.Errorf("Push() failed, top should not be nil")
		return
	}
	if stack.Top.Value != 10 {
		t.Errorf("Push() failed, expected top value to be 10, got %v", stack.Top.Value)
	}

	// Push additional elements and check the top and height
	stack.Push(20)
	stack.Push(30)
	if stack.Top.Value != 30 {
		t.Errorf("Push() failed, expected top value to be 30, got %v", stack.Top.Value)
	}

	if stack.Height != 3 {
		t.Errorf("Push() failed, expected height to be 3, got %v", stack.Height)
	}
}

// TestPop tests the Pop method of the Stack struct.
// It verifies that popping elements from the stack updates
// the top element and height correctly.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop the top element and check the value, top, and height
	value := stack.Pop()
	if value != 30 {
		t.Errorf("Pop() failed, expected value to be 30, got %v", value)
	}
	if stack.Top == nil {
		t.Errorf("Pop() failed, top should not be nil after popping one item from three")
		return
	}
	if stack.Top.Value != 20 {
		t.Errorf("Pop() failed, expected top value to be 20, got %v", stack.Top.Value)
	}

	if stack.Height != 2 {
		t.Errorf("Pop() failed, expected height to be 2, got %v", stack.Height)
	}

	// Pop the next element and check the value, top, and height
	value = stack.Pop()
	if value != 20 {
		t.Errorf("Pop() failed, expected value to be 20, got %v", value)
	}
	if stack.Top == nil {
		t.Errorf("Pop() failed, top should not be nil after popping two items from three")
		return
	}
	if stack.Top.Value != 10 {
		t.Errorf("Pop() failed, expected top value to be 10, got %v", stack.Top.Value)
	}

	// Pop the last element and check the value, top, and height
	value = stack.Pop()
	if value != 10 {
		t.Errorf("Pop() failed, expected value to be 10, got %v", value)
	}
	if stack.Top != nil {
		t.Errorf("Pop() failed, top should be nil after popping all items, got %v", stack.Top)
	}
	if stack.Height != 0 {
		t.Errorf("Pop() failed, expected height to be 0, got %v", stack.Height)
	}
}

// TestPopEmpty tests the Pop method of the Stack struct when the stack is empty.
// It verifies that popping from an empty stack returns nil and that subsequent
// push operations work correctly.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestPopEmpty(t *testing.T) {
	stack := Stack{}

	// Pop from an empty stack and check the value
	value := stack.Pop()
	if value != nil {
		t.Errorf("Pop() on empty stack should return nil, got %v", value)
	}

	// Push an element after popping from an empty stack and check the top and height
	stack.Push(42)
	if stack.Top == nil || stack.Top.Value != 42 {
		t.Errorf("Push() after Pop() on empty stack failed")
	}
	if stack.Height != 1 {
		t.Errorf("Push() after Pop() on empty stack: expected height to be 1, got %v", stack.Height)
	}
}
