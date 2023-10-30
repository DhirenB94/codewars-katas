package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1,3}))
}

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	fmt.Println(nums)
	
	streak := 1
	maxStreak := 1

	for i := 0; i < len(nums) -1; i++ {
		if nums[i] + 1 == nums[i+1] {
			streak ++
		} else {
			if streak > maxStreak {
				maxStreak = streak
			}
			streak = 1 
		}
		if streak > maxStreak {
			maxStreak = streak
		}
	}

	return maxStreak

	// numSet := make(map[int]bool)

	// // Create a set of unique numbers for efficient lookups
	// for _, num := range nums {
	// 	numSet[num] = true
	// }
	// fmt.Println(numSet)

	// longestStreak := 0

	// for num := range numSet {
	// 	fmt.Println("range", num)
	// 	if !numSet[num-1] {
	// 		currentNum := num
	// 		currentStreak := 1

	// 		// Check for consecutive elements greater than the current number
	// 		for numSet[currentNum+1] {
	// 			currentNum++
	// 			currentStreak++
	// 		}

	// 		longestStreak = max(longestStreak, currentStreak)
	// 	}
	// }

	// return longestStreak
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
