package main

func main() {
	nonogram([]int{0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1})
}

func nonogram(binaryArray []int) []int {
	m := make(map[int]int)
	mapKey := 0
	consecOnes := []int{}

	for _, number := range binaryArray {
		if number == 0 {
			mapKey++
		} else {
			m[mapKey] += 1
		}
	}
	for _, consec := range m {
		consecOnes = append(consecOnes, consec)
	}
	return consecOnes
}
