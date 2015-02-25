package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrAlreadyInitialized = errors.New("already initialized")
var ErrNotInitialized = errors.New("not initialized")
var ErrOutOfRange = errors.New("invalid stack")
var ErrTooManyItems = errors.New("too many items")

type stackData struct {
	startIndex uint32
	size       uint32
	length     uint32
}

type stack struct {
	elements []interface{}

	stackCount uint8
	stackData  []stackData
}

func (s *stack) initialize(t *testing.T, stackCount uint8, stackSize uint32) error {

	if s.stackCount > 0 {
		return ErrAlreadyInitialized
	}

	s.stackData = make([]stackData, stackCount)
	s.elements = make([]interface{}, uint32(stackCount)*stackSize)
	s.stackCount = stackCount

	for i := uint8(0); i < stackCount; i++ {
		s.stackData[i].startIndex = uint32(i) * stackSize
		s.stackData[i].size = stackSize

		t.Logf("Initialized stack %v (startIndex=%v, size=%v)", i, s.stackData[i].startIndex, s.stackData[i].size)
	}

	return nil
}

func (s *stack) push(t *testing.T, stack uint8, value interface{}) error {

	if s.stackCount == 0 {
		return ErrNotInitialized
	}

	if stack > s.stackCount-1 {
		return ErrOutOfRange
	}

	var stackData *stackData = &s.stackData[stack]

	if stackData.length == stackData.size {
		return ErrTooManyItems
	}

	index := stackData.startIndex + stackData.length
	s.elements[index] = value

	stackData.length++

	t.Logf("Pushed %v onto stack %v at index=%v (new length=%v)", value, stack, index, stackData.length)

	return nil
}

func (s *stack) pop(t *testing.T, stack uint8) (interface{}, error) {

	if s.stackCount == 0 {
		return nil, ErrNotInitialized
	}

	if stack > s.stackCount-1 {
		return nil, ErrOutOfRange
	}

	var stackData *stackData = &s.stackData[stack]

	if stackData.length == 0 {
		return nil, nil
	}

	index := stackData.startIndex + stackData.length - 1
	value := s.elements[index]

	stackData.length--

	t.Logf("Popped %v off stack %v from index=%v (new length=%v)", value, stack, index, stackData.length)

	return value, nil
}

func TestStack(t *testing.T) {
	var stack stack
	assert.NoError(t, stack.initialize(t, 3, 5))

	assert.NoError(t, stack.push(t, 0, 1))
	assert.NoError(t, stack.push(t, 0, 2))
	assert.NoError(t, stack.push(t, 0, 3))
	assert.NoError(t, stack.push(t, 0, 4))
	assert.NoError(t, stack.push(t, 0, 5))

	assert.NoError(t, stack.push(t, 1, 'a'))
	assert.NoError(t, stack.push(t, 1, 'b'))
	assert.NoError(t, stack.push(t, 1, 'c'))
	assert.NoError(t, stack.push(t, 1, 'd'))
	assert.NoError(t, stack.push(t, 1, 'e'))

	assert.NoError(t, stack.push(t, 2, 'A'))
	assert.NoError(t, stack.push(t, 2, 'B'))
	assert.NoError(t, stack.push(t, 2, 'C'))
	assert.NoError(t, stack.push(t, 2, 'D'))
	assert.NoError(t, stack.push(t, 2, 'E'))

	getValue := func(value interface{}, err error) interface{} {
		assert.NoError(t, err)
		return value
	}

	assert.Equal(t, 5, getValue(stack.pop(t, 0)))
	assert.Equal(t, 4, getValue(stack.pop(t, 0)))
	assert.Equal(t, 3, getValue(stack.pop(t, 0)))
	assert.Equal(t, 2, getValue(stack.pop(t, 0)))
	assert.Equal(t, 1, getValue(stack.pop(t, 0)))
	assert.Equal(t, nil, getValue(stack.pop(t, 0)))

	assert.Equal(t, 'e', getValue(stack.pop(t, 1)))
	assert.Equal(t, 'd', getValue(stack.pop(t, 1)))
	assert.Equal(t, 'c', getValue(stack.pop(t, 1)))
	assert.Equal(t, 'b', getValue(stack.pop(t, 1)))
	assert.Equal(t, 'a', getValue(stack.pop(t, 1)))
	assert.Equal(t, nil, getValue(stack.pop(t, 1)))

	assert.Equal(t, 'E', getValue(stack.pop(t, 2)))
	assert.Equal(t, 'D', getValue(stack.pop(t, 2)))
	assert.Equal(t, 'C', getValue(stack.pop(t, 2)))
	assert.Equal(t, 'B', getValue(stack.pop(t, 2)))
	assert.Equal(t, 'A', getValue(stack.pop(t, 2)))
	assert.Equal(t, nil, getValue(stack.pop(t, 2)))
}

func TestErrNotInitialized(t *testing.T) {
	var stack stack

	_, err := stack.pop(t, 0)
	assert.EqualError(t, err, ErrNotInitialized.Error())

	err = stack.push(t, 0, 'A')
	assert.EqualError(t, err, ErrNotInitialized.Error())
}

func TestErrOutOfRange(t *testing.T) {
	var stack stack

	assert.NoError(t, stack.initialize(t, 3, 5))

	err := stack.push(t, 99, 'A')
	assert.EqualError(t, err, ErrOutOfRange.Error())

	_, err = stack.pop(t, 99)
	assert.EqualError(t, err, ErrOutOfRange.Error())
}
