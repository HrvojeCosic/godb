package utils

import "errors"

type CircularList[T comparable] struct {
	head     *CircularListNode[T]
	capacity uint
	size     uint // current size of the list
}

type CircularListNode[T comparable] struct {
	value T
	next  *CircularListNode[T]
}

func NewCircularList[T comparable](capacity uint) *CircularList[T] {
	return &CircularList[T]{
		head: nil,
		capacity: capacity,
		size: 0,
	}
}

func (list CircularList[T]) Size() uint {
	return list.size
}

func (list *CircularList[T]) insert(val T) (*CircularListNode[T], error) {
	if list.capacity == list.size {
		return nil, errors.New("can not insert, list is at full capacity")
	}
	
	node := CircularListNode[T]{val, list.head}
	list.head = &node
	list.size++

	return &node, nil
}

func (list *CircularList[T]) remove(node *CircularListNode[T]) {
	if (list.head == nil) { return }
	if (list.head.value == node.value) {
		list.head = list.head.next
		list.size--
		return
	}

	current := list.head
	previous := list.head
	for i := 0; i < int(list.size); i++ {
		if ((*node).value == (*current).value) {
			previous.next = current.next
			list.size--
			break
		}
		current = (*current).next
		previous = current
	}
}