package main

//Create a function repeatString(str)
// It should takes the parameter being passed and determines if there is some substring K that can be repeated N > 1 times to produce the input string exactly as it appears
// Your function should return the longest substring K, and if there is none it shoudl return the string -1
// e.g "abcabcabcabc" should return "abcabc" but "abcabcabc" should return "abc"
//If the input string only contains 1 character then your program should return the string -1

func repeatSubstring(input string) string {
	finalSubstring := "-1"
	allSubstringRepeats := []string{}
	inputLength := len(input)

	if inputLength <= 1 {
		return finalSubstring
	}

	//iterate until half of the length of the input (if no repeats matching input found by this point, no repeats possible)
	for i := 1; i <= inputLength/2; i++ {
		repeatedSubstring := ""

		//get a substring to check for repeats, increasing the length by 1 each time
		substring := input[:i]

		//repeat the substring of length i, y number of times so that i * y = input length
		substringRepeatCount := inputLength / i

		//manually build the full string from the substring
		//this will do 10 loops for a substrinng of length 1, 5 of length 2 etc
		for j := 0; j < substringRepeatCount; j++ {
			repeatedSubstring += substring
		}

		if repeatedSubstring == input {
			allSubstringRepeats = append(allSubstringRepeats, substring)
		}
	}

	//of all the repeats, we want the longest one
	maxSubStringLength := 0

	if len(allSubstringRepeats) != 0 {
		for _, s := range allSubstringRepeats {
			if len(s) > maxSubStringLength {
				maxSubStringLength = len(s)
				finalSubstring = s
			}
		}
		return finalSubstring
	}

	return finalSubstring
}
