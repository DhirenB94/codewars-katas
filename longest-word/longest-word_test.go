package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestWord(t *testing.T) {
	t.Run("longest word", func(t *testing.T) {
		t.Run("empty string passed", func(t *testing.T) {
			input := ""
			output, _ := longestWord(input, "")

			assert.Equal(t, "", output)
		})
		t.Run("one word passed", func(t *testing.T) {
			input := "hello"
			output, _ := longestWord(input, "")

			assert.Equal(t, "hello", output)
		})
		t.Run("more than one word passed, longest returned", func(t *testing.T) {
			input := "Good morning"
			output, _ := longestWord(input, "")

			assert.Equal(t, "morning", output)
		})
		t.Run("more than one word with the same length passed, only 1st longest returned", func(t *testing.T) {
			input := "i love dogs "
			output, _ := longestWord(input, "")

			assert.Equal(t, "love", output)
		})
		t.Run("ignore punctuation to return the longest word", func(t *testing.T) {
			input := "hello there, all good???"
			output, _ := longestWord(input, "")

			assert.Equal(t, "hello", output)
		})
	})
	t.Run("final output", func(t *testing.T) {
		t.Run("return the longest word and the reversed longest word", func(t *testing.T) {
			input := "hello mate"

			expectedOutput := "hello"
			expectedFinalOutput := "olleh:"

			output, finalOutput := longestWord(input, "")

			assert.Equal(t, output, expectedOutput)
			assert.Equal(t, finalOutput, expectedFinalOutput)
		})
		t.Run("return correct output when challenge token provided", func(t *testing.T) {
			input := "hello mate"
			token := "uopxzfve15c"

			expectedOutput := "hello"
			expectedFinalOutput := "olleh:c51evfzxpou"

			output, finalOutput := longestWord(input, token)

			assert.Equal(t, output, expectedOutput)
			assert.Equal(t, finalOutput, expectedFinalOutput)
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
