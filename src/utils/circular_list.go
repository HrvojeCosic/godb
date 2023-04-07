package utils

import (
	"errors"
)

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

func (cl CircularList[T]) Head() *CircularListNode[T] {
	return cl.head
}

func (node CircularListNode[T]) Value() T {
	return node.value
}

func (list CircularList[T]) Next(current *CircularListNode[T]) *CircularListNode[T] {
	if (list.size == 0) {
		return nil
	}
	if (current.next == nil) {
		return list.head // go around to beginning
	}
	return current.next
}

func (list CircularList[T]) Size() uint {
	return list.size
}

func (list *CircularList[T]) Insert(val T) (*CircularListNode[T], error) {
	if list.capacity == list.size {
		return nil, errors.New("can not insert, list is at full capacity")
	}
	
	node := CircularListNode[T]{val, list.head}
	list.head = &node
	list.size++

	return &node, nil
}

func (list *CircularList[T]) Remove(val T) {
	if (list.head == nil) { return }
	if (list.head.value == val) {
		list.head = list.head.next
		list.size--
		return
	}

	current := list.head.next
	previous := list.head
	for i := 0; i < int(list.size); i++ {
		if (val == (*current).value) {
			previous.next = current.next
			list.size--
			break
		}
		current = (*current).next
		previous = current
	}
}