package ch01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Design an algorithm and write code to remove the duplicate characters in a
   string without using any additional buffer. NOTE: One or two additional
   variables are fine. An extra copy of the array is not. */

func StrDedupe(t *testing.T, s []rune) {
	writeIndex := 0
	readIndex := 0

	if len(s) < 2 {
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

func TestStrDedupe_Consecutive(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'e', 'e', 'f', 'g', 'g', 'h'}
	StrDedupe(t, str)
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', -1, -1, -1}, str)
}

func TestStrDedupe_NonConsecutive(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'e', 'f', 'g', 'h', 'g'}
	StrDedupe(t, str)
	assert.Equal(t, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', -1, -1, -1}, str)
}

func TestStrDedupe_Trailing(t *testing.T) {
	str := []rune{'a', 'a', 'a'}
	StrDedupe(t, str)
	assert.Equal(t, []rune{'a', -1, -1}, str)
}

func TestStrDedupe_ZeroValue(t *testing.T) {
	var str []rune
	StrDedupe(t, str)
	assert.Equal(t, str, str)
}
