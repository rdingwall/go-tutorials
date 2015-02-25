package main

import (
	"testing"
)

func pascalRow(n int) []int {
	var row, prevRow []int

	for i := 0; i < n+1; i++ {
		row = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			switch j {
			case 0, i:
				row[j] = 1

			default:
				row[j] = prevRow[j-1] + prevRow[j]
			}
		}
		
		prevRow = row
	}

	return row
}

func pascalRowR(n int) []int {
	if n == 0 {
		return []int{1}
	}

	row := make([]int, n+1)

	for i := 0; i < n+1; i++ {
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

func assertEquals(t *testing.T, expected, actual []int) {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, actual)
			return
		}
	}
}

func TestPascalRow(t *testing.T) {
	assertEquals(t, []int{1}, pascalRow(0))
	assertEquals(t, []int{1, 1}, pascalRow(1))
	assertEquals(t, []int{1, 2, 1}, pascalRow(2))
	assertEquals(t, []int{1, 3, 3, 1}, pascalRow(3))
	assertEquals(t, []int{1, 4, 6, 4, 1}, pascalRow(4))
	assertEquals(t, []int{1, 5, 10, 10, 5, 1}, pascalRow(5))
}

func TestPascalRowRecursive(t *testing.T) {
	assertEquals(t, []int{1}, pascalRowR(0))
	assertEquals(t, []int{1, 1}, pascalRowR(1))
	assertEquals(t, []int{1, 2, 1}, pascalRowR(2))
	assertEquals(t, []int{1, 3, 3, 1}, pascalRowR(3))
	assertEquals(t, []int{1, 4, 6, 4, 1}, pascalRowR(4))
	assertEquals(t, []int{1, 5, 10, 10, 5, 1}, pascalRowR(5))
}
