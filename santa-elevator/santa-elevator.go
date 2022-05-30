package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readFileLines("input.txt")
	santa(input)
}

func santa(input string) int {
	openBracket := strings.Count(input, "(")
	closedBracket := strings.Count(input, ")")

	floor := openBracket - closedBracket
	fmt.Println("Floor:", floor)

	return floor
}

func readFileLines(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
