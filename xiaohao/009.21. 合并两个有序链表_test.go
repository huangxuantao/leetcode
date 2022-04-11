package xiaohao

import (
	"testing"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := new(ListNode)
	ret := preHead

	var l1, l2 *ListNode
	l1 = list1
	l2 = list2
	for {
		if l1 == nil {
			ret.Next = l2
			return preHead.Next
		}

		if l2 == nil {
			ret.Next = l1
			return preHead.Next
		}

		if l1.Val < l2.Val {
			ret.Next = l1
			l1 = l1.Next
		} else {
			ret.Next = l2
			l2 = l2.Next
		}

		ret = ret.Next
	}
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

}