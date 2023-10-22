// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.

package main

func canConstruct(ransomNote string, magazine string) bool {
	lettersMap := make(map[rune]int)

	for _, m := range magazine {
		lettersMap[m]++
	}

	for _, rn := range ransomNote {
		if lettersMap[rn] == 0 {
			return false
		} else {
			lettersMap[rn]--
		}
	}
    return true
}