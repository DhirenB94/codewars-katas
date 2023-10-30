package main

import (
	"fmt"
	"sort"
)

func main() {
	merge([]int{1, 4, 6}, []int{2, 3, 5})
}

func merge(listOne, listTwo []int) []int {
	// output := []int{}
	// i := 0
	// j:= 0

	// for i <len(listOne) && j <len(listTwo) {
	// 	if listOne[i] < listTwo[j] {
	// 		output = append(output,listOne[i])
	// 		i++
	// 	} else {
	// 		output = append(output,listTwo[j])
	// 		j++
	// 	}
	// }

	// fmt.Println(output)

	// for ; i < len(listOne); i++ {
	// 	output = append(output, listOne[i])
	// }
	// for ; j < len(listTwo); j++ {
	// 	output = append(output, listTwo[j])
	// }
	// fmt.Println(output)
	// return output

	for _, v := range listTwo {
		listOne = append(listOne, v)
	}

	sort.Ints(listOne)

	fmt.Println(listOne)
	return listOne
}
