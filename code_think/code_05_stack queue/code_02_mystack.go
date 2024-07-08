package main

type MyStack2 struct {
	//创建两个队列
	queue1 []int
	queue2 []int
}

func Constructor() MyStack2 {
	return MyStack2{ //初始化
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyStack2) Push(x int) {
	//先将数据存在queue2中
	this.queue2 = append(this.queue2, x)
	//将queue1中所有元素移到queue2中，再将两个队列进行交换
	this.Move()
}

func (this *MyStack2) Move() {
	if len(this.queue1) == 0 {
		//交换，queue1置为queue2,queue2置为空
		this.queue1, this.queue2 = this.queue2, this.queue1
	} else {
		//queue1元素从头开始一个一个追加到queue2中
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:] //去除第一个元素
		this.Move()                   //重复
	}
}

func (this *MyStack2) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:] //去除第一个元素
	return val

}

func (this *MyStack2) Top() int {
	return this.queue1[0] //直接返回
}

func (this *MyStack2) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
