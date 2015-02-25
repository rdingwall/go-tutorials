package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDelete(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 42
	assert.Equal(t, 42, m["a"])
	delete(m, "a")
	assert.Equal(t, 0, m["a"])
}

func TestTryGet(t *testing.T) {
	m := make(map[string]int)

	m["b"] = 42

	i, ok := m["b"]

	assert.True(t, ok)
	assert.Equal(t, 42, i)
}

func TestEnumerateKeys(t *testing.T) {
	m := make(map[string]int)

	m["a"] = 42
	m["b"] = 62

	for key := range m {
		t.Logf("key: %v", key)
	}
}

func TestEnumerate(t *testing.T) {
	m := make(map[string]int)

	m["a"] = 42
	m["b"] = 62

	for key, value := range m {
		t.Logf("key: %v, value: %v", key, value)
	}
}

func TestLiteral(t *testing.T) {
	m := map[string]int{
		"a": 42,
		"b": 62, // trailing comma is required
	}

	m["c"] = 99
}

func TestEmptyLiteral(t *testing.T) {
	m := map[string]int{}

	m["c"] = 99
}

func TestStructKey(t *testing.T) {
	type Key struct {
		FirstName string
		LastName  string
	}

	m := make(map[Key]int)

	m[Key{"Aa", "Bb"}] = 42

	assert.Equal(t, 42, m[Key{"Aa", "Bb"}])
}

func TestConcurrency(t *testing.T) {
	var counter = struct {
		sync.RWMutex // embedded lock
		m            map[string]int
	}{m: make(map[string]int)}

	var wg sync.WaitGroup

	increment := func() {
		defer counter.Unlock() // LIFO defer order
		defer wg.Done()
		counter.Lock()
		counter.m["count"]++
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	assert.Equal(t, 3, counter.m["count"])
}
