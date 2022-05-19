package main

func SmallestIntegerFinder(numbers []int) int {
	minimumInteger := numbers[0]

	for _, arrayInteger := range numbers {
		if minimumInteger > arrayInteger {
			minimumInteger = arrayInteger
		}
	}
	return minimumInteger
}
