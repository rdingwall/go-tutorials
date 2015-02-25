package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func delete(e *listElement) {
	if e == nil {
		return
	}

	if e.next == nil {
		return
	}

	e.value = e.next.value
	e.next = e.next.next
}

func TestDelete(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	delete(elementAt(head, 2))
	assert.Equal(t, "7->2->2->4->5->2->3", printList(head))
}

func TestDelete_Head(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	delete(elementAt(head, 0))
	assert.Equal(t, "2->3->2->4->5->2->3", printList(head))
}

// not possible to delete at tail
