package xiaohao

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	f := new(ListNode)
	f.Next = head
	nodeMap := make(map[int]*ListNode)
	i := 0
	curr := f
	for curr != nil {
		nodeMap[i] = curr
		curr = curr.Next
		i++
	}

	//fmt.Println(len(nodeMap), len(nodeMap)-n-1, len(nodeMap)-n)
	nodeMap[len(nodeMap)-n-1].Next = nodeMap[len(nodeMap)-n].Next
	return f.Next
}
