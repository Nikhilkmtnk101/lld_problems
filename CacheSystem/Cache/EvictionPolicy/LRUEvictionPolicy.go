package EvictionPolicy

import "lld_problems/CacheSystem/Algorithms"

type (
	DoublyLinkedList     = Algorithms.DoublyLinkedList
	DoublyLinkedListNode = Algorithms.DoublyLinkedListNode
)

type LRUEvictionPolicy struct {
	mp  map[string]*DoublyLinkedListNode
	dll *DoublyLinkedList
}

func NewLRUEvictionPolicy() EvictionPolicy {
	return &LRUEvictionPolicy{
		mp:  make(map[string]*DoublyLinkedListNode),
		dll: Algorithms.NewDoublyLinkedList(),
	}
}

func (lru *LRUEvictionPolicy) KeyAccessed(key string) {
	node, exists := lru.mp[key]
	if !exists {
		node = Algorithms.NewDoublyLinkedListNode(key)
		lru.mp[key] = node
	} else {
		lru.dll.DetachNode(node)
	}
	lru.dll.InsertNodeAtFront(node)
}

func (lru *LRUEvictionPolicy) Evict() string {
	node := lru.dll.GetLastNode()
	delete(lru.mp, node.Data)
	lru.dll.DeleteNode(node)
	return node.GetData()
}
