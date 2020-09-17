package list

import (
	. "algorithm/public"
	"fmt"
	"math"
	"strconv"
)

//整数相加法
func SumOfListByNum(headOne, headTwo *LNode) *LNode {
	if headOne == nil || headOne.Next == nil {
		return headTwo
	}
	if headTwo == nil || headTwo.Next == nil {
		return headOne
	}
	//nodeOne := headOne.Next
	numOne := recursiveTraversal(headOne.Next, 0)
	numTwo := recursiveTraversal(headTwo.Next, 0)
	sum := numOne + numTwo
	fmt.Println(numOne, numTwo, sum)
	return SumConvertToList(sum)
}

func recursiveTraversal(node *LNode, count int) uint64 { //count:记录层数
	if node == nil {
		return 0
	}
	count++
	num := recursiveTraversal(node.Next, count)
	return num + uint64(math.Pow10(count-1))*uint64(node.Data.(int))
}

func SumConvertToList(sum uint64) *LNode {
	head := NewLNode()
	cur := head
	sumSlice := make([]byte, 0)
	sumSlice = strconv.AppendUint(sumSlice, sum, 10)
	//fmt.Println(string(sumSlice))
	for i := 0; i < len(sumSlice); i++ {
		cur.Next = NewLNode()
		cur.Next.Data, _ = strconv.Atoi(string(sumSlice[len(sumSlice)-i-1])) //链表赋值从个位开始
		cur = cur.Next
	}
	return head
}

//链表相加法
func SumOfListByAdd(h1, h2 *LNode) *LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	c := 0                   //记录进位
	sum := 0                 //记录两个结点相加的值
	p1 := h1.Next            //遍历h1
	p2 := h2.Next            //遍历h2
	resultHead := NewLNode() //相加后新的链表头结点
	p := resultHead          //指向链表resultHead最后一个结点
	//注意:此时p指针始终在相对p1、p2指针的前一个结点进行操作，并不同步
	for p1 != nil && p2 != nil { //两个链表一一对应的结点操作
		p.Next = NewLNode()                     //指向新创建的相加后和的结点
		sum = p1.Data.(int) + p2.Data.(int) + c //此处加c加的是上一轮的进位值
		p.Next.Data = sum % 10                  //两节点相加和，只保存个位
		c = sum / 10                            //进位，只保存十位，当然，最大也只是1
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	} //同时为nil的情况？下面相当于判断了两次，也可以
	//链表h2比h1长，接下来只需要考虑h2剩余结点的值
	if p1 == nil {
		for p2 != nil {
			p.Next = NewLNode() //指向为存储p2结点和进位相加后数值的新结点
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10 //和的个位
			c = sum / 10           //和的十位
			p = p.Next
			p2 = p2.Next
		}
	}
	//链表h1比h2长，接下来只需要考虑h1剩余结点的值，操作同上
	if p2 == nil {
		for p1 != nil {
			p.Next = NewLNode()
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}
	//处理进位结点
	if c == 1 {
		p.Next = NewLNode()
		p.Next.Data = 1
	}
	return resultHead
}
