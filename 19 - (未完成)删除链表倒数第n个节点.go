package main

import "fmt"

//type ListNode struct {
//     Val int
//     Next *ListNode
//}

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node //头节点
}

//链表插入数据
//从尾部添加元素
func (this *List) Append(data Object) {
	node := &Node{Data: data}
	if this.headNode == nil {
		this.headNode = node
		return
	}
	cur := this.headNode
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

func removeNthFromEnd(L *List, n int) *List {
	offset := 0
	h := L.headNode
	pre := &Node{Next:L.headNode}
	preItor := pre
	for h != nil {
		h = h.Next
		if offset < n  {
			offset += 1
			continue
		}
		preItor = preItor.Next
	}
	deleteNode := preItor.Next
	preItor.Next= deleteNode.Next
	deleteNode.Next = nil
	return pre.Next
}


func main()  {
    //初始化链表
	list := List{}
	//1.往单链表末尾追加元素2, 3, 4, 5
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	 node := list.headNode
	 for node.Next != nil {
	 	fmt.Println(node.Data)
	 	node = node.Next
	 }

	list = *removeNthFromEnd(&list, 3)
}

