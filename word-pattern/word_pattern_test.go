package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := wordPattern("abba", "dog cat cat dog")

		assert.True(t, output)
	})
	t.Run("case 2", func(t *testing.T) {
		output := wordPattern("abba", "dog cat cat fish")

		assert.False(t, output)
	})
	t.Run("case 3", func(t *testing.T) {
		output := wordPattern("aaaa", "dog cat cat dog")

		assert.False(t, output)
	})
	t.Run("case 4", func(t *testing.T) {
		output := wordPattern("abba", "dog dog dog")

		assert.False(t, output)
	})
	t.Run("case 5, incorrect lengths", func(t *testing.T) {
		output := wordPattern("aaaaa", "dog cat cat dog")

		assert.False(t, output)
	})
}
