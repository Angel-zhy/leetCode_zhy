package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode)AddNode(val int){
	if l.Next == nil {
		l.Next = &ListNode{val, nil}
	}else{
		l.Next.AddNode(val)
	}
}

func (l *ListNode)lPrint(){
	if l.Next != nil {
		l.Next.lPrint()
	}
	fmt.Println(l.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//创建ListNode结构体指针
	result := &ListNode{}
    lastNode := result
    temp :=0
    for l1 != nil || l2!=nil || temp != 0 {
		//相加处理
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			//移位（移动指针）
			l1 = l1.Next
		}else{
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			//移位（移动指针）
			l2 = l2.Next
		}else{
			v2 = 0
		}
		sum := v1+v2 + temp
		temp = sum/10
		bitNum := sum % 10
		//将数据添加到末端节点
		lastNode.Val = bitNum
		//是否准备新节点
		if l1 != nil || l2 != nil || temp != 0 {
			lastNode.Next = &ListNode{}
		}else{
			lastNode.Next = nil
		}
		lastNode = lastNode.Next
	}
	return result
}

func main()  {
     l1 := ListNode{2,nil}
     l1.AddNode(4)
     l1.AddNode(3)

	 l2 := ListNode{5,nil}
	 l2.AddNode(6)
	 l2.AddNode(4)

	 l3 := addTwoNumbers(&l1,&l2)
	 l3.lPrint()
}