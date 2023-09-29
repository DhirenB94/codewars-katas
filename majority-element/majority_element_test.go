package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []int{3, 2, 3}
		output := majorityElement(input)

		assert.Equal(t, 3, output)
	})
	t.Run("case 2", func(t *testing.T) {
		input := []int{0, 0, 1, 1, 1, 1, 1, 1, 2, 3, 3}
		output := majorityElement(input)

		assert.Equal(t, 1, output)
	})
	t.Run("case 3", func(t *testing.T) {
		input := []int{2, 2, 1, 1, 1, 2, 2}
		output := majorityElement(input)

		assert.Equal(t, 2, output)
	})
}
