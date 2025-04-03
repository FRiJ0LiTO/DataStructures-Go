# DataStructures

This project implements various data structures in Go, including a Queue, Stack, and HashMap. Each data structure provides methods for common operations and includes tests to verify their functionality.

## Project Structure

- `queue/`: Contains the implementation and tests for the Queue data structure.
- `stack/`: Contains the implementation and tests for the Stack data structure.
- `hashmap/`: Contains the implementation and tests for the HashMap data structure.
- `go.mod`: Go module file.
- `.gitignore`: Git ignore file.

## Data Structures

### Queue

A FIFO (First In, First Out) data structure.

#### Methods

- `Enqueue(value any)`: Adds a new element to the end of the queue.
- `Dequeue() any`: Removes and returns the front element of the queue.
- `PrintQueue()`: Prints all elements in the queue from front to back.

#### Tests

- `TestEnqueue(t *testing.T)`: Tests the `Enqueue` method.
- `TestDequeue(t *testing.T)`: Tests the `Dequeue` method.
- `TestEmptyQueue(t *testing.T)`: Tests the `Dequeue` method on an empty queue.
- `TestSingleItemQueue(t *testing.T)`: Tests the behavior of the queue with a single item.
- `TestEnqueueAfterEmptyingQueue(t *testing.T)`: Tests enqueuing after emptying the queue.
- `TestMultipleDequeues(t *testing.T)`: Tests multiple dequeues.

### Stack

A LIFO (Last In, First Out) data structure.

#### Methods

- `Push(value any)`: Adds a new element to the top of the stack.
- `Pop() any`: Removes and returns the top element of the stack.
- `PrintStack()`: Prints all elements in the stack from top to bottom.

#### Tests

- `TestPush(t *testing.T)`: Tests the `Push` method.
- `TestPop(t *testing.T)`: Tests the `Pop` method.
- `TestPopEmpty(t *testing.T)`: Tests the `Pop` method on an empty stack.

### HashMap

A map data structure with generic key and value types.

#### Methods

- `Put(key K, value V)`: Inserts or updates the value associated with the given key.
- `Get(key K) (V, bool)`: Retrieves the value associated with the given key.
- `Remove(key K)`: Deletes the value associated with the given key.
- `Keys() []K`: Returns a slice of all keys in the HashMap.
- `PrintMap()`: Prints all key-value pairs in the HashMap.

#### Tests

- `TestPut(t *testing.T)`: Tests the `Put` method.
- `TestGet(t *testing.T)`: Tests the `Get` method.
- `TestRemove(t *testing.T)`: Tests the `Remove` method.
- `TestKeys(t *testing.T)`: Tests the `Keys` method.
- `TestGetEmpty(t *testing.T)`: Tests the `Get` method on an empty map.
- `TestPutOverwrite(t *testing.T)`: Tests overwriting a value in the map.

## Getting Started

1. Clone the repository.

To run the tests, first navigate to the structure's folder and execute `go test`. For example:

```sh
cd stack
go test
```

You can do the same for the other structures (`queue` and `hashmap`):

```sh
cd queue
go test
```

```sh
cd hashmap
go test
```