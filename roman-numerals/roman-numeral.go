package main

import "fmt"

func main() {
	fmt.Println(romanNumeralConverter("XXI"))
}

var numeralMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanNumeralConverter(romanNumerals string) int {
	if len(romanNumerals) == 0 {
		return 0
	}

	firstNumeral := numeralMap[string(romanNumerals[0])]
	if len(romanNumerals) == 1 {
		return firstNumeral
	}
	nextNumeral := numeralMap[string(romanNumerals[1])]
	if nextNumeral > firstNumeral {
		combined := nextNumeral - firstNumeral
		return combined + romanNumeralConverter(romanNumerals[2:])
	}
	return firstNumeral + romanNumeralConverter(romanNumerals[1:])
}
