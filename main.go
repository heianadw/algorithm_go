package main

import (
	"algorithm/list"
	"algorithm/public"
	"fmt"
)

func main() {
	/*链表:1.链表的逆序*/
	fmt.Println("-----------1.1如何实现链表的逆序-----------")
	head11 := public.NewLNode() //不是nil
	//if head == nil {
	//	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
	//} else {
	//	fmt.Println(head)
	//}
	fmt.Println("<1>就地逆序")
	public.CreateList(head11, 10)
	public.PrintNode("逆序前：", head11)
	list.DirectReverse(head11)
	public.PrintNode("逆序后：", head11)

	fmt.Println("<2>递归逆序")
	public.PrintNode("逆序前：", head11)
	list.RecursiveReverse(head11)
	public.PrintNode("逆序后：", head11)

	fmt.Println("<3>插入逆序")
	public.PrintNode("逆序前：", head11)
	list.InsertReverse(head11)
	public.PrintNode("逆序后：", head11)

	head12 := public.NewLNode()
	fmt.Println("<4>从尾到头输出链表")
	public.CreateList(head12, 10)
	//public.PrintNode("顺序链表为：", head1)
	node1 := head12.Next
	fmt.Print("逆序输出为：")
	list.ReversePrint(node1, head12)

	/*链表:2.移除无序链表重复项*/
	fmt.Println("-----------1.2如何从无序链表中移除重复项-----------")
	head21 := public.NewLNode()
	fmt.Println("<1>顺序删除")
	public.CreateListDup(head21, 10)
	public.PrintNode("生成随机链表：", head21)
	list.RemoveDup(head21)
	public.PrintNode("删除后的链表：", head21)

	head22 := public.NewLNode()
	fmt.Println("<2>递归删除")
	public.CreateListDup(head22, 10)
	public.PrintNode("生成随机链表：", head22)
	list.RemoveDupRecursion(head22)
	public.PrintNode("删除后的链表：", head22)

	head23 := public.NewLNode()
	fmt.Println("<3>空间换时间")
	public.CreateListDup(head23, 10)
	public.PrintNode("随机链表：", head23)
	list.RemoveByMap(head23)
	public.PrintNode("删除后的链表：", head23)

	head24 := public.NewLNode()
	fmt.Println("<4>顺序链表的重复项删除")
	public.CreateSortedList(head24, 10)
	public.PrintNode("有序链表：", head24)
	list.RemoveSortedListDup(head24)
	public.PrintNode("删除重复后的链表：", head24)

	/*链表:3.两个单链表数字之和*/
	fmt.Println("-----------1.3如何计算两个单链表所代表的数之和-----------")
	fmt.Println("<1>整数相加法")
	head31, head32 := public.NewLNode(), public.NewLNode()
	err := public.CreateSumList(head31, head32, 19, 18, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	public.PrintNode("链表One：", head31)
	public.PrintNode("链表Two：", head32)
	head33 := list.SumOfListByNum(head31, head32)
	public.PrintNode("链表Sum：", head33)

	fmt.Println("<2>链表相加法")
	head34, head35 := public.NewLNode(), public.NewLNode()
	err = public.CreateSumList(head34, head35, 20, 23, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	public.PrintNode("链表1：", head34)
	public.PrintNode("链表2：", head35)
	head36 := list.SumOfListByAdd(head34, head35)
	public.PrintNode("链表3：", head36)

	/*链表:4.链表重新排序*/
	fmt.Println("-----------1.4如何对链表重新排序-----------")
	fmt.Println("<1>单链表的重新排序")
	head41 := public.CreateSequentialList(9)
	public.PrintNode("创建顺序链表：", head41)
	list.Reorder(head41)
	public.PrintNode("重排后的链表：", head41)

	fmt.Println("<2>求单链表的中间结点")
	head42 := public.CreateSequentialList(19)
	public.PrintNode("创建链表：", head42)
	middle := list.FindMiddleNodeByHavingHead(head42)
	fmt.Print("中间结点为:")
	for _, node := range middle {
		fmt.Printf("%v ", node.Data)
	}
	fmt.Println()

	/*链表:5.求链表倒数第k个元素*/
	fmt.Println("-----------1.5如何找出单链表中的倒数第k个元素-----------")
	fmt.Println("<1>顺序遍历两遍法")
	head51 := public.CreateSequentialList(9)
	public.PrintNode("创建单链表：", head51)
	k := 1
	v, err := list.FindTheKByCountBackwards(head51, k)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("单链表中的倒数第k(%d)个元素是:%v\n", k, v)
	}

	fmt.Println("<2>快慢指针法")
	head52 := public.CreateSequentialList(10)
	public.PrintNode("创建有序单链表：", head52)
	k = 9
	v, err = list.FindTheKByTwoPointer(head52, k)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("单链表中的倒数第k(%d)个元素是：%v\n", k, v)
	}

	fmt.Println("<3>将单链表右旋k个位置")
	head53 := public.CreateSequentialList(19)
	public.PrintNode("创建有序单链表：", head53)
	k = 7
	err = list.RotateRightK(head53, k)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("将单链表右旋k(%d)个位置\n", k)
		public.PrintNode("右旋后的单链表：", head53)
	}

	/*链表:6.检测一个较大的链表是否有环*/
	fmt.Println("-----------1.6如何检测一个较大的单链表是否有环-----------")
	fmt.Println("<1>蛮力法")
	head61 := public.CreateLoopList(9)
	public.PrintLoopList("打印有环的单链表：", head61)

	fmt.Println("<2>快慢指针法")
	head62 := public.CreateLoopList(10)
	public.PrintLoopList("链表为：", head62)
	meetNode := list.IsLoop(head62)
	if meetNode != nil {
		if meetNode == head62 {
			fmt.Println("空链表")
		} else {
			fmt.Println("链表有环")
			entranceNode := list.FindEntranceNodeOfLoop(head62, meetNode)
			fmt.Println("环的入口元素为：", entranceNode.Data)
		}
	} else {
		fmt.Println("链表无环")
	}

	/*链表:7.把链表相邻元素翻转*/
	fmt.Println("-----------1.7如何把链表相邻元素翻转-----------")
	fmt.Println("<1>交换值法") //并不推荐，但知道

	fmt.Println("<2>就地逆序")
	head72 := public.CreateSequentialList(11)
	public.PrintNode("创建顺序链表：", head72)
	list.ReverseTheConAdjacent(head72)
	public.PrintNode("相邻元素翻转：", head72)
}
