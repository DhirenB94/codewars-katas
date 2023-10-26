package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	startWindow := intervals[0]
	lastArray := intervals[len(intervals)-1]

	//compare 2nd num of current array with 1st number of next array. If <, then append startwindow 1st number with current array 2nd number. 
	//move start window to next array
	//if not, move to next next array until you find the 2nd number of current array is smaller that 1st num of next array
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] < intervals[i+1][0] {
			result = append(result, []int{startWindow[0], intervals[i][1]})
			startWindow = intervals[i+1]
		}
	}


	//deal with the last, if it is the same as the cureent startWindow, it means its not in any window so append
	//if they are not the same you need to append 1st num of startwindow plus 2nd num of last elem
	if startWindow[0] == lastArray[0] && startWindow[1] == lastArray[1] {
		result = append(result, startWindow)
	} else {
		result = append(result, []int{startWindow[0], lastArray[1]})
	}

	return result
}
