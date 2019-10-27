/**
Credit: https://github.com/eapache/queue

Note - this implementation is *not* thread safe

Original implementation updated to allow for the SQS requirement
of letting an element to go back to the front of the queue
*/
package queue

// Queue represents a first in, first out structure, backed by a self-expanding slice
type Queue struct {
	buf               []interface{}
	head, tail, count int
}

const (
	minLength = 16
)

// New constructs and returns a new queue.
func New() *Queue {
	return &Queue{
		buf: make([]interface{}, minLength),
	}
}

// Depth returns the number of elements currently stored in the queue.
func (q *Queue) Depth() int {
	return q.count
}

// resizes the queue to fit exactly twice its current contents
// this can result in shrinking if the queue is less than half-full
func (q *Queue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

// Enqueue puts an element on the end of the queue.
func (q *Queue) Enqueue(msg interface{}) {
	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = msg
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

// ReturnToQueue puts an element at the front of the queue.
func (q *Queue) ReturnToQueue(msg interface{}) {
	if q.count == len(q.buf) {
		q.resize()
	}

	q.head = (q.head - 1) & (len(q.buf) - 1)
	q.buf[q.head] = msg
	q.count++
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *Queue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *Queue) get(i int) interface{} {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		panic("queue: Get() called with index out of range")
	}
	// bitwise modulus
	return q.buf[(q.head+i)&(len(q.buf)-1)]
}

// Dequeue removes and returns the element from the front of the queue. If the
// queue is empty, the call will panic.
func (q *Queue) Dequeue() interface{} {
	if q.count <= 0 {
		panic("dequeue called on empty queue")
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--

	if len(q.buf) > minLength && (q.count<<2) == len(q.buf) {
		q.resize()
	}

	return ret
}
