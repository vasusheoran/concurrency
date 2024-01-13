package boundedblockingqueue

import "sync"

type BoundedBlockingQueue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
}

type queue struct {
	m    sync.Mutex
	c    sync.Cond
	data []interface{}
	cap  int
}

func (q *queue) Enqueue(item interface{}) {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.data) == q.cap {
		q.c.Wait()
	}

	q.data = append(q.data, item)
	q.c.Signal()
}

func (q *queue) Dequeue() interface{} {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.data) == 0 {
		q.c.Wait()
	}

	result := q.data[0]
	q.data = q.data[1:len(q.data)]
	q.c.Signal()
	return result
}

func New(capacity int) BoundedBlockingQueue {
	q := new(queue)
	q.c = sync.Cond{
		L: &q.m,
	}
	q.cap = capacity
	q.data = []interface{}{}
	return q
}
