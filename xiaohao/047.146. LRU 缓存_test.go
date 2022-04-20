package xiaohao

import "fmt"

type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := new(LinkNode)
	tail := new(LinkNode)
	head.next = tail
	tail.pre = head

	return LRUCache{
		m:    make(map[int]*LinkNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	defer func() {
		c := this.head.next
		for {
			if c.key != 0 {
				fmt.Print(c.key)
			}
			c = c.next
			if c == nil {
				break
			}
		}
		fmt.Println("")
	}()
	if v, has := this.m[key]; has {
		v.pre.next = v.next
		v.next.pre = v.pre

		v.pre = this.head
		v.next = this.head.next

		this.head.next.pre = v
		this.head.next = v

		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	defer func() {
		c := this.head.next
		for {
			if c.key != 0 {
				fmt.Print(c.key)
			}
			c = c.next
			if c == nil {
				break
			}
		}
		fmt.Println("")
	}()
	if v, has := this.m[key]; !has {
		if len(this.m) == this.cap {
			fmt.Println("delete", this.tail.pre.key)
			delete(this.m, this.tail.pre.key)
			this.tail.pre.pre.next = this.tail
			this.tail.pre = this.tail.pre.pre
		}

		this.m[key] = &LinkNode{
			key:   key,
			value: value,
		}

		this.m[key].pre = this.head
		this.m[key].next = this.head.next
		this.head.next.pre = this.m[key]
		this.head.next = this.m[key]

	} else {
		v.value = value

		v.pre.next = v.next
		v.next.pre = v.pre

		v.pre = this.head
		v.next = this.head.next

		this.head.next.pre = v
		this.head.next = v
	}
}
