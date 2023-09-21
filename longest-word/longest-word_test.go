package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestWord(t *testing.T) {
	t.Run("longest word", func(t *testing.T) {
		t.Run("empty string passed returns an error", func(t *testing.T) {
			input := ""
			output, err := longestWord(input, "")

			assert.Nil(t, output)
			assert.Error(t, err)
		})
		t.Run("one word passed", func(t *testing.T) {
			input := "hello"
			output, err := longestWord(input, "")

			assert.Equal(t, "hello", output.longestWord)
			assert.NoError(t, err)
		})
		t.Run("more than one word passed, longest returned", func(t *testing.T) {
			input := "Good morning"
			output, err := longestWord(input, "")

			assert.Equal(t, "morning", output.longestWord)
			assert.NoError(t, err)
		})
		t.Run("more than one word with the same length passed, only 1st longest returned", func(t *testing.T) {
			input := "i love dogs "
			output, err := longestWord(input, "")

			assert.Equal(t, "love", output.longestWord)
			assert.NoError(t, err)
		})
		t.Run("ignore punctuation to return the longest word", func(t *testing.T) {
			input := "hello there, all good???"
			output, err := longestWord(input, "")

			assert.Equal(t, "hello", output.longestWord)
			assert.NoError(t, err)
		})
	})
	t.Run("final output", func(t *testing.T) {
		t.Run("return the longest word and the reversed longest word as the final output when no token provided", func(t *testing.T) {
			input := "hello mate"

			expectedOutput := "hello"
			expectedFinalOutput := "olleh"

			output, err := longestWord(input, "")
			assert.NoError(t, err)

			assert.Equal(t, output.longestWord, expectedOutput)
			assert.Equal(t, output.finalOutput, expectedFinalOutput)
		})
		t.Run("return correct output when challenge token provided", func(t *testing.T) {
			input := "hello mate"
			token := "uopxzfve15c"

			expectedOutput := "hello"
			expectedFinalOutput := "olleh:c51evfzxpou"

			output, err := longestWord(input, token)
			assert.NoError(t, err)

			assert.Equal(t, output.longestWord, expectedOutput)
			assert.Equal(t, output.finalOutput, expectedFinalOutput)
		})
		t.Run("returns an error when token provided but no words", func(t *testing.T) {
			input := ""
			token := "uopxzfve15c"

			output, err := longestWord(input, token)
			assert.Nil(t, output)
			assert.Error(t, err)
			assert.ErrorContains(t, err, "input cannot be empty")
		})
	})
}

func TestRemoveSpeacialChars(t *testing.T) {
	t.Run("word should not be changed if no special chars", func(t *testing.T) {
		input := "hello"
		output := removeSpecialChars(input)

		assert.Equal(t, "hello", output)
	})
	t.Run("remove special chars", func(t *testing.T) {
		input := "hello%$?"
		output := removeSpecialChars(input)

		assert.Equal(t, "hello", output)
	})
	t.Run("remove punctuation", func(t *testing.T) {
		input := "?,hel.lo:"
		output := removeSpecialChars(input)

		assert.Equal(t, "hello", output)
	})
	t.Run("remove numbers", func(t *testing.T) {
		input := "h23ello4"
		output := removeSpecialChars(input)

		assert.Equal(t, "hello", output)
	})
}

func TestReverser(t *testing.T) {
	t.Run("reverse the word", func(t *testing.T) {
		input := "hello"
		output := reverser(input)

		assert.Equal(t, "olleh", output)
	})
}
