package public

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}

/*打印链表*/
func PrintNode(desc string, head *LNode) {
	fmt.Print(desc)
	if head.Next == nil {
		fmt.Println("空链表")
		return
	}
	curNode := head.Next
	for curNode != nil {
		fmt.Print(curNode.Data)
		if curNode.Next != nil {
			fmt.Print("——>")
		}
		curNode = curNode.Next
	}
	fmt.Println()
}

/*链表:1.链表的逆序*/
func CreateList(head *LNode, nodeNum int) {
	//随机数设置
	rand.Seed(time.Now().UnixNano())
	//listNum := make([]int,0,nodeNum)
	curNode := head
	for i := 0; i < nodeNum; i++ {
		//listNum = append(listNum,rand.Intn(100))
		node := NewLNode()
		node.Data = rand.Intn(100)
		curNode.Next = node
		curNode = node
	}
	curNode = head.Next
	fmt.Print("随机链表：")
	for curNode != nil {
		fmt.Print(curNode.Data)
		if curNode.Next != nil {
			fmt.Print("——>")
		}
		curNode = curNode.Next
	}
	fmt.Println()
}

/*链表:2.移除无序链表重复项*/
func CreateListDup(head *LNode, num int) {
	rand.Seed(time.Now().UnixNano())
	var dupNum int
	node := head
	for i := 0; i < num; i++ {
		node.Next = &LNode{}
		if dupNum != 0 && i%2 == 1 {
			node.Next.Data = dupNum
			dupNum = 0
		} else {
			node.Next.Data = rand.Intn(20)
			if node.Next.Data.(int)%2 == 0 {
				dupNum = node.Next.Data.(int)
			}
		}
		node = node.Next
	}
}

//创建有重复项的顺序链表
func CreateSortedList(head *LNode, num int) {
	rand.Seed(time.Now().UnixNano())
	begin := rand.Intn(100)
	cur := head
	for i := 1; i <= num; i++ {
		cur.Next = NewLNode()
		if i%3 == 0 {
			cur.Next.Data = begin + i - 1
		} else {
			cur.Next.Data = begin + i
		}
		cur = cur.Next
	}
}

/*链表:3.计算两个单链表数字之和*/
func CreateSumList(headOne, headTwo *LNode, numOne, numTwo, flag int) error {
	//如果是整数相加法，则num <= 18，考虑到uint64=18446744073709551615:20位，相加最大19位
	//或者其中之一是19位，另一个是18位
	if flag == 1 {
		if numOne+numTwo > 37 || numOne > 19 || numTwo > 19 {
			return errors.New("整数相加法单链表结点数过多（19结点+18结点为最大满足条件）")
		}
	}
	rand.Seed(time.Now().UnixNano())
	curOne := headOne
	for i := 0; i < numOne; i++ {
		curOne.Next = NewLNode()
		curOne.Next.Data = rand.Intn(10)
		curOne = curOne.Next
	}
	curTwo := headTwo
	//curTwo = nil
	for i := 0; i < numTwo; i++ {
		curTwo.Next = NewLNode()
		curTwo.Next.Data = rand.Intn(10)
		curTwo = curTwo.Next
	}
	return nil
}

/*链表:4.将单链表按一定的规则重新排序*/
//从1开始的顺序链表
func CreateSequentialList(num int) *LNode {
	head := NewLNode()
	cur := head
	for i := 1; i <= num; i++ {
		cur.Next = NewLNode()
		cur.Next.Data = i
		cur = cur.Next
	}
	return head
}

/*链表:6.检测一个较大的链表是否有环*/
//创建一个有环的链表
func CreateLoopList(num int) *LNode {
	if num <= 1 { //1个结点就不要环了
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	loopNum := rand.Intn(num - 1) //num为5时随机数是[0,4)，这样环不会闭在最后一个结点上
	head := NewLNode()
	cur := head
	var loopNode *LNode //用来指向闭环结点
	for i := 0; i < num; i++ {
		cur.Next = NewLNode()
		cur.Next.Data = i + 1
		cur = cur.Next
		//在cur=cur.Next之后判断并赋值，是考虑到有头结点的情况
		if i == loopNum {
			loopNode = cur
		}
	}
	cur.Next = loopNode //将最后一个结点指向闭环结点
	return head
}

//打印有环的链表 //暴力法
func PrintLoopList(desc string, head *LNode) {
	if head == nil || head.Next == nil {
		fmt.Println("空链表")
		return
	}
	fmt.Print(desc)
	nodeSet := make(map[*LNode]bool)
	cur := head
	for !nodeSet[cur.Next] { //当在map中找到相同地址时，退出循环 //键不存在，值为零值false
		fmt.Printf("%v——>", cur.Next.Data)
		nodeSet[cur.Next] = true
		cur = cur.Next
	}
	fmt.Print("环——>")
	//打印环一次
	last := cur               //记录最后一个结点
	nodeSet[cur.Next] = false //将闭环的第一个结点在map中置成不存在，以可以继续遍历
	for !nodeSet[cur.Next] {
		fmt.Printf("%v——>", cur.Next.Data)
		//nodeSet[cur.Next] = true //将闭环第一个节点再置成true，保证循环一遍后停止。——>不行
		cur = cur.Next //指向闭环的第一个结点
		if cur != last {
			nodeSet[cur.Next] = false //将下一个结点继续置成false
		} else {
			fmt.Print("......\n")
			nodeSet[cur.Next] = true //最后一个结点时结束循环
		}
	}
}
