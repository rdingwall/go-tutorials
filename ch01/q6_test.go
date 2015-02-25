package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rotate(t *testing.T, image [][]int) [][]int {
	n := len(image)

	output := make([][]int, n)
	for y := range image {
		output[y] = make([]int, n)
	}

	for y, row := range image {
		for x := range row {
			t.Logf("(%v, %v) = (%v, %v)", y, x, n-x-1, y)
			t.Logf("%v = %v", output[y][x], image[n-x-1][y])
			output[y][x] = image[n-x-1][y]
		}
	}

	return output
}

func TestRotate(t *testing.T) {

	image := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	assert.Equal(t, expected, rotate(t, image))
}
