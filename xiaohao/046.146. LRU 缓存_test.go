package xiaohao

import "fmt"

type LRUCache struct {
	usedP    int
	capacity int
	dataMap  map[int]int
	usedKey  []int
}

func Constructor(capacity int) LRUCache {
	var l LRUCache
	l.capacity = capacity
	l.dataMap = make(map[int]int, capacity)
	for i := 0; i < capacity; i++ {
		l.usedKey = append(l.usedKey, 0)
	}
	return l
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("get", this.usedP)
	if v, has := this.dataMap[key]; has {
		this.usedKey[this.usedP] = key
		if this.usedP+1 == this.capacity {
			this.usedP = 0
		} else {
			this.usedP++
		}
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, has := this.dataMap[key]; !has {
		delete(this.dataMap, this.usedKey[this.usedP])
	}

	this.dataMap[key] = value
	this.usedKey[this.usedP] = key
	if this.usedP+1 == this.capacity {
		this.usedP = 0
	} else {
		this.usedP++
	}
}
