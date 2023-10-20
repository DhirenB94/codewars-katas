// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

package main

func isSubsequence(s string, t string) bool {

	count := 0
	indexPosT := -1

	for _, valueS := range s {
		for indexT, valueT := range t {
			if valueS == valueT && indexT > indexPosT {
				count++
				indexPosT = indexT
				break
			}
		}
	}

	if count == len(s) {
		return true
	}

	return false
}

//Sudocode
//take each letter in s and check if it is in T
// add the condition of the order
//when a match is found, break the for loop
