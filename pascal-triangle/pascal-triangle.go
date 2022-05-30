package main

import "fmt"

func main() {
	pascal(7)
}
func pascal(rowNumber int) [][]int {
	//make the number of rows, with each row being 1 empty array
	// [[] [] [] [] [] [] []]
	row := make([][]int, rowNumber)
	//make each array have the same number of elements as the row its on --> eg row 7 will have an array with length 7
	//[[0] [0 0] [0 0 0] [0 0 0 0] [0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0 0]]
	for i := 0; i < rowNumber; i++ {
		row[i] = make([]int, i+1)
		//set 1st and last element as 1, on each array, on each row
		//[[1] [1 1] [1 0 1] [1 0 0 1] [1 0 0 0 1] [1 0 0 0 0 1] [1 0 0 0 0 0 1]]
		row[i][0] = 1
		row[i][i] = 1
		//loop through all positions of each array apart from 1st and last, only enters here from row 3 onwards
		for j := 1; j < i; j++ {
			//sum adjacent elements from the row above and replace the zero elements
			row[i][j] = row[i-1][j-1] + row[i-1][j]
		}
		fmt.Println(row[i])
	}

	return row
}
