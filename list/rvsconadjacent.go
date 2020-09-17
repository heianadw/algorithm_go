package list

import (
	. "algorithm/public"
)

//就地逆序。结点数为奇数，则在cur指向最后一个结点后戛然而止，如果为偶数，则在倒数第二个结点时翻转最后一次
func ReverseTheConAdjacent(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head.Next                    //当前遍历结点
	pre := head                         //当前结点的前驱结点
	var next *LNode                     //当前结点后继结点的后继结点
	for cur != nil && cur.Next != nil { //偶数结束于cur!=nil，奇数结束于cur.Next!=nil
		next = cur.Next.Next //保存第三个结点
		pre.Next = cur.Next  //前驱指向第二个结点
		cur.Next.Next = cur  //第二个结点指向第一个结点
		cur.Next = next      //第一个结点指向第三个结点
		pre = cur            //前驱结点指向当前结点
		cur = next           //当前结点指向后继结点
	}
}
