package main

import "fmt"

func main() {

	fmt.Println(longSub("absababstabs"))
}

func longSub(substr string) int {
	if len(substr) == 0 {
		return 0
	}

	hashMap := make(map[rune]int)
	streak := 0
	longestStreak := 1

	for _, v := range substr {
		_, ok := hashMap[v]
		if !ok {
			hashMap[v]++
			streak++
		} else {
			hashMap = make(map[rune]int)
			hashMap[v]++
			fmt.Println(streak)
			if streak > longestStreak {
				longestStreak = streak
			}
			streak = 1
		}
		if streak > longestStreak {
			longestStreak = streak
		}
	}
	return longestStreak
}
