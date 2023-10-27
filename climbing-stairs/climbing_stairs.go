package main

import "fmt"

func main() {
	climbStairs(5)
}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	stairCombos := []int{0, n - (n - 1), n - (n - 2)}

	for i := 1; i < n-1; i++ {
		stairCombos = append(stairCombos, stairCombos[i]+stairCombos[i+1])
	}

	fmt.Println(stairCombos[n])
	return stairCombos[n]
}

//OR

// func climbStairs(n int) int {
//     if n <= 2 {
//         return n
//     }

//     // Create an array to store the number of ways to reach each step.
//     dp := make([]int, n+1)
//     dp[1] = 1
//     dp[2] = 2

//     // Calculate the number of ways for each step from 3 to n.
//     for i := 3; i <= n; i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }

//     return dp[n]
// }
