package main

type duplicate struct {
	presentOnce  bool
	presentTwice bool
}

func removeDuplicate(nums []int) int {
	uniqueChecker := make(map[int]duplicate)
	result := []int{}

	for _, num := range nums {
		value := uniqueChecker[num]
		if value.presentOnce && !value.presentTwice {
			result = append(result, num)
			value.presentTwice = true
		}
		if !value.presentOnce && !value.presentTwice {
			result = append(result, num)
			value.presentOnce = true
		}
		//When you update the value inside the map, you need to assign it back to the map.
		//Maps in Go are passed by reference, so you need to reassign the updated value back to the map.
		uniqueChecker[num] = value
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	uniqueInNums := nums[:len(result)]
	return len(uniqueInNums)

}
