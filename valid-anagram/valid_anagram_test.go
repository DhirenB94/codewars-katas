package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := isAnagram("anagram", "nagaram")

		assert.True(t, output)
	})
	t.Run("case 2", func(t *testing.T) {
		output := isAnagram("rat", "car")

		assert.False(t, output)
	})
}
