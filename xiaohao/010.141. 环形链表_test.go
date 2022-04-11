package xiaohao

func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)
	a := head

	for a != nil {
		if _, has := nodeMap[a]; has {
			return true
		}
		nodeMap[a] = true
		a = a.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == head {
			return true
		}

		fast = fast.Next.Next
		head = head.Next
	}

	return false
}