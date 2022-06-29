package main

import (
	"fmt"
	"strconv"
	_ "strconv"
	"strings"
	"time"
)

//Implement the SumOfTimes method which takes a list of strings (minutes and seconds) and returns the sum of the times (hours, minutes, and seconds) as a single string.
//
//Example:
//
//Input: "12:32 34:01 15:23 9:27 55:22 25:56"
//
//Output: "2:32:41"

func main() {
	SumOfTimes([]string{"12:32", "34:01", "15:23", "9:27", "55:22", "02:25:56"})
}

func SumOfTimes(input []string) string {
	var singleTime time.Duration

	//format into time.duration
	for _, t := range input {
		splitByColon := strings.Split(t, ":")

		if len(splitByColon) < 3 {
			splitByColon = append([]string{"00"}, splitByColon...)
		}
		add, _ := time.ParseDuration(splitByColon[0] + "h" + splitByColon[1] + "m" + splitByColon[2] + "s")

		//add all the time.durations
		singleTime += add

	}

	//format back into a string
	totalseconds := int(singleTime.Seconds())

	hours := strconv.Itoa(totalseconds / 3600)
	mins := strconv.Itoa(totalseconds / 60 % 60)
	seconds := strconv.Itoa(totalseconds % 60)

	concat := fmt.Sprintf("%s:%s:%s", hours, mins, seconds)

	fmt.Println(concat)
	return concat

}
