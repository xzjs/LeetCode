package main

import "fmt"

type MinStack struct {
	Value int
	Stack []int
}

func Constructor() MinStack {
	return MinStack{0, []int{}}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Value = val
	} else {
		if this.Value > val {
			this.Value = val
		}
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	n := len(this.Stack)
	temp := this.Stack[n-1]
	this.Stack = this.Stack[0 : n-1]
	if temp == this.Value && len(this.Stack) > 0 {
		this.Value = this.Stack[0]
		for _, v := range this.Stack {
			if v < this.Value {
				this.Value = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Value
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
