package main

import (
	"fmt"
)

type ListNode6 struct {
	Val int
	Next *ListNode6
}

//1.末尾添加节点
func (h *ListNode6) Append6(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode6{Val:i}
}
//2.打印所有节点
func (h *ListNode6) Show6() {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}

//思路：将k个组合反转，拆分为一个一个的单链表翻转
func reverseKGroup(head *ListNode6, k int) *ListNode6 {
	head = &ListNode6{Next: head}  //增加一个前节点
	pre, start, end := head, head.Next, head
	for end.Next != nil {
		for i:=0; i<k && end != nil; i++ {
			end = end.Next   //end前进k
		}
		if end == nil {
			break
		}
		next := end.Next   //先把第k+1个后面的寄存下来
		end.Next = nil     //断开第k+1个后面的节点
		pre.Next = revese(start)   //翻转长度为k的链表
		start.Next = next    //start为k组合节点翻转之前的第一个节点，翻转之后变为最后一个节点；
		                     // 在最后一个节点后面拼接上刚才寄存下来的节点next
		pre, start, end = start,next,start   //重新赋值
	}
	return head.Next
}


//单链表反转（重点！！！）
func revese(head *ListNode6) *ListNode6{
    var pre *ListNode6
    cur := head
    for cur != nil {
    	cur.Next, pre, cur = pre, cur, cur.Next   //这句话非常重要！！！
	}
	return pre
}

func main()  {
	//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
	//k 是一个正整数，它的值小于或等于链表的长度。
	//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
	//示例：
	//给定这个链表：1->2->3->4->5
	//当 k = 2 时，应当返回: 2->1->4->3->5
	//当 k = 3 时，应当返回: 3->2->1->4->5

	list := []int{1,2,4,15,56,89,100}
	head := &ListNode6{Val:list[0]}

	for i:=1;i<len(list);i++ {
		head.Append6(list[i])
	}
	//head.Show6()

	//单链表翻转
	//head = revese(head)
	//head.Show6()

	//k个一组翻转链表
	head = reverseKGroup(head,3)
	head.Show6()
}
