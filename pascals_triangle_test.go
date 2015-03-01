package main

import (
	"testing"
)

// Calculate row N of Pascal's triangle.

func PascalTriangleRowRecursive(n uint) (row []uint) {
	if n == 0 {
		return []uint{1}
	}

	row = make([]uint, n+1)

	for i := uint(0); i < n+1; i++ {
		switch i {
		case 0, n:
			row[i] = 1

		default:
			prevRow := PascalTriangleRowRecursive(n - 1)
			row[i] = prevRow[i-1] + prevRow[i]
		}
	}

	return
}

func TestPascalTriangleRowRecursive(t *testing.T) {
	assertEqualContents(t, []uint{1}, PascalTriangleRowRecursive(0))
	assertEqualContents(t, []uint{1, 1}, PascalTriangleRowRecursive(1))
	assertEqualContents(t, []uint{1, 2, 1}, PascalTriangleRowRecursive(2))
	assertEqualContents(t, []uint{1, 3, 3, 1}, PascalTriangleRowRecursive(3))
	assertEqualContents(t, []uint{1, 4, 6, 4, 1}, PascalTriangleRowRecursive(4))
	assertEqualContents(t, []uint{1, 5, 10, 10, 5, 1}, PascalTriangleRowRecursive(5))
}

func PascalTriangleRowIterative(n uint) (row []uint) {
	row = make([]uint, n+1)
	for i := 0; i < int(n)+1; i++ {
		row[i] = 1
		for j := i - 1; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}
	return
}

func TestPascalTriangleRowIterative(t *testing.T) {
	assertEqualContents(t, []uint{1}, PascalTriangleRowIterative(0))
	assertEqualContents(t, []uint{1, 1}, PascalTriangleRowIterative(1))
	assertEqualContents(t, []uint{1, 2, 1}, PascalTriangleRowIterative(2))
	assertEqualContents(t, []uint{1, 3, 3, 1}, PascalTriangleRowIterative(3))
	assertEqualContents(t, []uint{1, 4, 6, 4, 1}, PascalTriangleRowIterative(4))
	assertEqualContents(t, []uint{1, 5, 10, 10, 5, 1}, PascalTriangleRowIterative(5))
}

func assertEqualContents(t *testing.T, expected, actual []uint) {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, actual)
			return
		}
	}
}
