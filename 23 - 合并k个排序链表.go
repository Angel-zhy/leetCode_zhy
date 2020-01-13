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


	//方法一：暴力法
	//思路：遍历所有链表，将多有节点的值放入一个数组中；
	//将这个数组排序，然后遍历所有元素得到正确顺序的值。
	//用遍历得到的值，创建一个新的有序链表。
	//时间复杂度：O(N\log N) ，其中 N 是节点的总数目

	//方法二：（逐一两两合并链表）
	//思路：将k个链表简化为两两链表合并
	//时间复杂度： O(kN) ，其中k 是链表的数目
	var sList []*ListNode4
	sList = append(sList,head1,head2,head3)
	L := mergeKLists(sList)
    L.Show4()


	//方法三：贪心算法、优先队列
	//复杂度分析：
	//时间复杂度：O(Nlogk)，这里 N 是这 k 个链表的结点总数，
	//每一次从一个优先队列中选出一个最小结点的时间复杂度是 O(logk)，故时间复杂度为 O(Nlogk)。



}
