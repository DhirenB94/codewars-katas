package main

import (
	"math"
)

func main() {
	breakIntoChange(0.03)
}

func breakIntoChange(amount float64) map[string]float64 {
	m1 := make(map[string]float64)
	//m2 := make(map[string]float64)

	amount, maxDivisibleAmount := divide(amount, 50)
	if maxDivisibleAmount != 0 {
		m1["£50"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 20)
	if maxDivisibleAmount != 0 {
		m1["£20"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 10)
	if maxDivisibleAmount != 0 {
		m1["£10"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 5)
	if maxDivisibleAmount != 0 {
		m1["£5"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 2)
	if maxDivisibleAmount != 0 {
		m1["£2"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 1)
	if maxDivisibleAmount != 0 {
		m1["£1"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 0.5)
	if maxDivisibleAmount != 0 {
		m1["50p"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 0.2)
	if maxDivisibleAmount != 0 {
		m1["20p"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 0.1)
	if maxDivisibleAmount != 0 {
		m1["10p"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 0.05)
	if maxDivisibleAmount != 0 {
		m1["5p"] = maxDivisibleAmount
	}

	amount, maxDivisibleAmount = divide(amount, 0.02)
	if maxDivisibleAmount != 0 {
		m1["2p"] = maxDivisibleAmount
	}

	if amount <= 0.01 {
		m1["1p"]++
	}

	// for k, v := range m1 {
	// 	if v > 0 {
	// 		m2[k] = v
	// 	}
	// }
	return m1
}

func divide(amount, number float64) (float64, float64) {
	dividedNum := math.Floor(amount / number)
	if dividedNum > 0 {
		amount = amount - dividedNum*number
	}
	return amount, dividedNum
}

//You're writing software for a cash machine (ATM) in the UK.
//Assume that machines have an unlimited number of coins & notes in the following denominations:
//£50, £20, £10, £5, £2, £1, 50p, 20p, 10p, 5p, 2p, 1p

//Write a function breakIntoChange(amount) that takes any non-negative amount in pounds and returns the minimum number of coins and notes representating that amount.
//e.g.
//breakIntoChange(3.45)
//{'£2': 1, '£1': 1, '20p': 2, '5p': 1}

//breakIntoChange(160)
//{'£50': 3, '£10': 1}

//breakIntoChange(0.13)
//{'10p': 1, '2p': 1, '1p': 1}
