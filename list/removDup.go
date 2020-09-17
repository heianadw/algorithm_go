package list

import (
	. "algorithm/public"
)

//顺序删除和递归删除的区别其实就在于一个从前往后删，一个从后往前删。
//顺序删除中的innerPre等于递归删除中的cur变量，innerCur变量相当于pointer变量
func RemoveDup(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	outerCur := head.Next //用于外层循环，指向链表第一个结点
	var innerCur *LNode   //用于内层循环遍历outerCur后面的结点
	var innerPre *LNode   //innerCur的前驱结点
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next //将需删除结点的前一个结点指向需删除结点的下一个结点，从而删除重复数字结点
				innerCur = innerCur.Next      //将内循环指针向后移，然后接着判断后继节点是否与外循环指针指向结点数字重复，重复的话继续删、继续移
			} else {
				innerPre = innerCur      //将内循环指针的前驱结点后移至内循环指针指向的结点
				innerCur = innerCur.Next //内循环指针本身后移到下一个结点
			}
		}
	}
}

func removeDupRecursionChild(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pointer *LNode //子链表的遍历指针
	cur := head        //当前递归层首结点，始终指向Pointer结点的前驱结点，负责重新建立链接
	//对以head.Next为首的子链表删除重复的结点
	head.Next = removeDupRecursionChild(head.Next) //啥也没干，就到最后一个结点返回完事，继续执行倒数第二个结点的递归函数
	pointer = head.Next                            //在倒数第二层递归中，pointer指向最后一个结点，cur指向倒数第二个结点
	//找出以head.Next为首的子链表中与head结点相同的结点并删除
	for pointer != nil {
		if head.Data == pointer.Data {
			cur.Next = pointer.Next //在倒数第二层递归函数中，pointer.Next为nil，相当于断开链接，直连后一个结点
			pointer = cur.Next      //.Next：左赋右指
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return head
}

func RemoveDupRecursion(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	head.Next = removeDupRecursionChild(head.Next) //传递第一个有效结点
}

//空间换时间，通过map变量来匹配重复结点
func RemoveByMap(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	listMap := make(map[interface{}]interface{})
	//cur := head.Next //首先指向第一个结点
	for cur, pre := head.Next, head; cur != nil; cur = cur.Next {
		if _, exist := listMap[cur.Data]; exist {
			pre.Next = cur.Next
		} else {
			listMap[cur.Data] = cur.Data
			pre = cur
		}
	}
}

//删除顺序链表中的重复项
func RemoveSortedListDup(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	for cur := head.Next; cur.Next != nil; {
		if cur.Data == cur.Next.Data {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}
