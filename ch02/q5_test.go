package ch02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Given a circular linked list, implement an algorithm which returns node at
   the begin- ning of the loop.

   DEFINITION: Circular linked list: A (corrupt) linked list in which a node’s
   next pointer points to an earlier node, so as to make a loop in the linked
   list.

   EXAMPLE:
   input: A -> B -> C -> D -> E -> C [the same C as earlier]
   output: C */

func FindLoopingElement(t *testing.T, head *ListElement) *ListElement {
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
	head := MakeLinkedList(1, 2, 3, 4, 5)
	ElementAt(head, 4).next = ElementAt(head, 2)
	loopElement := FindLoopingElement(t, head)
	assert.Equal(t, 5, loopElement.value)
}

func TestFindLoopingElement_ToHead(t *testing.T) {
	head := MakeLinkedList(1, 2, 3, 4, 5)
	ElementAt(head, 4).next = head
	loopElement := FindLoopingElement(t, head)
	assert.Equal(t, 5, loopElement.value)
}

func TestFindLoopingElement_Single(t *testing.T) {
	head := MakeLinkedList(1)
	head.next = head
	loopElement := FindLoopingElement(t, head)
	assert.Equal(t, 1, loopElement.value)
}
