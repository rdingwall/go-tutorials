package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findLoopingElement(t *testing.T, head *listElement) *listElement {
	slow := head
	fast := head
	
	if head.next == nil {
		return nil
	}
	
	for {
		t.Logf("slow=%v, fast=%v", slow.value, fast.value)
		
		if fast == slow.next {
			return fast.next.next
		}
		
		slow = slow.next
		fast = fast.next.next
	}
}

func TestFindLoopingElement(t *testing.T) {
	head := toLinkedList(1, 2, 3, 4, 5)
	elementAt(head, 4).next = elementAt(head, 2)
	loopElement := findLoopingElement(t, head)
	assert.Equal(t, 5, loopElement.value)
}

func TestFindLoopingElement_ToHead(t *testing.T) {
	head := toLinkedList(1, 2, 3, 4, 5)
	elementAt(head, 4).next = head
	loopElement := findLoopingElement(t, head)
	assert.Equal(t, 5, loopElement.value)
}

func TestFindLoopingElement_Single(t *testing.T) {
	head := toLinkedList(1)
	head.next = head
	loopElement := findLoopingElement(t, head)
	assert.Equal(t, 1, loopElement.value)
}


