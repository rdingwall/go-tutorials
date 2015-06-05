package ch01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Implement an algorithm to determine if a string has all unique characters.

func HasAllUniqueChars(s string) bool {
	occurrences := make(map[rune]bool)
	for _, c := range s {
		if occurrences[c] {
			return false
		}
		occurrences[c] = true
	}
	return len(s) > 0
}

func TestHasAllUniqueChars_Unique(t *testing.T) {
	assert.True(t, HasAllUniqueChars("abcdefghijklmnopqrstuvwyxz"))
}

func TestHasAllUniqueChars_NonUnique(t *testing.T) {
	assert.False(t, HasAllUniqueChars("abcdefggg"))
}

func TestHasAllUniqueChars_ZeroValue(t *testing.T) {
	assert.False(t, HasAllUniqueChars(""))
}
