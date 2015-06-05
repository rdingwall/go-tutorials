package ch08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Imagine a robot sitting on the upper left hand corner of an NxN grid. The
   robot can only move in two directions: right and down. How many possible
   paths are there for the robot? */

func CalculatePaths(t *testing.T, n, x, y uint) uint {
	switch {
	case x+1 >= n && y+1 >= n:
		return 0
	case x+1 == n:
		fallthrough
	case y+1 == n:
		return 1
	default:
		return CalculatePaths(t, n, x+1, y) + CalculatePaths(t, n, x, y+1)
	}
}

func TestCalculatePathsFromStart(t *testing.T) {
	assert.Equal(t, 70, CalculatePaths(t, 5, 0, 0))
}

func TestCalculatePathsFromEnd(t *testing.T) {
	assert.Equal(t, 0, CalculatePaths(t, 5, 4, 4))
}

func TestCalculatePathsOutOfRange(t *testing.T) {
	assert.Equal(t, 0, CalculatePaths(t, 5, 9, 8))
}
