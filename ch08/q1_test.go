package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fibRecursive(n uint) uint {
	switch n {
	case 0:
		return 1
	case 1:
		return 2
	default:
		return fibRecursive(n-1) + fibRecursive(n-2)
	}
}

func TestFibRecursive(t *testing.T) {
	assert.Equal(t, 1, fibRecursive(0))
	assert.Equal(t, 2, fibRecursive(1))
	assert.Equal(t, 3, fibRecursive(2))
	assert.Equal(t, 5, fibRecursive(3))
	assert.Equal(t, 8, fibRecursive(4))
	assert.Equal(t, 13, fibRecursive(5))
}

func fibIterative(n uint) uint {
	var sum, prev uint = 1, 1
	for i := uint(1); i < n+1; i++ {
		sum, prev = sum+prev, sum
	}
	return sum
}

func TestFibIterative(t *testing.T) {
	assert.Equal(t, 1, fibIterative(0))
	assert.Equal(t, 2, fibIterative(1))
	assert.Equal(t, 3, fibIterative(2))
	assert.Equal(t, 5, fibIterative(3))
	assert.Equal(t, 8, fibIterative(4))
	assert.Equal(t, 13, fibIterative(5))
}
