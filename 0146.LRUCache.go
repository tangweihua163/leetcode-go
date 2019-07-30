package leetcode_go

type LRUCache struct {
	m        map[int]*doubleLinkedListNode
	list     doubleLinkedList
	capacity int
	count    int
}

type doubleLinkedList struct {
	head *doubleLinkedListNode
	tail *doubleLinkedListNode
}

type doubleLinkedListNode struct {
	key  int
	val  int
	pre  *doubleLinkedListNode
	next *doubleLinkedListNode
}

func (list *doubleLinkedList) insert(p *doubleLinkedListNode) {
	if list.head == nil {
		list.head = p
		list.tail = p
	} else {
		p.next = list.head
		list.head.pre = p
		list.head = p
	}

}

func (list *doubleLinkedList) delete(p *doubleLinkedListNode) {
	if list.head == nil {
		return
	}

	if list.head == list.tail {
		if list.head == p {
			list.head = nil
			list.tail = nil
		}
		return
	}

	if p == list.head {
		list.head = list.head.next
		list.head.pre = nil
		p.next = nil
		return
	}

	if p == list.tail {
		list.tail = list.tail.pre
		list.tail.next = nil
		p.pre = nil
		return
	}

	p.pre.next = p.next
	p.next.pre = p.pre
	p.next = nil
	p.pre = nil

}

func (list *doubleLinkedList) evict() (bool, int) {
	if list.tail == nil {
		return false, 0
	} else if list.tail == list.head {
		k := list.head.key

		list.head = nil
		list.tail = nil
		return true, k
	} else {
		k := list.tail.key

		p := list.tail
		list.tail = list.tail.pre
		list.tail.next = nil
		p.pre = nil
		return true, k
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		count:    0,
		list:     doubleLinkedList{head: nil, tail: nil},
		m:        make(map[int]*doubleLinkedListNode),
	}
}

func (cache *LRUCache) Get(key int) int {

	if cache.count == 0 {
		return -1
	}

	if v, ok := cache.m[key]; ok {
		cache.list.delete(v)
		cache.list.insert(v)
		return v.val
	} else {
		return -1
	}
}

func (cache *LRUCache) Put(key int, value int) {

	if p, ok := cache.m[key]; ok {
		cache.list.delete(p)
		cache.list.insert(p)
		p.val = value
		return
	}

	if cache.count < cache.capacity {
		p := &doubleLinkedListNode{
			key:  key,
			val:  value,
			pre:  nil,
			next: nil,
		}
		cache.list.insert(p)

		cache.m[key] = p

		cache.count++

		return
	} else {

		if suc, k := cache.list.evict(); suc {
			delete(cache.m, k)
		}

		p := &doubleLinkedListNode{
			key:  key,
			val:  value,
			pre:  nil,
			next: nil,
		}

		cache.list.insert(p)
		cache.m[key] = p

	}

}
