package main

import "fmt"

type ListNode2 struct {
	Val int
	Next *ListNode2
}
//1.末尾添加节点
func (h *ListNode2) Append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode2{Val:i}
}
//2.打印所有节点
func (h *ListNode2) Show() {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}
//3.
func removeNthFromEnd(list *ListNode2, n int) *ListNode2 {
	dummy := ListNode2{0,list}
	first := &dummy
	second := &dummy
	// 两个指针之间的距离是n
	for i := 1; i <= n; i++ {   //先走3步（和下一个循环中first.Next有关，用first.Next而不是first）
		first = first.Next
	}
	//fmt.Println(first.Val,second.Val)
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	// 两个指针之间的距离是n+1
	//for i := 1; i <= n+1; i++ {   //先走4步（和下一个循环中first有关）
	//	first = first.Next
	//}
	////fmt.Println(first.Val,second.Val)
	//for first != nil {
	//	first = first.Next
	//	second = second.Next
	//}

	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	list := []int{1,2,3,4,5}
	head := &ListNode2{Val:list[0]}
	tail := head  //需要头尾两个指针
	for i:=1;i<len(list);i++ {
		//方法一 数组直接构建链表
		tail.Next = &ListNode2{Val:list[i]}   //tail.Next => (*tail).Next  golang的隐式转换
		tail = tail.Next
		//head.Append(list[i]) //方法二 追加构建链表
	}
	head.Show()

	//删除倒数第n个节点
	fmt.Println("删除节点后----------")
	removeNthFromEnd(head,4)
	head.Show()
}

