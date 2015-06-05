package ch02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* You have two numbers represented by a linked list, where each node contains
   a sin- gle digit. The digits are stored in reverse order, such that the 1’s
   digit is at the head of the list. Write a function that adds the two numbers
   and returns the sum as a linked list.

   EXAMPLE:
   Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
   Output: 8 -> 0 -> 8 */

func AddValues(t *testing.T, head1, head2 *ListElement) (output *ListElement) {
	// sum elements
	var sum int

	for e, factor := head1, 1; e != nil; e = e.next {
		sum += e.value * factor
		factor *= 10
	}

	for e, factor := head2, 1; e != nil; e = e.next {
		sum += e.value * factor
		factor *= 10
	}

	// build results
	var prev *ListElement

	for sum > 0 {

		e := new(ListElement)
		e.value = sum % 10

		if output == nil {
			output = e
		} else {
			prev.next = e
		}

		sum /= 10
		prev = e
	}

	return
}

func TestAddValues(t *testing.T) {
	// 513 + 295 = 808
	head1 := MakeLinkedList(3, 1, 5)
	head2 := MakeLinkedList(5, 9, 2)
	output := AddValues(t, head1, head2)
	assert.Equal(t, "8->0->8", PrintLinkedList(output))
}

func TestAddValues_DifferentLengths(t *testing.T) {
	// 1513 + 295 = 1808
	head1 := MakeLinkedList(3, 1, 5, 1)
	head2 := MakeLinkedList(5, 9, 2)
	output := AddValues(t, head1, head2)
	assert.Equal(t, "8->0->8->1", PrintLinkedList(output))
}

func TestAddValues_Nil(t *testing.T) {
	head1 := MakeLinkedList(3, 1, 5)
	head2 := (*ListElement)(nil)
	output := AddValues(t, head1, head2)
	assert.Equal(t, "3->1->5", PrintLinkedList(output))
}
