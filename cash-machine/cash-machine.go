package main

import (
	"fmt"
	"math"
)

func main() {
	breakIntoChange(333)
}

func breakIntoChange(amount float64) map[string]float64 {
	m1 := make(map[string]float64)
	m2 := make(map[string]float64)

	amount, fiftyPounds := divide(amount, 50)
	m1["£50"] = fiftyPounds

	amount, twentyPounds := divide(amount, 20)
	m1["£20"] = twentyPounds

	amount, tenPounds := divide(amount, 10)
	m1["£10"] = tenPounds

	amount, fivePounds := divide(amount, 5)
	m1["£5"] = fivePounds

	amount, twoPounds := divide(amount, 2)
	m1["£2"] = twoPounds

	amount, onePound := divide(amount, 1)
	m1["£1"] = onePound

	amount, fiftyPence := divide(amount, 0.5)
	m1["50p"] = fiftyPence

	amount, twentyPence := divide(amount, 0.2)
	m1["20p"] = twentyPence

	amount, tenPence := divide(amount, 0.1)
	m1["10p"] = tenPence

	amount, fivePence := divide(amount, 0.05)
	m1["5p"] = fivePence

	amount, twoPence := divide(amount, 0.02)
	m1["2p"] = twoPence

	onePence := math.Round(amount / 0.01)
	if onePence > 0 {
		m1["1p"] = onePence
		amount = amount - onePence*0.01
	}

	for k, v := range m1 {
		if v > 0 {
			m2[k] = v
		}
	}

	fmt.Println(m2)
	return m2
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
