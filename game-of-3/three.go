package main

import "fmt"

//Given any number, repeatedly divide the number by 3 till you reach 1.
//Add or subtract 1 whenever division by 3 is not possible. At each stage output the number.
//e.g. if you start with 100, it’s not divisible by 3, so subtract 1.
//100, 99
//Then keep dividing by 3
//100, 99, 33, 11
//11 is not divisible by 3, so add 1.
//100, 99, 33, 11, 12
//Keep going till you get to 1.
//100, 99, 33, 11, 12, 4, 3, 1

//plus one or minus one when it does so and ouput that
//100 -1

func main() {
	three(100)
}

func three(number int) []int {
	if number <= 2 {
		return nil
	}

	numbers := []int{}

	for number >= 3 {
		//if not divisible by 3, append then subtract
		if number %3 == 1 {
			numbers = append(numbers, number)
			number = number - 1
		} else if number%3 == 2 {
			numbers = append(numbers, number)
			number = number + 1
		} else {
			numbers = append(numbers, number)
			number = number / 3
		}
	}

	numbers = append(numbers, 1)
	fmt.Println(numbers)
	return numbers

	// if number == 0 {
	// 	fmt.Println("zero is not allowed")
	// 	return []int{}
	// }

	// numbers := []int{number}

	// for number > 1 {
	// 	switch {
	// 	case number%3 == 2:
	// 		number = number + 1
	// 		numbers = append(numbers, number)
	// 		continue
	// 	case number%3 == 1:
	// 		number = number - 1
	// 		numbers = append(numbers, number)
	// 		continue
	// 	case number%3 == 0:
	// 		number = number / 3
	// 		numbers = append(numbers, number)
	// 		continue
	// 	}

	// }
	// fmt.Println(numbers)
	// return numbers
}
