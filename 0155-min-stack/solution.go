type MinStack struct {
	data [][2]int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	if len(this.data) == 0 {
		this.data = append(this.data, [2]int{value, value})
	} else {
		this.data = append(this.data, [2]int{value, min(value, this.data[len(this.data)-1][1])})
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data) - 1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
