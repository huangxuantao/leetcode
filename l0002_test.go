package leetcode

/**
* Definition for singly-linked list.
//* type ListNode struct {
//*     Val int
//*     Next *ListNode
//* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1 []int
	for {
		n1 = append(n1, l1.Val)
		if l1.Next == nil {
			break
		}

		l1 = l1.Next
	}

	var n2 []int
	for {
		n2 = append(n2, l2.Val)
		if l2.Next == nil {
			break
		}

		l2 = l2.Next
	}

	len1 := len(n1)
	len2 := len(n2)
	lenMax := 0
	if len1 > len2 {
		lenMax = len1
	} else {
		lenMax = len2
	}

	var n3 []int
	var j int
	for i := 0; i < lenMax; i++ {
		var v int
		if i < len1 {
			v += n1[i]
		}

		if i < len2 {
			v += n2[i]
		}

		v += j

		if v >= 10 {
			j = 1
			v = v % 10
		} else {
			j = 0
		}

		n3 = append(n3, v)
	}

	if j == 1 {
		n3 = append(n3, 1)
	}


	len3 := len(n3)
	l3 := new(ListNode)
	t := l3
	for i, v := range n3 {
		t.Val = v

		if i == len3-1 {
			break
		}
		t.Next = new(ListNode)
		t = t.Next
	}

	return l3
}
