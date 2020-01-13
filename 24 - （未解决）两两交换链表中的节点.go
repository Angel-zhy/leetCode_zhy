package main

import (
	"fmt"
)

type ListNode5 struct {
	Val int
	Next *ListNode5
}

//1.末尾添加节点
func (h *ListNode5) Append5(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode5{Val:i}
}
//2.打印所有节点
func (h *ListNode5) Show5() {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}

//递归写法(两两交换其中相邻的节点)
func swapPairs(head *ListNode5) *ListNode5 {
	if head == nil || head.Next == nil {
		return  head
	}
	var next = head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	fmt.Println(next.Val)
	return next
}

//p := head  //p为第一个节点
//head = head.Next  //head表头为原链表第二个节点
//p.Next = head.Next  //原链表的第一个节点的Next节点为原链表的第三个节点（即第一、二位节点已经交换）
//head.Next = p
//p.Next = swapPairs(p.Next)  //递归
//return  head

func main()  {
	//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
    //实现链表
    nodeArr := []int{1,2,3,4}
    head := &ListNode5{nodeArr[0],nil}
    for i:=1; i<len(nodeArr); i++ {
    	head.Append5(nodeArr[i])
	}
	//head.Show5()

	swapPairs(head)

    head.Show5()   //1->3->4 （为什么交换后，少了个2节点？？？）

}