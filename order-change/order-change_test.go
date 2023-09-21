package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderChangePoint(t *testing.T) {
	t.Run("return 0 if array is smaller than 3", func(t *testing.T) {
		input := []int{0, 41}

		changePosition := orderChangePoint(input)

		assert.Equal(t, 0, changePosition)
	})
	t.Run("return -1 if no point of change", func(t *testing.T) {
		input := []int{-4, -2, 0, 6, 10}

		changePosition := orderChangePoint(input)

		assert.Equal(t, -1, changePosition)
	})
	t.Run("return point of order change in ascending array", func(t *testing.T) {
		input := []int{1, 2, 4, 6, 4, 3, 1}

		changePosition := orderChangePoint(input)

		assert.Equal(t, 3, changePosition)
	})
	t.Run("return point of order change in decending array", func(t *testing.T) {
		input := []int{5, 4, 3, 10, 11}

		changePosition := orderChangePoint(input)

		assert.Equal(t, 2, changePosition)
	})

	t.Run("when point of order change is the last number", func(t *testing.T) {
		input := []int{2, 4, 6, 10, 5}

		changePosition := orderChangePoint(input)

		assert.Equal(t, 3, changePosition)
	})
}
