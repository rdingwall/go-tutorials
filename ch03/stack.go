package main

import (
	"bytes"
	"fmt"
)

type Stack struct {
	top *stackElement
}

type stackElement struct {
	value interface{}
	next *stackElement
}

func (s *Stack) Pop() interface{} {
	if stack.top == nil {
		return nil
	}
	
	value := stack.top.value
	
	stack.top = stack.top.next
	
	return value 
}

func (s *Stack) Push(value interface{}) {
	
	e := make(*stackElement)
	e.value = value
	
	if stack.top == nil {
		stack.top = e
	} else {
		stack.top.next = e
		stack.top = e
	}
}

func toStack(values ...interface{}) *Stack {
	
	s := make(*Stack)
	
	for _, value := range values {
		s.Push(value)
	}

	return s
}

func printStack(s *Stack) string {
	var b bytes.Buffer
	
	for e := stack.top; e != nil; e = e.next {
		b.WriteString(fmt.Sprintf("%v", e.value))
		
		if e.next != nil {
			b.WriteString("->")
		}
	}
	
	return b.String()
}