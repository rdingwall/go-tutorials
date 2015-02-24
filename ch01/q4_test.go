package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func areAnagrams(t *testing.T, s1 string, s2 string) bool {	
	if len(s1) != len(s2) {
		return false
	}
	
	for _, c1 := range s1 {
		var found bool
		t.Logf("Looking for %c..", c1)
		
		for _, c2 := range s2 {
			if c1 == c2 {
				t.Logf("Found %c!", c1)
				found = true
				break
			}
		}
		
		if !found {
			t.Logf("Did not find %c!", c1)
			return false
		}
	}
	
	for _, c2 := range s2 {
		var found bool

		t.Logf("Looking for %c..", c2)
		
		for _, c1 := range s1 {
			if c2 == c1 {
				t.Logf("Found %c!", c2)
				found = true
				break
			}
		}
		
		if !found {
			t.Logf("Did not find %c!", c2)
			return false
		}
	}
	
	return true
}

func TestAreAnagrams_True(t *testing.T) {
	assert.True(t, areAnagrams(t, "abcde", "becad"))
}

func TestAreAnagrams_False(t *testing.T) {
	assert.False(t, areAnagrams(t, "bbcde", "becxd"))
}

func TestAreAnagrams_False2(t *testing.T) {
	assert.False(t, areAnagrams(t, "bbcdd", "becxd"))
}

func TestAreAnagrams_Same(t *testing.T) {
	assert.True(t, areAnagrams(t, "abcde", "abcde"))
}

func TestAreAnagrams_Empty(t *testing.T) {
	assert.True(t, areAnagrams(t, "", ""))
}

func TestAreAnagrams_EmptyNonEmpty(t *testing.T) {
	assert.False(t, areAnagrams(t, "", "xxx"))
}

func TestAreAnagrams_DifferentLengths(t *testing.T) {
	assert.False(t, areAnagrams(t, "abcdexxx", "becxd"))
}