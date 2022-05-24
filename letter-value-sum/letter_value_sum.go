package main

import "fmt"

func main() {
	letterValueSum("microspectrophotometries")

}
func letterValueSum(word string) int {
	nums := []int{}
	for _, letter := range word {
		numberValue := int(letter - 96)
		nums = append(nums, numberValue)
	}
	sum := add(nums)
	fmt.Println(sum)
	return sum
}

func add(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
