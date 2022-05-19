package main

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {
	numbers := []int{32, 15, 88, 2}

	got := SmallestIntegerFinder(numbers)
	want := 2

	if got!= want {
		t.Errorf("got %v, want %v",got, want)
	}
}
