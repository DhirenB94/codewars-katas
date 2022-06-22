package main

import (
	"fmt"
	"math"
	"strings"
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
		leftOfWord, rightOfWord := dotSplitter(totalDots)
		alignedString = leftOfWord + word + rightOfWord

	}
	fmt.Println(alignedString)
	return alignedString

}

func dots(numOfDots int) string {
	var str strings.Builder
	for i := 0; i < numOfDots; i++ {
		str.WriteString(".")
	}
	return str.String()
}

func dotSplitter(s string) (leftOfWord string, rightOfWord string) {
	float := float64(len(s))
	halfLength := int(math.Round(float / 2))

	if halfLength%2 == 0 {
		leftOfWord = s[0:halfLength]
		rightOfWord = s[halfLength:]
	} else {
		leftOfWord = s[0:halfLength]
		rightOfWord = s[halfLength:]
	}
	return leftOfWord, rightOfWord
}
