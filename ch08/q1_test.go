package ch08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write a method to generate the nth Fibonacci number.

func FibRecursive(n uint) uint {
	switch n {
	case 0:
		return 1
	case 1:
		return 2
	default:
		return FibRecursive(n-1) + FibRecursive(n-2)
	}
}

func TestFibRecursive(t *testing.T) {
	assert.Equal(t, 1, FibRecursive(0))
	assert.Equal(t, 2, FibRecursive(1))
	assert.Equal(t, 3, FibRecursive(2))
	assert.Equal(t, 5, FibRecursive(3))
	assert.Equal(t, 8, FibRecursive(4))
	assert.Equal(t, 13, FibRecursive(5))
}

func FibIterative(n uint) (sum uint) {
	sum = 1
	var prev uint = 1
	for i := uint(1); i < n+1; i++ {
		sum, prev = sum+prev, sum
	}
	return
}

func TestFibIterative(t *testing.T) {
	assert.Equal(t, 1, FibIterative(0))
	assert.Equal(t, 2, FibIterative(1))
	assert.Equal(t, 3, FibIterative(2))
	assert.Equal(t, 5, FibIterative(3))
	assert.Equal(t, 8, FibIterative(4))
	assert.Equal(t, 13, FibIterative(5))
}
