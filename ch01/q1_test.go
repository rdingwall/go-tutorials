package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hasAllUniqueChars(s string) bool {	
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
	assert.True(t, hasAllUniqueChars("abcdefghijklmnopqrstuvwyxz"))
}

func TestHasAllUniqueChars_NonUnique(t *testing.T) {
	assert.False(t, hasAllUniqueChars("abcdefggg"))
}

func TestHasAllUniqueChars_ZeroValue(t *testing.T) {
	assert.False(t, hasAllUniqueChars(""))
}
