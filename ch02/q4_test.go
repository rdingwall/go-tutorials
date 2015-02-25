package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addValues(t *testing.T, head1, head2 *listElement) *listElement {
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
	var output, prev *listElement

	for sum > 0 {

		e := new(listElement)
		e.value = sum % 10

		if output == nil {
			output = e
		} else {
			prev.next = e
		}

		sum /= 10
		prev = e
	}

	return output

}

func TestAddValues(t *testing.T) {
	// 513 + 295 = 808
	head1 := toLinkedList(3, 1, 5)
	head2 := toLinkedList(5, 9, 2)
	output := addValues(t, head1, head2)
	assert.Equal(t, "8->0->8", printList(output))
}

func TestAddValues_DifferentLengths(t *testing.T) {
	// 1513 + 295 = 1808
	head1 := toLinkedList(3, 1, 5, 1)
	head2 := toLinkedList(5, 9, 2)
	output := addValues(t, head1, head2)
	assert.Equal(t, "8->0->8->1", printList(output))
}

func TestAddValues_Nil(t *testing.T) {
	head1 := toLinkedList(3, 1, 5)
	head2 := (*listElement)(nil)
	output := addValues(t, head1, head2)
	assert.Equal(t, "3->1->5", printList(output))
}
