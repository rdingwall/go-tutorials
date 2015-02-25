package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeZeros(t *testing.T, matrix [][]int) {

	rowZero := func(matrix [][]int, y int) {
		row := matrix[y]
		for x := range row {
			row[x] = 0
		}
	}

	colZero := func(matrix [][]int, x int) {
		for y := range matrix {
			matrix[y][x] = 0
		}
	}

	for y, row := range matrix {
		for x, value := range row {
			if value == 0 {
				defer rowZero(matrix, y)
				defer colZero(matrix, x)
			}
		}
	}
}

func TestRowColSetToZero(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 9},
	}

	expected := [][]int{
		{1, 2, 0},
		{0, 0, 0},
		{7, 8, 0},
	}

	writeZeros(t, matrix)

	assert.Equal(t, expected, matrix)
}
