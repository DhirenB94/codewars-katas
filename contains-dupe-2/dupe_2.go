package main
// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
// Basically are there duplicates within K distance of eachother

// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true

// Example 2:
// Input: nums = [1,0,1,1], k = 1
// Output: true

// Example 3:
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{5, 6, 7, 5, 6, 7}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	//loop through the array, store each number in a map numberFromaArray:indexPosInArray
	//if something already exists for that number lookup, it would represent the 1st occurence of thar numberFromArray and have its indexPos stored
	// check currentIndexPos and idexPos from the map are within k distance.
	//if so, return true, if not, assign the new index pos

	//[1,2,3,1], k = 3

	dupeMap := make(map[int]int)

	for currentIndex, v := range nums {
		mapIndex, ok := dupeMap[v]
		if ok {
			if currentIndex-mapIndex <= k {
				return true
			}
		}
		dupeMap[v] = currentIndex
	}

	return false
}
