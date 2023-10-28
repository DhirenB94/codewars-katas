package main

func main() {
	_ = trailingZeroes(11)

}

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n

	}

	return count
}
