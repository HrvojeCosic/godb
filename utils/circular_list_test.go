package utils

import (
	"testing"
)

func TestInsert(t *testing.T) {
	list := NewCircularList[int](2)
	
	i1 := 9
	list.Insert(i1) //nolint:all
	if (list.size != 1) {
		t.Errorf("Expected list size to be 1, but got %d", list.size)
	}
	if (list.head.value != i1) {
		t.Errorf("Expected inserted value to be %d, but got %d", i1, list.head.value)
	}

	i2 := 61
	list.Insert(61) //nolint:all
	if (list.head.value != i2) {
		t.Errorf("Expected last inserted value to be %d, but got %d", i2, list.head.value)
	}

	i3 := 981
	list.Insert(i3) //nolint:all
	if (list.size != 2) {
		t.Error("Expected list to reject element after reaching max capacity")
	}
}

func TestRemove(t *testing.T) {
	list := NewCircularList[int](4)
	i1, i2, i3, i4 := 125, 98, 4, 99
	node1, _ := list.Insert(i1)
	node2, _ := list.Insert(i2)
	node3, _ := list.Insert(i3)
	node4, _ := list.Insert(i4)

	list.Remove(node1.value)
	if (list.head != node4 || list.head.next != node3 || list.head.next.next != node2) {
		t.Errorf("Expected list's remaining values to be %d, %d, and %d", node4.value, node3.value, node2.value)
	}

	list.Remove(node3.value)
	if (list.size != 2) {
		t.Errorf("Expected list have 2 elements left, but found %d elements", list.size)
	}
	if (list.head != node4 || node4.next != node2) {
		t.Errorf("Expected list's remaining values to be %d, and %d", node4.value, node2.value)
	}
}

func TestRemoveOnlyElement(t *testing.T) {
	list := NewCircularList[int](1)
	iNode, _ := list.Insert(104)

	list.Remove(iNode.value)
	if (list.head != nil) {
		t.Error("Expected list's head to be nil")
	}
	if (list.size != 0) {
		t.Errorf("Expected list's size to be 0, but got %d", list.size)
	}
}

func TestNext(t *testing.T) {
	list := NewCircularList[int](50)
	iNode1, _ := list.Insert(104)
	next1 := list.Next(list.head)
	if (next1.value != iNode1.value) {
		t.Errorf("Expected next value to be the same as current, but found %d", next1.value)
	}

	iNode2, _ := list.Insert(51)
	next2 := list.Next(list.head)
	if (next2.value != iNode1.value) {
		t.Errorf("Expected next value to be %d, but found %d", iNode1.value, next2.value)
	}

	newFirstVal := 321
	iNode3, _ := list.Insert(newFirstVal)
	next3 := list.Next(iNode1)
	if (next3.value != newFirstVal) {
		t.Errorf("Expected next value to circle back to first element's value (%d), but found %d", newFirstVal, next3.value)
	}

	list.Remove(iNode2.value)
	next4 := list.Next(iNode3)
	if (next4.value != iNode1.value) {
		t.Errorf("Expected next value to be %d after removing iNode2, but found %d", iNode1.value, next4.value)
	}
}