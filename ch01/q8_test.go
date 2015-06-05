package ch01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Assume you have a method isSubstring which checks if one word is a substring
   of another. Given two strings, s1 and s2, write code to check if s2 is a
   rotation of s1 using only one call to isSubstring (i.e., “waterbottle” is a
   rotation of “erbottlewat”). */

func IsRotation(t *testing.T, s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}

func TestIsRotation(t *testing.T) {
	assert.True(t, IsRotation(t, "waterbottle", "erbottlewat"))
}

func TestIsRotationFalse(t *testing.T) {
	assert.False(t, IsRotation(t, "waterboxx", "erbottlewat"))
}

func TestIsRotationEmptyS2(t *testing.T) {
	assert.True(t, IsRotation(t, "xx", ""))
}

func TestIsRotationEmptyS1(t *testing.T) {
	assert.False(t, IsRotation(t, "", "xx"))
}

func TestIsRotationEmpty(t *testing.T) {
	assert.True(t, IsRotation(t, "", ""))
}
