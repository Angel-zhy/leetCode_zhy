package main

import "container/list"

func main() {
	// 一.使用队列实现栈的下列操作：
	// push(x) -- 元素 x 入栈
	// pop() -- 移除栈顶元素
	// top() -- 获取栈顶元素
	// empty() -- 返回栈是否为空
	//1.使用标准库container/list包装
}

type MyStack struct {
	*list.List
}

func Constructor() MyStack {
    return MyStack{
		list.New()
	}
}

func (m *MyStack) Push(x int) {
     m.PushBack(x)
}

func (m *MyStack) Pop() int {
     if m.len() == 0 {
		 return -1
	 }
	 return m.Remove(m.Back().(int))
}

func (m *MyStack) Top() int {
    if m.Len() == 0{
		return -1
	}
	return m.Back().Value.(int)
}

func (m *MyStack) Empty() bool {
    return m.Len() == 0
}


//二.使用栈来实现队列
	//https://www.cnblogs.com/TimLiuDream/p/10087768.html	
	/*注意:
		你只能使用队列的基本操作--
	    也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
		你所使用的语言也许不支持队列。
		你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
	    你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
	*/