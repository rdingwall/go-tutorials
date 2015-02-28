package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Describe how you could use a single array to implement three stacks.

var (
	ErrInvalidCapacity = errors.New("invalid capacity")
	ErrOutOfRange      = errors.New("invalid stack")
	ErrTooManyItems    = errors.New("too many items")
)

type MultiStack struct {
	storage []interface{}
	stacks  [][]interface{}
}

func NewMultiStack(count, capacity uint) (ms *MultiStack, err error) {
	if count == 0 || capacity == 0 {
		return nil, ErrInvalidCapacity
	}

	ms = new(MultiStack)
	ms.storage = make([]interface{}, count*capacity)
	ms.stacks = make([][]interface{}, count)

	for i := uint(0); i < count; i++ {
		startIndex := i * capacity
		capacity := startIndex + capacity
		ms.stacks[i] = ms.storage[startIndex:startIndex:capacity]
	}

	return
}

func (ms *MultiStack) Push(stack uint, value interface{}) (err error) {
	if stack+1 > uint(len(ms.stacks)) {
		return ErrOutOfRange
	}

	s := ms.stacks[stack]

	if len(s) == cap(s) {
		return ErrTooManyItems
	}

	s = s[:len(s)+1]
	s[len(s)-1] = value
	ms.stacks[stack] = s

	return
}

func (ms *MultiStack) Pop(stack uint) (value interface{}, err error) {
	if stack+1 > uint(len(ms.stacks)) {
		return nil, ErrOutOfRange
	}

	s := ms.stacks[stack]

	if len(s) == 0 {
		return
	}

	value = s[len(s)-1]
	s = s[:len(s)-1]
	ms.stacks[stack] = s

	return
}

func TestPushPop(t *testing.T) {
	ms, err := NewMultiStack(3, 5)
	assert.NoError(t, err)

	assert.NoError(t, ms.Push(0, 1))
	assert.NoError(t, ms.Push(0, 2))
	assert.NoError(t, ms.Push(0, 3))
	assert.NoError(t, ms.Push(0, 4))
	assert.NoError(t, ms.Push(0, 5))
	assert.EqualError(t, ms.Push(0, 6), ErrTooManyItems.Error())

	assert.NoError(t, ms.Push(1, 'a'))
	assert.NoError(t, ms.Push(1, 'b'))
	assert.NoError(t, ms.Push(1, 'c'))
	assert.NoError(t, ms.Push(1, 'd'))
	assert.NoError(t, ms.Push(1, 'e'))
	assert.EqualError(t, ms.Push(1, 'f'), ErrTooManyItems.Error())

	assert.NoError(t, ms.Push(2, 'A'))
	assert.NoError(t, ms.Push(2, 'B'))
	assert.NoError(t, ms.Push(2, 'C'))
	assert.NoError(t, ms.Push(2, 'D'))
	assert.NoError(t, ms.Push(2, 'E'))
	assert.EqualError(t, ms.Push(2, 'F'), ErrTooManyItems.Error())

	getValue := func(value interface{}, err error) interface{} {
		assert.NoError(t, err)
		return value
	}

	assert.Equal(t, 5, getValue(ms.Pop(0)))
	assert.Equal(t, 4, getValue(ms.Pop(0)))
	assert.Equal(t, 3, getValue(ms.Pop(0)))
	assert.Equal(t, 2, getValue(ms.Pop(0)))
	assert.Equal(t, 1, getValue(ms.Pop(0)))
	assert.Equal(t, nil, getValue(ms.Pop(0)))

	assert.Equal(t, 'e', getValue(ms.Pop(1)))
	assert.Equal(t, 'd', getValue(ms.Pop(1)))
	assert.Equal(t, 'c', getValue(ms.Pop(1)))
	assert.Equal(t, 'b', getValue(ms.Pop(1)))
	assert.Equal(t, 'a', getValue(ms.Pop(1)))
	assert.Equal(t, nil, getValue(ms.Pop(1)))

	assert.Equal(t, 'E', getValue(ms.Pop(2)))
	assert.Equal(t, 'D', getValue(ms.Pop(2)))
	assert.Equal(t, 'C', getValue(ms.Pop(2)))
	assert.Equal(t, 'B', getValue(ms.Pop(2)))
	assert.Equal(t, 'A', getValue(ms.Pop(2)))
	assert.Equal(t, nil, getValue(ms.Pop(2)))
}

func TestErrInvalidCapacity(t *testing.T) {
	ms, err := NewMultiStack(0, 0)
	assert.Nil(t, ms)
	assert.EqualError(t, err, ErrInvalidCapacity.Error())
}

func TestErrOutOfRange(t *testing.T) {
	ms, err := NewMultiStack(3, 5)
	assert.NoError(t, err)

	err = ms.Push(99, 'A')
	assert.EqualError(t, err, ErrOutOfRange.Error())

	_, err = ms.Pop(99)
	assert.EqualError(t, err, ErrOutOfRange.Error())
}
