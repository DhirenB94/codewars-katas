package main

//Create a function that takes an array of numbers and returns the index at which the numbers stop increasing/decreasing and the oppposite occurs
//if no order change point is found return -1
//the array will contain a minimum of 3 ints

func orderChangePoint(input []int) int {
	arrayLength := len(input)
	if arrayLength < 3 {
		return 0
	}

	for i := 1; i < arrayLength; i++ {
		before := i - 1
		after := i + 1
		current := i

		//last position
		if i == arrayLength-1 {
			if input[before] > input[current] && input[before] > input[before-1] {
				return i
			}
			if input[before] < input[current] && input[before] < input[before-1] {
				return i
			}
			return -1
		}

		//rest of the array
		if input[before] < input[current] && input[after] < input[current] {
			return i
		}
		if input[before] > input[current] && input[after] > input[current] {
			return i
		}
	}
	return -1
}
