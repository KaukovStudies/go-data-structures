package queue

import "errors"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return -1, errors.New("queue is empty - nothing to dequeue")
	}

	firstItem := q.items[0]

	q.items = q.items[1:]

	return firstItem, nil
}
