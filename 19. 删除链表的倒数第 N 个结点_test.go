package leetcode

import (
	"strconv"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 1
	next := head.Next
	for {
		if next != nil {
			l++
			next = next.Next
		} else {
			break
		}
	}

	count := 1
	current := head
	next = head.Next
	for {
		if l-n == 0 {
			return head.Next
		}

		if count == l-n {
			current.Next = next.Next
			break
		}

		if next != nil {
			count++
			current = next
			next = next.Next
		} else {
			break
		}
	}

	return head
}

func generateListNode(head *ListNode, n int) {
	if head.Val < n {
		head.Next = new(ListNode)
		head.Next.Val = head.Val + 1
		generateListNode(head.Next, n)
	}
}

func printListNode(head *ListNode) string {
	var ret string

	ret += strconv.Itoa(head.Val)
	next := head.Next
	for {
		if next != nil {
			ret += strconv.Itoa(next.Val)
			next = next.Next
		} else {
			break
		}
	}

	return ret
}

func Test_removeNthFromEnd(t *testing.T) {
	header := new(ListNode)
	header.Val = 1
	generateListNode(header, 5)
	t.Log(printListNode(header))

	t.Log(printListNode(removeNthFromEnd(header, 2)))
}
