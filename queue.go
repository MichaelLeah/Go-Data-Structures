// Package godatastructures provides basic implementation of various data structures such as Stacks and Queues.
package godatastructures

// Queue offers a storage mechanism where it is first in, first out
type Queue struct {
	data  []interface{} // Stores any data type
	front int           // The first entry of the Queue
	rear  int           // The last entry of the queue
}

// Insert n items into the Queue
func (q *Queue) Insert(items ...interface{}) {
	for i := range items {
		q.data = append(q.data, items[i])
	}
	q.rear = (len(q.data) - 1)
}

// Remove an item from the Queue.
func (q *Queue) Remove() {
	if len(q.data) > q.front {
		q.data[q.front] = nil
		q.front++
	}
}

// Peek returns the first item of the Queue
func (q *Queue) Peek() interface{} {
	if (len(q.data) - 1) >= q.front {
		return q.data[q.front]
	}

	return nil
}
