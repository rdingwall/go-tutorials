package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	count := 0
	var once sync.Once

	increment := func() {
		count++
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}

	wg.Wait()

	assert.Equal(t, 1, count)
}

func TestRWMutex(t *testing.T) {
	var mutex sync.RWMutex
	var i int32
	var wg sync.WaitGroup

	write := func() {
		defer mutex.Unlock()
		defer wg.Done()
		mutex.Lock()
		i++
	}

	read := func() {
		defer mutex.RUnlock()
		defer wg.Done()
		mutex.RLock()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()

		wg.Add(1)
		go read()
	}

	wg.Wait() // no deadlocks
}

func TestPool(t *testing.T) {
	var i int32
	new := func() interface{} {
		return atomic.AddInt32(&i, 1)
	}

	pool := sync.Pool{New: new}

	for i := 0; i < 100; i++ {
		pool.Get()
	}

	assert.NotEqual(t, 0, i)
}
