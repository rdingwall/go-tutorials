package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findNthToLastElement(t *testing.T, head *listElement, n int) *listElement {
	
	nthLast := head
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

	return nthLast
}

func TestFindNthToLastElement(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	e := findNthToLastElement(t, head, 3)
	assert.Equal(t, 5, e.value)
}

func TestFindNthToLastElement_OutOfRange(t *testing.T) {
	head := toLinkedList(7, 2, 3, 2, 4, 5, 2, 3)
	e := findNthToLastElement(t, head, 999)
	
	assert.Equal(t, (*listElement)(nil), e)
}

func TestFindNthToLastElement_First(t *testing.T) {
	head := toLinkedList(7, 2, 3)
	e := findNthToLastElement(t, head, 3)
	assert.Equal(t, 7, e.value)
}

func TestFindNthToLastElement_Last(t *testing.T) {
	head := toLinkedList(7, 2, 3)
	e := findNthToLastElement(t, head, 1)
	assert.Equal(t, 3, e.value)
}