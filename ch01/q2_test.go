package ch01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write code to reverse a C-Style String.

func StrRev(s []rune) {

	if s == nil {
		return
	}

	var length int
	for i, c := range s {
		if c == -1 {
			length = i
			break
		}
	}

	for i := 0; length-i-1 > i; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}

	s[length] = -1
}

func TestStrRev(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', -1}
	StrRev(str)
	assert.Equal(t, []rune{'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a', -1}, str)
}

func TestStrRev_ZeroValue(t *testing.T) {
	var str []rune
	StrRev(str)
	assert.Equal(t, str, str)
}

func TestStrRev_Empty(t *testing.T) {
	str := []rune{-1}
	StrRev(str)
	assert.Equal(t, []rune{-1}, str)
}
