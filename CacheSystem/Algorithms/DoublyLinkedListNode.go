package Algorithms

type DoublyLinkedListNode struct {
	Data string
	prev *DoublyLinkedListNode
	next *DoublyLinkedListNode
}

func NewDoublyLinkedListNode(data string) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Data: data,
	}
}

func (node *DoublyLinkedListNode) GetData() string {
	return node.Data
}
