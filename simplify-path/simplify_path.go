package main

import (
	"strings"
)

func main() {
	simplifyPath("/a/./b/../../c/")
}

func simplifyPath(path string) string {
	stack := []string{}
	result := ""

	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part == ".." {
			// Go up one directory level by popping from the stack
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
		if part != "" && part != "." && part != ".." {
			stack = append(stack, part)
		}
	}

	result = "/" + strings.Join(stack, "/")
	return result
}
