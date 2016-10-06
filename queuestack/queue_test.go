package queuestack

import "testing"

func TestQueueIsInitiallyEmpty(t *testing.T) {
	q := NewQueue()

	if !q.Empty() {
		t.Error("newly created queue should contain no elements!")
	}
}

func TestQueueIsNonEmptyAfterEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)

	if q.Empty() {
		t.Error("queue should not be empty after enqueue-ing any elements!")
	}
}

func TestQueueIsEmptyAfterEnqueueAndDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	v, err := q.Dequeue()

	if !q.Empty() {
		t.Error("queue should be empty after enqueue-ing and dequeue-ing an element!")
	}

	if err != nil {
		t.Error("there should not have been anything erroneous occuring after enqueueing and dequeueing an element!")
	}

	if v != 1 {
		t.Error("value dequeued should be the same as what was enqueued got:", v, "want:", 1)
	}
}

func TestDequeueingAnEmptyQueueErrors(t *testing.T) {
	q := NewQueue()

	_, err := q.Dequeue()

	if err == nil {
		t.Error("queue should have errored when it was dequeued while empty!")
	}
}

func TestPeekingAnEmptyQueueErrors(t *testing.T) {
	q := NewQueue()

	_, err := q.Peek()

	if err == nil {
		t.Error("queue should have errored when it was dequeued while empty!")
	}
}

func TestPeekReturnsHeadOfQueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)

	v, err := q.Peek()

	if err != nil {
		t.Error("there should have been nothing erroneous that occured.")
	}

	if v != 1 {
		t.Error("queue should have returned the first value in line!")
	}
}

func TestQueueIsFIFOCompliant(t *testing.T) {
	q := NewQueue()

	values := []interface{}{0, 1, 2, 3}

	for _, v := range values {
		q.Enqueue(v)
	}

	for _, v := range values {
		dequeued, err := q.Dequeue()

		if err != nil {
			t.Fatal("encountered unexpected error:", err)
		}

		if v != dequeued {
			t.Error("queue is not FIFO compliant")
		}
	}
}
