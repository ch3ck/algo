package queuestack

import "errors"

// Queue represents a collection data structure with FIFO (First In First Out) semantics).
type Queue interface {
	Enqueue(value interface{})
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	Empty() bool
}

type queue struct {
	data []interface{}
	size int
}

func (q *queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
	q.size++
}

func (q *queue) Dequeue() (interface{}, error) {
	if q.size > 0 {
		value := q.data[0]
		q.size--
		q.data = q.data[1:]

		return value, nil
	}

	return nil, errors.New("no such element")
}

func (q *queue) Peek() (interface{}, error) {
	if q.size > 0 {
		value := q.data[0]
		return value, nil
	}

	return nil, errors.New("no such element")
}

func (q *queue) Empty() bool {
	return q.size == 0
}

// NewQueue returns an empty queue.
func NewQueue() Queue {
	return &queue{}
}
