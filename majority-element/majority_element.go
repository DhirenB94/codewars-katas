package main

//Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

func main() {
	majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
}

func majorityElement(nums []int) int {

	majority := (len(nums) / 2)

	elementCount := make(map[int]int)

	for _, num := range nums {
		count := elementCount[num]
		count++
		elementCount[num] = count
	}

	for key, ec := range elementCount {
		if ec > majority {
			return key
		}
	}
	return 0
}
