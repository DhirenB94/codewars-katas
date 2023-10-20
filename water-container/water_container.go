package main

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

import "fmt"

func main() {
	output := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	fmt.Println("output", output)
}

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		// Calculate the width between the two lines.
		width := right - left

		// Calculate the height of the container, which is determined by the minimum height of the two lines.
		containerHeight := min(height[left], height[right])

		// Calculate the area using the width and height.
		area := width * containerHeight

		// Update the maximum area if the current area is larger.
		if area > maxArea {
			maxArea = area
		}

		// Move the pointers inward, keeping the higher line as it is more likely to form a larger container.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
