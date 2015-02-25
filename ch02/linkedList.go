package main

import (
	"bytes"
	"fmt"
)

type listElement struct {
	value int
	next  *listElement
}

func toLinkedList(values ...int) *listElement {
	var head, tail *listElement

	for _, value := range values {
		e := new(listElement)
		e.value = value

		if head == nil {
			head = e
		}

		if tail != nil {
			tail.next = e
		}

		tail = e
	}

	return head
}

func elementAt(head *listElement, n int) *listElement {
	e := head
	for i := 0; i < n; i++ {
		if e == nil {
			return nil
		}

		e = e.next
	}

	return e
}

func printList(head *listElement) string {
	var b bytes.Buffer

	for e := head; e != nil; e = e.next {
		b.WriteString(fmt.Sprintf("%v", e.value))

		if e.next != nil {
			b.WriteString("->")
		}
	}

	return b.String()
}

func toSlice(head *listElement) []int {
	values := make([]int, 0)
	for e := head; e != nil; e = e.next {
		values = append(values, e.value)
	}
	return values
}
