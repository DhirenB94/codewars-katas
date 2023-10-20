// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

package main

func isSubsequence(s string, t string) bool {
	// In this code, we use two pointers i and j to iterate through the strings s and t.
	// The idea is to compare characters in both strings, advancing i when a match is found.
	// If i reaches the length of s, it means that we have successfully found all characters of s in the same relative order within t, and we return true. Otherwise, we return false.
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)



	//Other Solution
	
	// count := 0
	// indexPosT := -1

	// for _, valueS := range s {
	// 	for indexT, valueT := range t {
	// 		if valueS == valueT && indexT > indexPosT {
	// 			count++
	// 			indexPosT = indexT
	// 			break
	// 		}
	// 	}
	// }

	// if count == len(s) {
	// 	return true
	// }

	// return false
}

//Sudocode
//take each letter in s and check if it is in T
// add the condition of the order
//when a match is found, break the for loop
