package utils

import "testing"

func TestInsert(t *testing.T) {
	list := NewCircularList[int](2)
	
	i1 := 9
	list.insert(i1)
	if (list.size != 1) {
		t.Errorf("Expected list size to be 1, but got %d", list.size)
	}
	if (list.head.value != 1) {
		t.Errorf("Expected inserted value to be %d, but got %d", i1, list.head.value)
	}

	i2 := 61
	list.insert(61)
	if (list.head.value != i2) {
		t.Errorf("Expected last inserted value to be %d, but got %d", i2, list.head.value)
	}

	i3 := 981
	list.insert(i3)
	if (list.size != 2) {
		t.Error("Expected list to reject element after reaching max capacity")
	}
}

func TestRemove(t *testing.T) {
	list := NewCircularList[int](4)
	i1, i2, i3, i4 := 125, 98, 4, 99
	node1, _ := list.insert(i1)
	node2, _ := list.insert(i2)
	node3, _ := list.insert(i3)
	node4, _ := list.insert(i4)

	list.remove(node1)
	if (list.size != 3) {
		t.Errorf("Expected list have 3 elements left, but found %d elements", list.size)
	}
	if (list.head != node4 || list.head.next != node3 || list.head.next.next != node2) {
		t.Errorf("Expected list's remaining values to be %d, %d and %d", node4.value, node3.value, node2.value)
	}
}

func TestRemoveOnlyElement(t *testing.T) {
	list := NewCircularList[int](1)
	iNode, _ := list.insert(104)

	list.remove(iNode)
	if (list.head != nil) {
		t.Error("Expected list's head to be nil")
	}
	if (list.size != 0) {
		t.Errorf("Expected list's size to be 0, but got %d", list.size)
	}
}