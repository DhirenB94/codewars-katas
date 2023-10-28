package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	spacesSplit := strings.Split(s, " ")

	for i := len(spacesSplit) - 1; i >= 0; i-- {
		if spacesSplit[i] != "" {
			return len(spacesSplit[i])
		}
	}

	return 0

}
