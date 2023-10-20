package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		expectedOutput := 7

		output := maxProfit2(input)

		assert.Equal(t, expectedOutput, output)
	})
	t.Run("case 2", func(t *testing.T) {
		input := []int{2, 4, 1}
		expectedOutput := 2

		output := maxProfit2(input)

		assert.Equal(t, expectedOutput, output)
	})

	t.Run("case 3", func(t *testing.T) {
		input := []int{7, 6, 4, 3, 1}
		expectedOutput := 0

		output := maxProfit2(input)

		assert.Equal(t, expectedOutput, output)
	})
	t.Run("case 3", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expectedOutput := 4

		output := maxProfit2(input)

		assert.Equal(t, expectedOutput, output)
	})
}
