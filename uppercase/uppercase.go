package main

import (
	"strings"
)

func main() {
	uppercase("How can mirrors be real if our eyes aren't real")
}

func uppercase(input string) string {
	return strings.Title(input)
}
