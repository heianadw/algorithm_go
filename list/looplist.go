package list

import . "algorithm/public"

//判断单链表是否有环
//时间复杂度O(n)，空间复杂度O(1)
func IsLoop(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head //调用时判断是否==head，如果是，则打印空链表
	}
	slow := head.Next //每次一步
	fast := head.Next //每次两步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

//找出环的入口点
func FindEntranceNodeOfLoop(head *LNode, meetNode *LNode) *LNode {
	first := head.Next
	second := meetNode
	for first != second { //(n-1)*r + 相遇点到入口点的距离 == 起始点到入口点的距离
		first = first.Next
		second = second.Next
	}
	return first //环入口点
}
