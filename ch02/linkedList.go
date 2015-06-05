package ch02

import (
	"bytes"
	"fmt"
)

type ListElement struct {
	value int
	next  *ListElement
}

func MakeLinkedList(values ...int) *ListElement {
	var head, tail *ListElement

	for _, value := range values {
		e := new(ListElement)
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

func ElementAt(head *ListElement, n int) *ListElement {
	e := head
	for i := 0; i < n; i++ {
		if e == nil {
			return nil
		}

		e = e.next
	}

	return e
}

func PrintLinkedList(head *ListElement) string {
	var b bytes.Buffer

	for e := head; e != nil; e = e.next {
		b.WriteString(fmt.Sprintf("%v", e.value))

		if e.next != nil {
			b.WriteString("->")
		}
	}

	return b.String()
}

func MakeSlice(head *ListElement) []int {
	values := make([]int, 0)
	for e := head; e != nil; e = e.next {
		values = append(values, e.value)
	}
	return values
}
