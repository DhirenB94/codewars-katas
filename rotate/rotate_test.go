package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	t.Run("case one", func(t *testing.T) {
		got := rotateRight([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		want := []int{5, 6, 7, 1, 2, 3, 4}

		assert.Equal(t, want, got)

	})
	t.Run("case one", func(t *testing.T) {
		got := rotateRight([]int{-1, -100, 3, 99}, 2)
		want := []int{3, 99, -1, -100}

		assert.Equal(t, want, got)

	})

}
