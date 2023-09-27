package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3

		merge(nums1, m, nums2, n)

		assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	})
	t.Run("case 2", func(t *testing.T) {
		nums1 := []int{1}
		m := 1
		nums2 := []int{}
		n := 0

		merge(nums1, m, nums2, n)

		assert.Equal(t, []int{1}, nums1)
	})
	t.Run("case 3", func(t *testing.T) {
		nums1 := []int{0}
		m := 0
		nums2 := []int{1}
		n := 1

		merge(nums1, m, nums2, n)

		assert.Equal(t, []int{1}, nums1)
	})
	t.Run("case 4", func(t *testing.T) {
		nums1 := []int{2,0}
		m := 1
		nums2 := []int{1}
		n := 1

		merge(nums1, m, nums2, n)

		assert.Equal(t, []int{1,2}, nums1)
	})
	t.Run("case 5", func(t *testing.T) {
		nums1 := []int{1,0}
		m := 1
		nums2 := []int{2}
		n := 1

		merge(nums1, m, nums2, n)

		assert.Equal(t, []int{1,2}, nums1)
	})

}
