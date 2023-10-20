package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsequence(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := "abc"
		s2 := "ahbgdc"

		output := isSubsequence(s, s2)

		assert.True(t, output)

	})
	t.Run("case 2", func(t *testing.T) {
		s := "axc"
		s2 := "ahbgdc"

		output := isSubsequence(s, s2)

		assert.False(t, output)
	})

	t.Run("case 3", func(t *testing.T) {
		s := "ace"
		s2 := "abcde"

		output := isSubsequence(s, s2)

		assert.True(t, output)
	})
	t.Run("case 4", func(t *testing.T) {
		s := "aec"
		s2 := "abcde"

		output := isSubsequence(s, s2)

		assert.False(t, output)
	})

	t.Run("case 5", func(t *testing.T) {
		s := "bb"
		s2 := "ahbgdc"

		output := isSubsequence(s, s2)

		assert.False(t, output)
	})
	t.Run("case 5", func(t *testing.T) {
		s := "ab"
		s2 := "baab"

		output := isSubsequence(s, s2)

		assert.True(t, output)
	})
}
