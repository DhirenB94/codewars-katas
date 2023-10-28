package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	//approach like long addition, if the last number plus one gives a carrt (so if the last number is a 9)
	//then set that number as 0 and carry the one to the number before (continue loop and set carry as 1)
	//keep doing this until there is no carry e.g 8499, at the 4 adding one no longer gives a carry so break out of loop and set carry as 0
	//if outside the loop there is still a carry, it means you have gone up by a factor of 10 (e.g 9999 --> 10000)so set the 1st number as 0 and append 1 to the start
	carry := 0

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i]++
			carry = 0
			break
		}
	}

	if carry == 1 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits

}
