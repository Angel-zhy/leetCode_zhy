package main

import "fmt"

type ListNode4 struct {
	Val int
	Next *ListNode4
}

//1.末尾添加节点
func (h *ListNode4) Append4(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode4{Val:i}
}
//2.打印所有节点
func (h *ListNode4) Show4() {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}


//合并k个有序连边
func mergeKLists(lists []*ListNode4) *ListNode4 {
	   li := lists[0]
      for i:=1; i<len(lists); i++ {
		  li = mergeTwoLists4(li, lists[i])
	  }
	  return  li
}


//合并2个链表
func mergeTwoLists4(l1 *ListNode4, l2 *ListNode4) *ListNode4 {
	if l1 == nil {
		return  l2
	}else  if l2 == nil {
		return l1
	}else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists4(l1.Next,l2)
		return l1
	}else {
		l2.Next = mergeTwoLists4(l1,l2.Next)
		return l2
	}
}

func main(){
	//合并k个有序列表，并分析复杂度
//链表实现
	list1 := []int{1,2,4,15}
	list2 := []int{1,8,10,15}
	list3 := []int{2,3,5,9,17,23}
	head1 := &ListNode4{Val:list1[0]}
	head2 := &ListNode4{Val:list2[0]}
	head3 := &ListNode4{Val:list3[0]}

	for i:=1;i<len(list1);i++ {
		head1.Append4(list1[i])
	}
	for i:=1;i<len(list2);i++ {
		head2.Append4(list2[i])
	}
	for i:=1;i<len(list3);i++ {
		head3.Append4(list3[i])
	}
	//head1.Show4()
	//head2.Show4()
	//head3.Show4()

	//方法一：
	//思路：将k个链表简化为两两链表合并
	var sList []*ListNode4
	sList = append(sList,head1,head2,head3)

	L := mergeKLists(sList)
    L.Show4()
}
