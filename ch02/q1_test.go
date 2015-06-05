package ch02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write code to remove duplicates from an unsorted linked list.

func RemoveDuplicates(head *ListElement) *ListElement {
	values := make(map[int]bool, 1)
	var prev *ListElement

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

func TestRemoveDuplicates(t *testing.T) {
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	head = RemoveDuplicates(head)
	assert.Equal(t, "7->2->3->4->5", PrintLinkedList(head))
}

/* FOLLOW UP: How would you solve this problem if a temporary buffer is not
   allowed? */

func RemoveDuplicates2(head *ListElement) *ListElement {

	contains := func(head, tail *ListElement, value int) bool {
		for e := head; e != nil && e != tail; e = e.next {
			if e.value == value {
				return true
			}
		}
		return false
	}

	var prev *ListElement

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

func TestRemoveDuplicates_NoBuffer(t *testing.T) {
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	head = RemoveDuplicates2(head)
	assert.Equal(t, "7->2->3->4->5", PrintLinkedList(head))
}
