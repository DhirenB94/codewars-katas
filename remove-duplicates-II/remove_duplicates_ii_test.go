package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []int{1, 1, 1, 2, 2, 3}
		output := removeDuplicate(input)

		assert.Equal(t, 5, output)
	})
	t.Run("case 2", func(t *testing.T) {
		input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		output := removeDuplicate(input)

		assert.Equal(t, 7, output)
	})
}
