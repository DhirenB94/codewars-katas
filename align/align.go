package main

import (
	"fmt"
)

//Align (left/right/centre) a string within a given width.
//
//Fill the empty spaces with a dot '.'
//
//e.g. for the string "SaltPay" and width 11
//
//align(“SaltPay”, 11, "L") gives "SaltPay....", aligned to the left
//align(“SaltPay”, 11, "R") gives "....SaltPay", aligned to the right
//align(“SaltPay”, 11, "C") gives "..SaltPay..", aligned in the centre
//
//When centre alignment is not exactly possible, leave more spaces to the left.
//
//e.g. align(”SaltPay”, 12, "C") gives "...SaltPay.."

func main() {
	align("SaltPay", 16, "C")

}

func align(word string, width int, direction string) string {
	numOfDots := width - len(word)
	totalDots := dots(numOfDots)

	var alignedString string

	switch direction {
	case "L":
		alignedString = word + totalDots
	case "R":
		alignedString = totalDots + word
	case "C":
		fewerDots, moreDots := dotSplitter(totalDots)
		alignedString = moreDots + word + fewerDots

	}
	fmt.Println(alignedString)
	return alignedString

}

// dots will return x number of dots as a string e.g 5 = "....."
func dots(numOfDots int) string {
	dots := ""
	for i := 0; i < numOfDots; i++ {
		dots = dots + "."
	}
	return dots

	// var str strings.Builder
	// for i := 0; i < numOfDots; i++ {
	// 	str.WriteString(".")
	// }
	// return str.String()
}

func dotSplitter(s string) (fewerDots, moreDots string) {
	halfLength := len(s) / 2
	//rounds down if a decimal eg 4/2 = 2 and 5/2 =2

	//you want less[x:y] = numbers from x to y but NOT including y
	fewerDots = s[0:halfLength]
	//more on the left [y:] INCLUDING y till the end
	moreDots = s[halfLength:]
	//they will be the same in the case of an even number

	// float := float64(len(s))
	// halfLength := int(math.Round(float / 2))
	return fewerDots, moreDots
}
