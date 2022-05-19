package main

import "fmt"


//Write a function that merges two sorted lists of the same length into a new sorted list.

//    [1,4,6],[2,3,5] → [1,2,3,4,5,6].

//   You can do this quicker than concatenating them followed by a sort.

func main () {
	fmt.Println(merge([]int{1,4,6},[]int{2,3,5} ))
}

func merge (listOne, listTwo []int) []int {
	output := []int{}
	i := 0
	j:= 0

	for i <len(listOne) && j <len(listTwo) {
		if listOne[i] < listTwo[j] {
			output = append(output,listOne[i])
			i++
		} else {
			output = append(output,listTwo[j])
			j++
		}
	}
	for ; i < len(listOne); i++ {
		output = append(output, listOne[i])
	}
	for ; j < len(listTwo); j++ {
		output = append(output, listTwo[j])
	}
	return output
}