package Algorithms

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList) IsHeadNode(node *DoublyLinkedListNode) bool {
	return node == dll.head
}

func (dll *DoublyLinkedList) IsTailNode(node *DoublyLinkedListNode) bool {
	return node == dll.tail
}

func (dll *DoublyLinkedList) InsertNodeAtFront(node *DoublyLinkedListNode) {
	if dll.IsEmpty() {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
}

func (dll *DoublyLinkedList) GetFirstNode() *DoublyLinkedListNode {
	return dll.head
}

func (dll *DoublyLinkedList) InsertNodeAtBack(data string) {
	node := NewDoublyLinkedListNode(data)
	if dll.IsEmpty() {
		dll.head = node
		dll.tail = node
	} else {
		node.prev = dll.tail
		dll.tail.next = node
		dll.tail = node
	}
}

func (dll *DoublyLinkedList) GetLastNode() *DoublyLinkedListNode {
	return dll.tail
}

func (dll *DoublyLinkedList) DetachNode(node *DoublyLinkedListNode) {
	if node == nil {
		return
	}

	if dll.IsHeadNode(node) {
		dll.head = node.next
	} else {
		node.prev.next = node.next
	}

	if dll.IsTailNode(node) {
		dll.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}

func (dll *DoublyLinkedList) DeleteNode(node *DoublyLinkedListNode) {
	if node == nil {
		return
	}
	dll.DetachNode(node)
	node.prev = nil
	node.next = nil
}
