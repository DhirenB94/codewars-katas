package main

func main() {
	spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
}

const (
	Right = "right"
	Left  = "left"
	Down  = "down"
	Up    = "up"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	bottomEdge := len(matrix) - 1   //total number of arrays
	rightEdge := len(matrix[0]) - 1 //total length of each array
	topEdge := 0
	leftEdge := 0

	current := 0 //keep track of elmeents inserted into the grid
	direction := Right
	x := 0
	y := 0

	fullGrid := len(matrix) * len(matrix[0])
	grid := make([]int, fullGrid)

	for current < fullGrid {
		grid[current] = matrix[y][x] //1st element (0 index) of the grid, is 1st array, 1st pos in the array = 1
		switch direction {
		case Right:
			//when you reach the rightedge, need to start moving down
			//this means all elems of the 1st array have been mapped and we dont want any of these repeated,
			// so we need to move the topEdge downwards by indexing the next array in the matrix array
			if x == rightEdge {
				direction = Down
				y++
				topEdge++
			} else {
				x++
			}
		case Down:
			//when you hit the bottom edge, this means the right most elems of all array have been mapped
			//so you need to bring the right edge in 1 (reduce the lenght of each array)
			if y == bottomEdge {
				direction = Left
				x--
				rightEdge--
			} else {
				y++
			}
		case Left:
			//when you have reach the leftedge, this means that all the elems of the last array have been mapped
			//so you need to move the bottom edge upwards and eliminate the last array by indexing lastarray -1
			if x == leftEdge {
				direction = Up
				y--
				bottomEdge--
			} else {
				x--
			}
		case Up:
			//at this point we have eliminated the 1st array (topedge +1) so we hit the top on array 2
			//this means all the left most elems (array[0]) of each array have been mapped, so we need to move the left edge in by 1
			if y == topEdge {
				direction = Right
				x++
				leftEdge++
			} else {
				y--
			}
		}
		current++
	}
	return grid

}
