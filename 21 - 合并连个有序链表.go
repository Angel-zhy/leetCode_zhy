package main

import "fmt"

type ListNode3 struct {
	Val int
	Next *ListNode3
}
//1.末尾添加节点
func (h *ListNode3) Append3(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode3{Val:i}
}
//2.打印所有节点
func (h *ListNode3) Show3() {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}

//递归

//3.联合两个有序链表 （ 推荐 ）
//时间复杂度：O(n + m)
//         因为每次调用递归都会去掉 l1 或者 l2 的头元素（直到至少有一个链表为空），
//         函数 mergeTwoList 中只会遍历每个元素一次。
//         所以，时间复杂度与合并后的链表长度为线性关系。
//空间复杂度：O(n + m)
//        调用 mergeTwoLists 退出时 l1 和 l2 中每个元素都一定已经被遍历过了，
//        所以 n + m 个栈帧会消耗 O(n + m) 的空间。
func mergeTwoLists(l1 *ListNode3, l2 *ListNode3) *ListNode3 {
	if (l1 == nil) {   //l1空链表，直接返回另一条
		return l2
	} else if (l2 == nil) {  //l2空链表
		return l1
	} else if l1.Val < l2.Val {   //递归
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//4.联合两个有序链表
func mergeTwoLists2(l1 *ListNode3, l2 *ListNode3) *ListNode3 {
	// 如果有一条为nil， 直接返回另一条
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 定义头head和移动指针node
	var head, node *ListNode3
	// 选小的那条开始
	// 前面判断过l1和l2了，所以这里不可能是空指针
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}
	// 循环比较，同时更新
	// 始终选择小的值连到node上
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 指向小的那条
			node.Next = l1
			// 小的往前走
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		// 【重要】node往前走
		node = node.Next
	}
	// 循环完了，把剩下的直接放到后面，因为本身就是有序的，所以接在后面就行
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	// 返回最初的那个头
	// 这个头不能动，否则找不到了
	return head
}



func main()  {
	list1 := []int{1,2,4,15}
	list2 := []int{1,3,4,6,9,11}
	head1 := &ListNode3{Val:list1[0]}
	head2 := &ListNode3{Val:list2[0]}
	tail1 := head1  //需要头尾两个指针

	for i:=1;i<len(list1);i++ {
		//方法一 数组直接构建链表
		tail1.Next = &ListNode3{Val:list1[i]}   //tail.Next => (*tail).Next  golang的隐式转换
		tail1 = tail1.Next
		//head.Append3(list[i]) //方法二 追加构建链表
	}
	for i:=1;i<len(list2);i++ {
		head2.Append3(list2[i]) //方法二 追加构建链表
	}
	//head1.Show3()
	//head2.Show3()

    mergeList := mergeTwoLists(head1,head2)
    //mergeList := mergeTwoLists2(head1,head2)
	mergeList.Show3()
}
