package leetcode

import (
	"testing"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret *ListNode

	switch {
	case list1 == nil && list2 == nil:
		return ret
	case list1 == nil:
		ret = new(ListNode)
		ret.Val = list2.Val
		ret.Next = list2.Next
		return ret
	case list2 == nil:
		ret = new(ListNode)
		ret.Val = list1.Val
		ret.Next = list1.Next
		return ret
	default:
		ret = new(ListNode)
		if list1.Val <= list2.Val {
			//fmt.Println(list1.Val, "<=", list2.Val)
			ret.Val = list1.Val
			ret.Next = mergeTwoLists(list1.Next, list2)
		} else {
			//fmt.Println(list1.Val, ">", list2.Val)
			ret.Val = list2.Val
			ret.Next = mergeTwoLists(list1, list2.Next)
		}
	}

	return ret
}

func Test_mergeTwoLists(t *testing.T) {
	list1 := new(ListNode)
	list1.Val = 1
	list1.Next = new(ListNode)
	list1.Next.Val = 2
	list1.Next.Next = new(ListNode)
	list1.Next.Next.Val = 4


	list2 := new(ListNode)
	list2.Val = 1
	list2.Next = new(ListNode)
	list2.Next.Val = 3
	list2.Next.Next = new(ListNode)
	list2.Next.Next.Val = 4

	t.Log(printListNode(list1))
	t.Log(printListNode(list2))
	ret := mergeTwoLists(list1, list2)
	t.Log(printListNode(ret))
}