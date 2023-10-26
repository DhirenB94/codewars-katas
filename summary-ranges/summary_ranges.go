package main

import (
	"strconv"
)

func main() {
	summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
}

func summaryRanges(nums []int) []string {
	stringArray := []string{}

	if len(nums) == 0 {
		return []string{}
	}

	//keep track of the start of the range window
	startRange := nums[0]
	last := nums[len(nums)-1]

	for i := 0; i < len(nums)-1; i++ {
		//when 2 consecutive numbers dont have a difference of 1, append the start of the window and the end of the window
		//move the start of the new window to the end of the window before + 1
		if nums[i+1]-nums[i] != 1 {
			if startRange == nums[i] {
				stringArray = append(stringArray, strconv.Itoa(startRange))
			} else {
				stringArray = append(stringArray, strconv.Itoa(startRange)+"->"+strconv.Itoa(nums[i]))
			}
			startRange = nums[i+1]
		}
	}

	if startRange == last {
		stringArray = append(stringArray, strconv.Itoa(startRange))
	} else {
		stringArray = append(stringArray, strconv.Itoa(startRange)+"->"+strconv.Itoa(last))
	}

	return stringArray
}
