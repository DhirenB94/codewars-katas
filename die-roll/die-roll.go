//Roll a six-sided die a million times. Print the frequency of each side coming up.
//
//Note: Run the program a few times. Do you get the same numbers every time?

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	count1, count2, count3, count4, count5, count6 := 0, 0, 0, 0, 0, 0
	for i := 0; i < 1_000_000; i++ {
		min := 1
		max := 6
		roll := rand.Intn(max-min+1) + min
		switch roll {
		case 1:
			count1++
		case 2:
			count2++
		case 3:
			count3++
		case 4:
			count4++
		case 5:
			count5++
		case 6:
			count6++
		}
	}

	fmt.Printf("Number One rolled %v times\n", count1)
	fmt.Printf("Number Two rolled %v times\n", count2)
	fmt.Printf("Number Three rolled %v times\n", count3)
	fmt.Printf("Number Four rolled %v times\n", count4)
	fmt.Printf("Number Five rolled %v times\n", count5)
	fmt.Printf("Number Six rolled %v times\n", count6)
}
