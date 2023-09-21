package main
// 
//import (
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestThree(t *testing.T) {
//	t.Run("when given zero, should return nil slice", func(t *testing.T) {
//		got := three(0)
//		want := []int{}
//
//		assert.Equal(t, want, got, "zero is not allowed")
//	})
//	t.Run("when given 3", func(t *testing.T) {
//		got := three(1)
//		want := []int{3, 1}
//
//		assert.Equal(t, want, got, "")
//	})
//
//	t.Run("when given 1, should return 1 only ", func(t *testing.T) {
//		got := three(1)
//		want := []int{1}
//
//		assert.Equal(t, want, got, "")
//	})
//	t.Run("when given 2, should add one", func(t *testing.T) {
//		got := three(2)
//		want := []int{2, 3, 1}
//
//		assert.Equal(t, want, got, "")
//	})
//
//}