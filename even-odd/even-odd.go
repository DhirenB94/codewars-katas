package main

import "fmt"

func main () {
	fmt.Println(EvenOrOdd(45))
}


func EvenOrOdd(number int) string{
	if number % 2 == 0 {
		return "Even"
	}
	return "Odd"
}
