package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := "A man, a plan, a canal: Panama"
		output := isPalindrome(input)

		assert.True(t, output)

	})
	t.Run("case 2", func(t *testing.T) {
		input := "race a car"
		output := isPalindrome(input)

		assert.False(t, output)
	})

	t.Run("case 3", func(t *testing.T) {
		input := " "
		output := isPalindrome(input)

		assert.True(t, output)
	})
}
