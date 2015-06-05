package ch02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Implement an algorithm to delete a node in the middle of a single linked
   list, given only access to that node.

   EXAMPLE:
   Input: the node ‘c’ from the linked list a->b->c->d->e
   Result: nothing is returned, but the new linked list looks like a->b->d->e */

func Delete(e *ListElement) {
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
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	Delete(ElementAt(head, 2))
	assert.Equal(t, "7->2->2->4->5->2->3", PrintLinkedList(head))
}

func TestDelete_Head(t *testing.T) {
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	Delete(ElementAt(head, 0))
	assert.Equal(t, "2->3->2->4->5->2->3", PrintLinkedList(head))
}

// not possible to delete at tail
