package list

import (
	. "algorithm/public"
	"fmt"
)

//就地逆序，也即直接逆序
func DirectReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre, cur *LNode //nil
	//if pre == nil {
	//	fmt.Println("****************************************************")
	//}
	/*这个思路跟讲解的图示一点也不一样
	首先将头结点后的第一个结点设置成next结点，整个操作也是围绕着next结点来进行的，
	先将next结点指向的下一个结点置成cur结点，相当于保存了要操作的next结点的后续结点，
	然后将next结点指向的后续结点改成指向next结点的前驱结点，第一次时将next结点的后续结点置成nil，
	即当成尾结点，没有后续结点。接着将next结点本身置成pre结点，并将next结点的原后续结点改成next结点，
	也即cur结点。所以事实上是先移动cur，再移动next与cur指向同一结点。周而复始，循环往复，直到
	cur结点和next结点为nil，且pre结点为逆序前的尾结点为止。跳出循环，并将头结点指向pre结点，即
	原尾结点，完成整个链表的逆序。
	*/
	next := node.Next //头结点后的第一个结点
	for next != nil {
		cur = next.Next //将第二个结点置成当前结点
		next.Next = pre //将第一个结点的后续结点指向前驱结点(第一次置成nil)
		pre = next      //将第一个结点置成前驱结点
		next = cur      //将next(第一个结点)指向当前结点(第二个结点)
	}
	node.Next = pre //处理头结点
}

//递归逆序
func recursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node //直到最后一个结点的next为空，则返回最后一个结点
	}
	newHead := recursiveReverseChild(node.Next)
	node.Next.Next = node //在node表示第一个结点的时候，将第二个结点的next指向node，完成最后一个逆序的连接
	node.Next = nil       //在node表示第一个结点的时候，将node的next赋值为nil，即成为逆序后链表的尾结点
	return newHead        //返回最后一个结点，直到第一个结点
}
func RecursiveReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	firstNode := node.Next
	//递归函数调用
	newHead := recursiveReverseChild(firstNode)
	node.Next = newHead //将头结点的next指向最后一个结点，完成逆序
}

func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur, next *LNode //当前结点，后继结点
	cur = node.Next.Next
	node.Next.Next = nil //设置链表第一个结点为尾结点
	//把遍历到的结点插入到头结点后面
	for cur != nil {
		next = cur.Next //第二个结点开始，保存后继节点为next //保存第三个结点的后继节点
		//fmt.Println(next.Data)
		cur.Next = node.Next //将第二个结点的next指向第一个结点 //将第三个结点的next指向第二个结点
		node.Next = cur      //将头结点指向第二个结点 //将头结点指向第三个结点
		cur = next           //将第三个结点设置成当前结点，继续下一轮操作 //将第四个结点设置成当前结点，继续下一轮
		//fmt.Println(cur.Data)
	}
}

func ReversePrint(node, head *LNode) { //参数非头结点,head只用来判断打印——>
	if node == nil {
		return
	}
	ReversePrint(node.Next, head)
	if node != head.Next {
		fmt.Printf("%d——>", node.Data)
	} else {
		fmt.Println(node.Data)
	}
}
