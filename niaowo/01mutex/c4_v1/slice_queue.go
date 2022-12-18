package c4_v1

import "sync"

type SliceQueue struct {
	sync.Mutex
	data []any
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]any, 0, n)}
}

func (q *SliceQueue) Enqueue(v any) {
	q.Lock()
	q.data = append(q.data, v)
	q.Unlock()
}

func (q *SliceQueue) Dequeue() any {
	q.Lock()
	if len(q.data) == 0 {
		q.Unlock()
		return nil
	}

	v := q.data[0]
	q.data = q.data[1:]
	q.Unlock()
	return v

}
