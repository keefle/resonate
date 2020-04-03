package fuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	path := "to/joe/man"
	assert.Equal(t, splitPath(path), []string{"to", "joe", "man"})
}
