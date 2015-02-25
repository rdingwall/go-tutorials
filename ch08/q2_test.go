package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculatePaths(n, x, y uint) uint {
	var pathCount uint

	if x < n-1 {
		pathCount += calculatePaths(n, x+1, y)
	}

	if y < n-1 {
		pathCount += calculatePaths(n, x, y+1)
	}

	return pathCount
}

func CalculatePathsTest(t *testing.T) {
	assert.Equal(t, 0, calculatePaths(5, 4, 4))
	assert.Equal(t, 70, calculatePaths(5, 0, 0))
}
