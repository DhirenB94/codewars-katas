package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[rune]int)

	for _, ss := range s {
		hashMap[ss]++
	}

	for _, tt := range t {
		if hashMap[tt] == 0 {
			return false
		} else {
			hashMap[tt]--
		}
	}

	return true
}
