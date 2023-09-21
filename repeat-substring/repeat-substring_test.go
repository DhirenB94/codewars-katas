package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatSubstring(t *testing.T) {
	t.Run("input only contains zero character, return -1", func(t *testing.T) {
		input := ""
		output := repeatSubstring(input)

		assert.Equal(t, "-1", output)
	})
	t.Run("input only contains one character, return -1", func(t *testing.T) {
		input := "a"
		output := repeatSubstring(input)

		assert.Equal(t, "-1", output)
	})
	t.Run("repeat found", func(t *testing.T) {
		input := "aaa"
		output := repeatSubstring(input)

		assert.Equal(t, "a", output)
	})
	t.Run("no repeat found", func(t *testing.T) {
		input := "abc"
		output := repeatSubstring(input)

		assert.Equal(t, "-1", output)
	})
	t.Run("repeat sequence found", func(t *testing.T) {
		input := "abab"
		output := repeatSubstring(input)

		assert.Equal(t, "ab", output)
	})
	t.Run("repeat sequence not found", func(t *testing.T) {
		input := "abac"
		output := repeatSubstring(input)

		assert.Equal(t, "-1", output)
	})
	t.Run("abcabcabc should return abc", func(t *testing.T) {
		input := "abcabcabc"
		output := repeatSubstring(input)

		assert.Equal(t, "abc", output)
	})
	t.Run("abcabcabcabc should return abcabc", func(t *testing.T) {
		input := "abcabcabcabc"
		output := repeatSubstring(input)

		assert.Equal(t, "abcabc", output)
	})

}
