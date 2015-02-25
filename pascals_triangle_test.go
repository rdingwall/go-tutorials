package main

import (
	"testing"
)

// Calculate row N of pascal's triangle in place
func pascalRow(n uint) []uint {
	row := make([]uint, n+1)
	for i := 0; i < int(n)+1; i++ {
		row[i] = 1
		for j := i - 1; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}
	return row
}

// Calculate row N of Pascal's triangle recursively
func pascalRowR(n uint) []uint {
	if n == 0 {
		return []uint{1}
	}

	row := make([]uint, n+1)

	for i := uint(0); i < n+1; i++ {
		switch i {
		case 0, n:
			row[i] = 1

		default:
			prevRow := pascalRowR(n - 1)
			row[i] = prevRow[i-1] + prevRow[i]
		}
	}

	return row
}

func assertEquals(t *testing.T, expected, actual []uint) {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, actual)
			return
		}
	}
}

func TestPascalRow(t *testing.T) {
	assertEquals(t, []uint{1}, pascalRow(0))
	assertEquals(t, []uint{1, 1}, pascalRow(1))
	assertEquals(t, []uint{1, 2, 1}, pascalRow(2))
	assertEquals(t, []uint{1, 3, 3, 1}, pascalRow(3))
	assertEquals(t, []uint{1, 4, 6, 4, 1}, pascalRow(4))
	assertEquals(t, []uint{1, 5, 10, 10, 5, 1}, pascalRow(5))
}

func TestPascalRowRecursive(t *testing.T) {
	assertEquals(t, []uint{1}, pascalRowR(0))
	assertEquals(t, []uint{1, 1}, pascalRowR(1))
	assertEquals(t, []uint{1, 2, 1}, pascalRowR(2))
	assertEquals(t, []uint{1, 3, 3, 1}, pascalRowR(3))
	assertEquals(t, []uint{1, 4, 6, 4, 1}, pascalRowR(4))
	assertEquals(t, []uint{1, 5, 10, 10, 5, 1}, pascalRowR(5))
}
