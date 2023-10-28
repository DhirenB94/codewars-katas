package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	// total := 0
	// for i := 0; i < len(s); i++ {
	// 	//last pos
	// 	if i == len(s)-1 {
	// 		total = total + romanNumerals[string(s[i])]
	// 		break
	// 	}
	// 	if string(s[i]) == "I" && string(s[i+1]) == "V" {
	// 		total = total + 4
	// 		i++
	// 		continue
	// 	}
	// 	if string(s[i]) == "I" && string(s[i+1]) == "X" {
	// 		total = total + 9
	// 		i++
	// 		continue
	// 	}
	// 	if string(s[i]) == "X" && string(s[i+1]) == "L" {
	// 		total = total + 40
	// 		i++
	// 		continue
	// 	}
	// 	if string(s[i]) == "X" && string(s[i+1]) == "C" {
	// 		total = total + 90
	// 		i++
	// 		continue
	// 	}
	// 	if string(s[i]) == "C" && string(s[i+1]) == "D" {
	// 		total = total + 400
	// 		i++
	// 		continue
	// 	}
	// 	if string(s[i]) == "C" && string(s[i+1]) == "M" {
	// 		total = total + 900
	// 		i++
	// 		continue
	// 	}
	// 	total = total + romanNumerals[string(s[i])]
	// }

	// return total
	if len(s) == 0 {
		return 0
	}

	firstNumeral := romanNumerals[string(s[0])]
	if len(s) == 1 {
		return firstNumeral
	}
	nextNumeral := romanNumerals[string(s[1])]
	if nextNumeral > firstNumeral {
		combined := nextNumeral - firstNumeral
		return combined + romanToInt(s[2:])
	}
	return firstNumeral + romanToInt(s[1:])
}
