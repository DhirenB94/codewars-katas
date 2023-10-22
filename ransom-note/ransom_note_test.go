package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNote(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := canConstruct("a", "b")

		assert.False(t, output)
	})
	t.Run("case 2", func(t *testing.T) {
		output := canConstruct("aa", "ab")

		assert.False(t, output)
	})
	t.Run("case 1", func(t *testing.T) {
		output := canConstruct("aa", "aab")

		assert.True(t, output)
	})
}