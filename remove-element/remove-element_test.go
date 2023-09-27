package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		numberOfNumsRemaining := removeElement([]int{3, 2, 2, 3}, 3)

		assert.Equal(t, 2, numberOfNumsRemaining)
	})
	t.Run("case 1", func(t *testing.T) {
		numberOfNumsRemaining := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

		assert.Equal(t, 5, numberOfNumsRemaining)
	})

}
