package main

import (
	"strconv"
)

func main() {
	evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
}

func evalRPN(tokens []string) int {
	//move along the array until you reach an operator
	//do topstack-1 operator topstack and put element back onto the stack as the top

	//["2","1","+","3","*"] == 9

	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			newTopELem := stack[len(stack)-2] + stack[len(stack)-1]
			//remove last 2 elems and replace with new
			stack = stack[:len(stack)-2]
			stack = append(stack, newTopELem)
		case "*":
			newTopELem := stack[len(stack)-2] * stack[len(stack)-1]
			//remove last 2 elems and replace with new
			stack = stack[:len(stack)-2]
			stack = append(stack, newTopELem)
		case "/":
			newTopELem := stack[len(stack)-2] / stack[len(stack)-1]
			//remove last 2 elems and replace with new
			stack = stack[:len(stack)-2]
			stack = append(stack, newTopELem)
		case "-":
			newTopELem := stack[len(stack)-2] - stack[len(stack)-1]
			//remove last 2 elems and replace with new
			stack = stack[:len(stack)-2]
			stack = append(stack, newTopELem)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
