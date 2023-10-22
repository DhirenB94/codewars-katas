package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets

//1. find all combinations that make 0
//1. eliminate fduplicates

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {

	tripletsArray := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplets := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplets)

					tripletsArray = append(tripletsArray, triplets)
				}
			}
		}
	}
	return uniqueTriplet(tripletsArray)
}

func uniqueTriplet(triplets [][]int) [][]int {
	tripletMap := make(map[[3]int]bool)
	var result [][]int

	for _, triplet := range triplets {
		key := [3]int{triplet[0], triplet[1], triplet[2]}
		if tripletMap[key] == false {
			tripletMap[key] = true
			result = append(result, triplet)
		}
	}

	return result
}
