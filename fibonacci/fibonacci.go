package main

import (
	"fmt"
	"log"
)

//Fibonacci number (Fibonacci sequence), named after mathematician Fibonacci, is a sequence of numbers that looks like this:

//0, 1, 1, 2, 3, 5, 8, 13, ...
//You get first two starting numbers, 0 and 1, and the next number in the sequence is always the sum of the previous two numbers.
//Write a function fib, that takes one parameter n, and returns the nth number from the Fibonacci sequence, counting from 0.
//For example fib(0) returns 0, fib(4) returns 3, fib(15) returns 610.

//Bonus: can you find fib(92)?

func main() {
	fib(4)
}

func fib(n int) int{
	if n < 0 {
		log.Fatalln("Cannot be a negative number")
	}

	fibonacci := []int{0, n -(n-1)}

	for i:= 0; i <n-1; i++ {
		fibonacci = append(fibonacci, fibonacci[i] + fibonacci [i+1])
	}
	fmt.Println(fibonacci[n])
	return fibonacci[n]
}
