package queue

import "errors"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return -1, errors.New("Queue is empty. Nothing to dequeue")
	}

	firstItem := q.items[0]

	q.items = q.items[1:]

	return firstItem, nil
}
