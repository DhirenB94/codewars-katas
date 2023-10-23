// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
// pattern contains only lower-case English letters.
// s contains only lowercase English letters and spaces ' '.
// s does not contain any leading or trailing spaces.
// All the words in s are separated by a single space.

package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	pToS := make(map[byte]string)
	sToP := make(map[string]byte)

	stringArray := strings.Split(s, " ")

	if len(pattern) != len(stringArray) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		patternChar, sChar := pattern[i], stringArray[i]

		value, exists := pToS[patternChar]
		if exists && value != sChar {
			return false
		} else {
			pToS[patternChar] = sChar
		}

		v, e := sToP[sChar]
		if e && v != patternChar {
			return false
		} else {
			sToP[sChar] = patternChar
		}
	}
	return true
}
