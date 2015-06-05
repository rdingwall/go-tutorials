package ch02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Implement an algorithm to find the nth to last element of a singly linked
   list. */

func FindNthToLastElement(t *testing.T, head *ListElement, n int) (nthLast *ListElement) {

	nthLast = head
	e := head

	// two pointers, gap between them
	for i := 0; i < n; i++ {
		if e == nil {
			return nil
		}

		e = e.next
	}

	// increment both until e hits tail
	for ; e != nil; e = e.next {
		nthLast = nthLast.next
	}

	return
}

func TestFindNthToLastElement(t *testing.T) {
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	e := FindNthToLastElement(t, head, 3)
	assert.Equal(t, 5, e.value)
}

func TestFindNthToLastElement_OutOfRange(t *testing.T) {
	head := MakeLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	e := FindNthToLastElement(t, head, 999)

	assert.Equal(t, (*ListElement)(nil), e)
}

func TestFindNthToLastElement_First(t *testing.T) {
	head := MakeLinkedList(7, 2, 3)
	e := FindNthToLastElement(t, head, 3)
	assert.Equal(t, 7, e.value)
}

func TestFindNthToLastElement_Last(t *testing.T) {
	head := MakeLinkedList(7, 2, 3)
	e := FindNthToLastElement(t, head, 1)
	assert.Equal(t, 3, e.value)
}
