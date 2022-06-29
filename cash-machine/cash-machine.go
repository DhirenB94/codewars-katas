package main

import (
	"fmt"
	"math"
)

func main() {
	breakIntoChange(0.13)
}

func breakIntoChange(amount float64) map[string]float64 {
	m := make(map[string]float64)

	fiftyPounds := math.Floor(amount / 50)
	if fiftyPounds > 0 {
		amount = amount - fiftyPounds*50
		m["£50"] = fiftyPounds
	}

	twentyPounds := math.Floor(amount / 20)
	if twentyPounds > 0 {
		m["£20"] = twentyPounds
		amount = amount - twentyPounds*20
	}

	tenPounds := math.Floor(amount / 10)
	if tenPounds > 0 {
		m["£10"] = tenPounds
		amount = amount - tenPounds*10
	}

	fivePounds := math.Floor(amount / 5)
	if fivePounds > 0 {
		m["£5"] = fivePounds
		amount = amount - fivePounds*5
	}

	twoPounds := math.Floor(amount / 2)
	if twoPounds > 0 {
		m["£2"] = twoPounds
		amount = amount - twoPounds*2
	}

	onePound := math.Floor(amount / 1)
	if onePound > 0 {
		m["£1"] = onePound
		amount = amount - onePound*1
	}

	fiftyPence := math.Floor(amount / 0.5)
	if fiftyPence > 0 {
		m["50p"] = fiftyPence
		amount = amount - fiftyPence*0.5
	}

	twentyPence := math.Floor(amount / 0.2)
	if twentyPence > 0 {
		m["20p"] = twentyPence
		amount = amount - twentyPence*0.2
	}

	tenPence := math.Floor(amount / 0.1)
	if tenPence > 0 {
		m["10p"] = tenPence
		amount = amount - tenPence*0.1
	}

	fivePence := math.Floor(amount / 0.05)
	if fivePence > 0 {
		m["5p"] = fivePence
		amount = amount - fivePence*0.05
	}

	twoPence := math.Floor(amount / 0.02)
	if twoPence > 0 {
		m["2p"] = twoPence
		amount = amount - twoPence*0.02
	}

	onePence := math.Round(amount / 0.01)
	if onePence > 0 {
		m["1p"] = onePence
		amount = amount - onePence*0.01
	}
	fmt.Println(m)
	return m

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
