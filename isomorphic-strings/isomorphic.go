// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to itself.
// Assume s length and t length are the same and only consist of vali ascii characters

package main

func isIsomorphic(s string, t string) bool {
	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {

		//for each letter in s, check if there is already a value stored in the map.
		//if there is, and it does not match the corresponding t letter in the same position, then there is a discrepency
		//otherwise, store the corresponding t letter in the s map
		sValue, sExists := sToT[s[i]]
		if sExists && sValue != t[i] {
			return false
		} else {
			sToT[s[i]] = t[i]
		}

		//have to do the same for t to s
		tValue, tExists := tToS[t[i]]
		if tExists && tValue != s[i] {
			return false
		} else {
			tToS[t[i]] = s[i]
		}

	}
	return true
}
