package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	lowercase := strings.ToLower(s)
	spacesRemoved := strings.ReplaceAll(lowercase, " ", "")
	nonAlphaNumericsremoved := removeNonAlphaNumeric(spacesRemoved)

	i := 0
	j := len(nonAlphaNumericsremoved) - 1

	for i < j {
		if nonAlphaNumericsremoved[i] != nonAlphaNumericsremoved[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func removeNonAlphaNumeric(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := reg.ReplaceAllString(input, "")
	return result
}

//Sudocode
//format sentence
//find the halfway point and then compare by going forwards and backwards
//as soon as they are not equal, return false
//2 point solution
