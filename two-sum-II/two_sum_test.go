package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("return index of the 2 numbers that sum to make the target number", func(t *testing.T) {
		t.Run("case 1", func(t *testing.T) {
			input := []int{2, 7, 11, 15}
			expectedOutput := []int{1, 2}

			output := twoSum(input, 9)

			assert.Equal(t, expectedOutput, output)

		})
		t.Run("case 2", func(t *testing.T) {
			input := []int{2, 3, 4}
			expectedOutput := []int{1, 3}

			output := twoSum(input, 6)

			assert.Equal(t, expectedOutput, output)

		})
		t.Run("case 3", func(t *testing.T) {
			input := []int{-1, 0}
			expectedOutput := []int{1, 2}

			output := twoSum(input, -1)

			assert.Equal(t, expectedOutput, output)

		})
		t.Run("case 2", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 4, 9, 56, 90}
			expectedOutput := []int{4, 5}

			output := twoSum(input, 8)

			assert.Equal(t, expectedOutput, output)

		})
	})

}
