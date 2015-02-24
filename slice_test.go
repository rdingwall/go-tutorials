package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceLength(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	len := len(slice)
	assert.Equal(t, 5, len)
}

func TestSliceCapacity(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	cap := cap(slice)
	assert.Equal(t, 5, cap)
}

func TestSlicing(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[1:3]
	assert.Equal(t, []rune{'b', 'c'}, slice)
}

func TestSlicingToEnd(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[:3]
	assert.Equal(t, []rune{'a', 'b', 'c'}, slice)
}

func TestSlicingFromStart(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[3:]
	assert.Equal(t, []rune{'d', 'e'}, slice)
}

func TestSlicingStartToEnd(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[:] // sharing same storage as runeArray
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e'}, slice)
}

func TestSliceAndArrayTypesDifferent(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[:]
	assert.NotEqual(t, array, slice) // [5]rune (array) != []rune (slice)
}

func TestModifyUnderlyingElements(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[1:3]
	slice[0], slice[1] = 'B', 'C'
	assert.Equal(t, [...]rune{'a', 'B', 'C', 'd', 'e'}, array)
}

func TestResliceToCapacity(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice := array[2:3]
	slice = slice[:cap(slice)] // can only grow up to capacity, not below zero
	assert.Equal(t, []rune{'c', 'd', 'e'}, slice)
}

func TestIncreaseCapacity(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := make([]rune, len(slice), cap(slice)*2)
	copy(slice2, slice)
	assert.Equal(t, slice, slice2)
}

func TestAppend(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := append(slice, 'f', 'g', 'h')
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}, slice2)
}

func TestAppendToNilSlice(t *testing.T) {
	slice2 := append([]rune(nil), 'a', 'b', 'c')
	assert.Equal(t, []rune{'a', 'b', 'c'}, slice2)
}

func TestConcat(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := []rune{'f', 'g', 'h'}
	slice3 := append(slice, slice2...)
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}, slice3)
}

func TestCopyToAllowUnderlyingArrayToBeGarbageCollected(t *testing.T) {
	array := [...]rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := array[1:3]
	slice3 := append([]rune(nil), slice2...)
	for i := range array {
		array[i] = 0
	}
	assert.Equal(t, []rune{'b', 'c'}, slice3)
}

func Test2DSlice(t *testing.T) {
	matrix := make([][]rune, 3)
	matrix[0] = []rune{'x', 'o', 'x'}
	matrix[1] = []rune{'o', 'o', 'x'}
	matrix[2] = []rune{'x', 'x', 'o'}
}

func TestCut(t *testing.T) {
	cut := func(s []rune, index int) []rune {
		return append(s[:index], s[index+1:]...)
	}

	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := cut(slice, 2)
	assert.Equal(t, []rune{'a', 'b', 'd', 'e'}, slice2)
}

func TestInsert(t *testing.T) {
	insert := func(s []rune, r rune, index int) []rune {
		s = append(s, 0)
		copy(s[index+1:], s[index:])
		s[index] = r
		return s
	}

	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	slice2 := insert(slice[:], 'X', 2)
	assert.Equal(t, []rune{'a', 'b', 'X', 'c', 'd', 'e'}, slice2)
}
