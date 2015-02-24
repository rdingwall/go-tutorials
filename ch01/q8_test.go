package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isRotation(t *testing.T, s1, s2 string) bool {
	return strings.Contains(s1 + s1, s2)
}

func TestIsRotation(t *testing.T) {
	assert.True(t, isRotation(t, "waterbottle", "erbottlewat"))
}

func TestIsRotationFalse(t *testing.T) {
	assert.False(t, isRotation(t, "waterboxx", "erbottlewat"))
}

func TestIsRotationEmptyS2(t *testing.T) {
	assert.True(t, isRotation(t, "xx", ""))
}

func TestIsRotationEmptyS1(t *testing.T) {
	assert.False(t, isRotation(t, "", "xx"))
}

func TestIsRotationEmpty(t *testing.T) {
	assert.True(t, isRotation(t, "", ""))
}