package main

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice.
// The relative order of the elements should be kept the same.
// If there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.

// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

func removeDuplicate(nums []int) int {
	uniqueChecker := make(map[int]string)
	result := []int{}

	for _, num := range nums {
		if uniqueChecker[num] == "once" {
			result = append(result, num)
			uniqueChecker[num] = "twice"
		}
		if uniqueChecker[num] == "" {
			result = append(result, num)
			uniqueChecker[num] = "once"
		}
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	uniqueInNums := nums[:len(result)]
	return len(uniqueInNums)

}
