package main

//For a better solution to achive a 0(1) for get min
//add a second field to the MinStack struck, which would be []int to keep a track of the min numbers added
//when you push, the 1st pushed is added onto both stacks
//the next ones, if the elem to be pushed is <= the last pushed (last elem on min stack), then add to both, if not only add to the main stack
//when popping, if the last elem on the stack is = the last elem on the min stack, pop from both
//now GetMin can just get the last number on the min stack

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}

}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]

	if len(this.stack) > 0 {
		for i := 0; i < len(this.stack); i++ {
			if this.stack[i] < min {
				min = this.stack[i]
			}
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
