package list

import (
	. "algorithm/public"
	"errors"
	"fmt"
)

var emptyListErr = errors.New("空链表")
var invalidKErr = errors.New("元素取值不合法")
var theKOutOfBoundsErr = errors.New("k值超出元素个数")

//遍历两次找到k元素//不考虑有无头结点，都从第一个元素开始计数
//时间复杂度O(n),空间复杂度O(1)
func FindTheKByCountBackwards(head *LNode, k int) (interface{}, error) {
	if head == nil || head.Next == nil {
		return nil, emptyListErr
	}
	if k <= 0 {
		return nil, invalidKErr
	}
	n := 0           //结点总数，排除头结点
	cur := head.Next //遍历指针 //这里从第一个结点开始，保证n计数时cur处于第一个结点的位置，移动后不nil才计数
	for cur != nil {
		cur = cur.Next //第一个结点
		n++
	}
	j := n - k + 1 //正数第j个结点等于倒数第k个结点
	if j <= 0 {
		return nil, theKOutOfBoundsErr
	}
	cur = head //准备二次遍历//和cur!=nil的判断条件有所区别，这里从头结点开始遍历，循环几次移动到第几个元素
	//总之，计算结点数从第一个结点开始，找第几个结点从头结点开始
	for i := 0; i < j; i++ {
		cur = cur.Next //循环结束时cur指向倒数第k个结点
	}
	return cur.Data, nil
}

//快慢指针法//无论有无头结点，始终将快指针移动k个结点即可，有头结点则指向第k个结点，无头结点则指向第k+1个结点，后面的判断条件不变，for fast!=nil
//有无头结点只要保证快慢指针起始位置一致即可，同时指向头结点或者同时指向第一个结点
//时间复杂度是O(n)，因为时间不会随着结点的增加而使增长率改变，空间复杂度是O(1)
func FindTheKByTwoPointer(head *LNode, k int) (interface{}, error) {
	if head == nil || head.Next == nil {
		return nil, emptyListErr
	}
	if k <= 0 {
		return nil, invalidKErr
	}
	var fast, slow *LNode //快慢指针
	fast, slow = head, head
	//将快指针先移动k个元素
	var i int
	//i和fast的对应关系，紧跟着判断条件来，i=0时fast指向头结点，i=1时fast指向结点1
	for i = 0; i < k && fast != nil; i++ { //为避免k值过大，引起恐慌，加上条件fast!=nil
		fast = fast.Next
	}
	//遍历结束，如果fast==nil则说明k取值超出整个链表长度，则报错返回
	//带头结点的遍历可以如此判断，但是不带头结点的遍历需要用i<k来判断，因为不带头结点时k值如果和链表长度一致，那么fast移动结束后为nil，此时
	//是正常情况，且此时i=k //由于临界点的情况比较特殊，比如有带头结点的情况k值刚好比链表结点个数多1，则fast==nil时i==k，所以
	//fast==nil&&i<k条件并不能满足有无头结点两种情况，最终还是需要分别判断
	if fast == nil {
		return nil, theKOutOfBoundsErr
	}
	//否则的话，将快慢指针同时移动，直到快指针为nil时，慢指针指向的就是倒数第k个元素
	//快指针先将前k个元素排除掉，然后生成快指针镜像指针慢指针，跟随快指针走到链表终点的下一结点即可
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Data, nil //返回此时slow指针指向的结点的数据即可
}

//将单链表向右旋转k个位置//类似于链表逆序中的插入法
//时间复杂度O(n)，空间复杂度O(1)
func RotateRightK(head *LNode, k int) error {
	if head == nil || head.Next == nil {
		return emptyListErr
	}
	if k <= 0 {
		return invalidKErr
	}
	var fast, slow *LNode   //快慢指针，慢指针需要指向倒数第k+1个结点，快指针需要指向最后一个结点
	fast, slow = head, head //同查找倒数第k个元素一样，从head起步
	//fast指针依然先行k步
	for i := 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return theKOutOfBoundsErr
	}
	if fast.Next == nil { //如果fast正好是最后一个结点，则说明链表整个右旋，也就不用旋转
		fmt.Println("单链表整个右旋，不变")
		return nil
	}
	//移动slow指针至倒数第k+1个结点
	for fast.Next != nil { //此时的判断条件变为fast.Next
		slow = slow.Next
		fast = fast.Next
	}
	//移动结束后，fast指向最后一个结点，此时将fast.Next指向第一个结点，将head指向第k个结点，将slow.Next置nil
	fast.Next = head.Next
	head.Next = slow.Next
	slow.Next = nil
	return nil
}
