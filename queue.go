package main

import "errors"

type Queue struct {
	elements []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(el any) {
	q.elements = append(q.elements, el)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	el := q.elements[0]
	q.elements = q.elements[1:]
	return el, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}
