// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

package main

import (
	"fmt"
	"sort"
)

func main() {
	answer := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(answer)
}

func groupAnagrams(strs []string) [][]string {
	// Initialize a map to store anagrams.
	anagramMap := make(map[string][]string)

	// Iterate through the input strings.
	for _, str := range strs {
		// Convert the string to a byte slice to be able to sort it.
		strBytes := []byte(str)
		// Sort the characters in the string.
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		// Convert the sorted byte slice back to a string.
		sortedStr := string(strBytes)

		// Check if the sorted string exists in the map.
		if _, ok := anagramMap[sortedStr]; !ok {
			// If not, create an entry with an empty slice.
			anagramMap[sortedStr] = []string{}
		}

		// Append the original string (unsorted) to the slice associated with the sorted string.
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// Create a result slice to store the grouped anagrams.
	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
