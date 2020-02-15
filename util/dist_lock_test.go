package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	lm := NewLockManager()

	assert.True(t, lm.Lock("hello"), "newly locked file must return true")
	assert.False(t, lm.Lock("hello"), "already locked file must return false")
	assert.True(t, lm.Unlock("hello"), "already locked file should be capable of unlocking")
}
