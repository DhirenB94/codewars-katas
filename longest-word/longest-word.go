package main

import (
	"strings"
	"unicode"
)

//Create a function that will take a parameter being passed and return the longest word in the string
//If there are 2 words with the same length, return only the first word
//ignore punctuation/numbers for the word length but these are viable inputs

//Then, take the output of this function, combine it with this challengeToken - uopxzfve15c, in reverse order, seperated by a colon
// eg input= fun&!! time -- output = time, final output = emit:c51evfzxpou

func longestWord(input, challengeToken string) (string, string) {
	longestWord := ""

	seperatedWords := strings.Split(input, " ")
	formattedWords := []string{}

	for _, sw := range seperatedWords {
		formattedWords = append(formattedWords, removeSpecialChars(sw))
	}

	if input == "" {
		return longestWord, ""
	}

	wordLength := 0

	for _, word := range formattedWords {
		if len(word) > wordLength {
			wordLength = len(word)
			longestWord = word
		}
	}

	reversedWord := reverser(longestWord)
	reversedToken := reverser(challengeToken)

	finalOutput := reversedWord + ":" + reversedToken
	return longestWord, finalOutput
}

func removeSpecialChars(chars string) string {
	formattedString := ""
	for _, char := range chars {
		if unicode.IsLetter(char) {
			formattedString += string(char)
		}
	}
	return formattedString
}

func reverser(word string) string {
	reversedWord := ""

	wordLength := len(word)

	for i := wordLength - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	return reversedWord
}
