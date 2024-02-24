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

	fmt.Printf("Current Queue: %v \n", q.elements)

	testCases := []struct {
		expected int
	}{
		{2},
		{3},
		{5},
	}

	for _, tc := range testCases {
		el, err := q.Dequeue()
		if err != nil {
			t.Errorf("unexpeted error: %v", err)
		}
		if el != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, el)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue, got %v", q.elements)
	}
}

func TestQueueFloat(t *testing.T) {
	q := NewQueue()

	q.Enqueue(2.35)
	q.Enqueue(35.57)
	q.Enqueue(5711.38)

	fmt.Printf("Current Queue: %v \n", q.elements)

	testCases := []struct {
		expected float64
	}{
		{2.35},
		{35.57},
		{5711.38},
	}

	for _, tc := range testCases {
		el, err := q.Dequeue()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if el != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, el)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue, got %v", q.elements)
	}
}

func TestQueueString(t *testing.T) {
	q := NewQueue()

	q.Enqueue("zyx110")
	q.Enqueue("abc789")
	q.Enqueue("mcu123")

	fmt.Printf("Current Queue: %v \n", q.elements)

	testCases := []struct {
		expected string
	}{
		{"zyx110"},
		{"abc789"},
		{"mcu123"},
	}

	for _, tc := range testCases {
		el, err := q.Dequeue()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if el != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, el)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("expected empty queue, got %v", q.elements)
	}
}
