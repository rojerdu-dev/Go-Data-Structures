package main

import (
	"fmt"
	"testing"
)

func TestQueueInt(t *testing.T) {
	q := NewQueue()

	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(5)

	fmt.Println(q.elements)
	el, err := q.Dequeue()
	if err != nil || el != 2 {
		t.Errorf("expected 2, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != 3 {
		t.Errorf("expected 3, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != 5 {
		t.Errorf("expected 5, got %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue")
	}
}

func TestQueueFloat(t *testing.T) {
	q := NewQueue()

	q.Enqueue(2.3)
	q.Enqueue(33.57)
	q.Enqueue(55.11)

	el, err := q.Dequeue()
	if err != nil || el != 2.3 {
		t.Errorf("expected 2.3, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != 33.57 {
		t.Errorf("expected 33.57, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != 55.11 {
		t.Errorf("expected 55.11, got %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue")
	}
}

func TestQueueString(t *testing.T) {
	q := NewQueue()

	q.Enqueue("z")
	q.Enqueue("y")
	q.Enqueue("x")

	el, err := q.Dequeue()
	if err != nil || el != "z" {
		t.Errorf("expected z, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != "y" {
		t.Errorf("expected y, got %v", err)
	}

	el, err = q.Dequeue()
	if err != nil || el != "x" {
		t.Errorf("expected x, got %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue")
	}
}
