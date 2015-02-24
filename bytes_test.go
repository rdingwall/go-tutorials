package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringBuffer(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("aaa ")
	b.WriteString("bbb")
	assert.Equal(t, "aaa bbb", b.String())
}