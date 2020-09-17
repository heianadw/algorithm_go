package list

import . "algorithm/public"

//找出链表Head的中间节点，把链表从中间断成两个子链表。
//偶数时对半断开，奇数时从中间结点的前一个结点断开
//当结点数为偶数时，无头结点查找中间结点，慢指针停留的地方为第二个中间结点，前一个结点是第一个中间结点
//偶数时，当有头结点的时候，慢指针停留的地方为第一个中间结点，后一个结点为第二个中间结点
//奇数情况下，无论有无头结点，慢指针始终指向中间结点。
func findMiddleNode(head *LNode) *LNode { //此head不是头结点，而是第一个结点，因此为无头结点找中间结点//快指针走奇数13579
	if head == nil || head.Next == nil {
		return head
	}
	fast := head    //每次遍历两个结点的指针 //通常情况下，快指针指向的结点数是慢指针指向结点数的二倍
	slow := head    //每次遍历一个结点的指针
	slowPre := head //慢指针的前驱结点
	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil //分割链表为两部分 //偶数个结点时，
	return slow        //slow成为后半部分链表的头结点
}

//对不带头结点的单链表翻转
func reverse(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *LNode //前驱结点和当前结点
	for head != nil {    //就地逆序:pre,head,next，以head为操作结点
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func Reorder(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur1 := head.Next //前半部分链表的第一个结点
	mid := findMiddleNode(head.Next)
	cur2 := reverse(mid) //后半部分链表逆序后的第一个结点
	var tmp *LNode
	//合并链表
	for cur1.Next != nil { //cur1和cur2各移动了一个结点，为一个循环结束。再重新从各自的第二个结点开始。
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2 //将第二部分的最后一个结点和第一部分的最后一个结点串起来//偶数的时候第二部分剩一个结点，奇数的时候剩两个结点
}

//求带头结点的单链表的中间结点//返回结点指针，而不是数字
func FindMiddleNodeByHavingHead(head *LNode) []*LNode {
	if head == nil || head.Next == nil {
		return nil
	}
	//middle := make([]*LNode,0)
	var fast, slow *LNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { //偶数
		return []*LNode{slow, slow.Next}
	} else { //奇数
		return []*LNode{slow}
	}
	//return middle
}
