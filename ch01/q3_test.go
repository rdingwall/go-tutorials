package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strDedupe(t * testing.T, s []rune) {
	writeIndex := 0
	readIndex := 0
	
	if (len(s) < 2) {
		return
	}

	// crap but works
	for {
		
		if readIndex == len(s) || writeIndex == len(s) {
			break
		}
		
		var i int
		var found bool
		for i = 0; i <= writeIndex; i++ {
			
			if i == readIndex {
				continue
			}
			
			if s[i] == s[readIndex] {
				found = true
				break
			}
		}
		
		if found {
			readIndex++
			continue
		}
		
		s[writeIndex] = s[readIndex]
		writeIndex++
		readIndex++
	}
	
	// Zero out trailing runes
	for i := writeIndex; i < len(s); i++ {
		s[i] = -1
	}
}

func TestRemoveDuplicateChars_Consecutive(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'e', 'e', 'f', 'g', 'g', 'h'}
	strDedupe(t, str)
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', -1, -1, -1}, str)
}

func TestRemoveDuplicateChars_NonConsecutive(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'e', 'f', 'g', 'h', 'g'}
	strDedupe(t, str)
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', -1, -1, -1}, str)
}

func TestRemoveDuplicateChars_Trailing(t *testing.T) {
	str := []rune{'a', 'a', 'a'}
	strDedupe(t, str)
	assert.Equal(t, []rune{'a', -1, -1}, str)
}

func TestRemoveDuplicateChars_ZeroValue(t *testing.T) {
	var str []rune
	strDedupe(t, str)
	assert.Equal(t, str, str)
}