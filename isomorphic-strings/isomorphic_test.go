package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNote(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := isIsomorphic("egg", "add")

		assert.True(t, output)
	})
	t.Run("case 2", func(t *testing.T) {
		output := isIsomorphic("foo", "bar")

		assert.False(t, output)
	})
	t.Run("case 3", func(t *testing.T) {
		output := isIsomorphic("paper", "title")

		assert.True(t, output)
	})
	t.Run("case 4", func(t *testing.T) {
		output := isIsomorphic("bbbaaaba", "aaabbbba")

		assert.False(t, output)
	})
	t.Run("case 5", func(t *testing.T) {
		output := isIsomorphic("badc", "babc")

		assert.False(t, output)
	})
}
