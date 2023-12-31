package datastructure

type LinkedList struct {
	Head *interface{}
	Tail *interface{}
}

type Node struct {
	Data interface{}
	Next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) AddToFront(data interface{}) {
	ll.Head = &data

}
func (ll *LinkedList) AddToBack(data interface{}) {
	ll.Tail = &data
}
