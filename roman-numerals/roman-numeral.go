package main

import "fmt"

func main() {
	// fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(intToRoman(20))

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

// var intRomansNumnerals = map[int]string{
// 	1: "I",
// 	4: "IV",
// 	5: "V",
// 	9: "IX",
// 	10: "X",
// 	40: "XL",
// 	50: "L",
// 	90: "XC",
// 	100: "C",
// 	400: "CD",
// 	500: "D",
// 	900: "CM",
// 	1000: "M",
// }

func intToRoman(num int) string {
	romNumeral := ""

	for num > 0 {
		if num >= 1000 {
			romNumeral += "M"
			num -= 1000
			continue
		}
		if num >= 900 {
			romNumeral += "CM"
			num -= 900
			continue
		}
		if num >= 500 {
			romNumeral += "D"
			num -= 500
			continue
		}
		if num >= 400 {
			romNumeral += "CD"
			num -= 400
			continue
		}
		if num >= 100 {
			romNumeral += "C"
			num -= 100
			continue
		}
		if num >= 90 {
			romNumeral += "XC"
			num -= 90
			continue
		}
		if num >= 50 {
			romNumeral += "L"
			num -= 50
			continue
		}
		if num >= 40 {
			romNumeral += "XL"
			num -= 40
			continue
		}
		if num >= 10 {
			romNumeral += "X"
			num -= 10
			continue
		}
		if num >= 9 {
			romNumeral += "IX"
			num -= 9
			continue
		}
		if num >= 5 {
			romNumeral += "V"
			num -= 5
			continue
		}
		if num >= 4 {
			romNumeral += "IV"
			num -= 4
			continue
		}
		if num >= 1 {
			romNumeral += "I"
			num -= 1
			continue
		}
	}
	return romNumeral
}
