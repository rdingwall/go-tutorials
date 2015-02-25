package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(head *listElement) *listElement {
	values := make(map[int]bool, 1)
	var prev *listElement

	for e := head; e != nil; {
		if values[e.value] {
			prev.next = e.next // delete node
		} else {
			values[e.value] = true
			prev = e
		}

		e = e.next
	}

	return head
}

func removeDuplicates2(head *listElement) *listElement {

	contains := func(head, tail *listElement, value int) bool {
		for e := head; e != nil && e != tail; e = e.next {
			if e.value == value {
				return true
			}
		}
		return false
	}

	var prev *listElement

	for e := head; e != nil; {
		if contains(head, e, e.value) {
			prev.next = e.next // delete node
		} else {
			prev = e
		}

		e = e.next
	}

	return head
}

func TestRemoveDuplicates(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	head = removeDuplicates(head)
	assert.Equal(t, "7->2->3->4->5", printList(head))
}

func TestRemoveDuplicates_NoBuffer(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	head = removeDuplicates2(head)
	assert.Equal(t, "7->2->3->4->5", printList(head))
}
