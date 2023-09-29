package main

//Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

func main() {
	majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
}

// this soltution 0(n) for time and complexity as its dependent on the length of the array
// space is 0(n) because the map size is deendent on the array size
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

//Could you solve the problem in linear time and in O(1) space?
//Boyer-Moore Voting Algorithm.

//Time is still 0(n) as we have to loop through the array of size n
// but space is 0(1) because it ses only two variables: candidate and count.
// These variables do not depend on the size of the input array.
// They remain constant regardless of the input size.

func majorityElementImproved(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}

	return candidate
}
