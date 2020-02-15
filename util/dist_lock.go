package util

import (
	"log"
	"sync"
)

// Dinning philosophers

type LockManager struct {
	locks map[string]bool
	mu    sync.Mutex
}

func NewLockManager() *LockManager {
	return &LockManager{
		locks: make(map[string]bool),
	}
}

func (l *LockManager) Lock(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.locks[key] {
		return false
	}

	l.locks[key] = true

	return true
}

func (l *LockManager) Unlock(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.locks[key] {
		log.Fatal("Why are you trying to unlock something that you didn't even lock")
		return false
	}

	l.locks[key] = false

	return true
}
