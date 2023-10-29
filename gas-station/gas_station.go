package main

import "fmt"

func main() {
	canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, currentGas, startStation := 0, 0, 0

    for i := 0; i < len(gas); i++ {
        totalGas += gas[i] - cost[i]
        currentGas += gas[i] - cost[i]

        if currentGas < 0 {
            // If we can't reach the next station, update the starting station to the next station.
            startStation = i + 1
			fmt.Println(startStation)
            currentGas = 0
        }
    }

    // If the total gas is negative, it means it's not possible to complete the circuit. (total gas has to be >= cost)
    if totalGas < 0 {
        return -1
    }

    return startStation
}
