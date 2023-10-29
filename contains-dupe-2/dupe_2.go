package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	//loop through the array, store each number in a map number:indexPos
	//if something already exists for that number lookup, check its within k distance.
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
