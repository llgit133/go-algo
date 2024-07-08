package main

// 通过切片实现一个栈
// 由于只是辅助实现算法题目，因此不做异常情况处理
type MyStack1 []int

func (s *MyStack1) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack1) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *MyStack1) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack1) Size() int {
	return len(*s)
}

func (s *MyStack1) Empty() bool {
	return s.Size() == 0
}

// ---------- 分界线 ----------

type MyQueue struct {
	stackIn  *MyStack1
	stackOut *MyStack1
}

func Constructor1() MyQueue {
	return MyQueue{
		stackIn:  &MyStack1{},
		stackOut: &MyStack1{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

// fillStackOut 填充输出栈
func (this *MyQueue) fillStackOut() {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)
		}
	}
}
