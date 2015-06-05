package ch01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Given an image represented by an NxN matrix, where each pixel in the image
   is 4 bytes, write a method to rotate the image by 90 degrees. */

func Rotate(t *testing.T, image [][]int) (output [][]int) {
	n := len(image)

	output = make([][]int, n)
	for y := range image {
		output[y] = make([]int, n)
	}

	for y, row := range image {
		for x := range row {
			output[y][x] = image[n-x-1][y]
		}
	}

	return
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

	assert.Equal(t, expected, Rotate(t, image))
}
