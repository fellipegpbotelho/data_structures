package main

import (
	"errors"
	"fmt"
)

// Node represents a node of a linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// Insert inserts new node at the end of linked list
func (linkedList *LinkedList) Insert(value int) {
	node := Node{}
	node.value = value
	if linkedList.len == 0 {
		linkedList.head = &node
		linkedList.len++
		return
	}

	item := linkedList.head
	for i := 0; i < linkedList.len; i++ {
		nodeIsTail := item.next == nil
		if nodeIsTail {
			item.next = &node
			linkedList.len++
			return
		}
		item = item.next
	}
}

// GetAt returns node at given position from linked list
func (linkedList *LinkedList) GetAt(position int) *Node {
	node := linkedList.head
	if position < 0 {
		return node
	}
	positionDoNotExists := position > (linkedList.len - 1)
	if positionDoNotExists {
		return nil
	}
	for i := 0; i < position; i++ {
		node = node.next
	}
	return node
}

// InsertAt inserts new node at given position
func (linkedList *LinkedList) InsertAt(position int, value int) {
	newNode := Node{}
	newNode.value = value
	if position < 0 {
		fmt.Println("Position can not be negative")
		return
	}
	if position == 0 {
		linkedList.head = &newNode
		linkedList.len++
		return
	}
	if position > linkedList.len {
		fmt.Println("Position do not exists")
		return
	}
	node := linkedList.GetAt(position)
	newNode.next = node
	previousNode := linkedList.GetAt(position - 1)
	previousNode.next = &newNode
	linkedList.len++
}

// Print displays all the nodes from linked list
func (linkedList *LinkedList) Print() {
	if linkedList.len == 0 {
		fmt.Println("No nodes in list")
		return
	}
	node := linkedList.head
	for i := 0; i < linkedList.len; i++ {
		fmt.Println("Node: ", node)
		node = node.next
	}
}

// Search returns node position with given value from linked list
func (linkedList *LinkedList) Search(value int) int {
	node := linkedList.head
	for i := 0; i < linkedList.len; i++ {
		if node.value == value {
			return i
		}
		node = node.next
	}
	return -1
}

// DeleteAt deletes node at given position from linked list
func (linkedList *LinkedList) DeleteAt(position int) error {
	if position < 0 {
		fmt.Println("Position can not be negative")
		return errors.New("position can not be negative")
	}
	if linkedList.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
	previousNode := linkedList.GetAt(position - 1)
	if previousNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	previousNode.next = linkedList.GetAt(position).next
	linkedList.len--
	return nil
}

// DeleteByValue deletes nodes having given value from linked list
func (linkedList *LinkedList) DeleteByValue(value int) error {
	node := linkedList.head
	if linkedList.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < linkedList.len; i++ {
		if node.value == value {
			if i > 0 {
				previousNode := linkedList.GetAt(i - 1)
				previousNode.next = linkedList.GetAt(i).next
			} else {
				linkedList.head = node.next
			}
			linkedList.len--
			return nil
		}
		node = node.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

func main() {
	linkedList := LinkedList{}

	fmt.Println("\n************* Insert and InsertAt *************")
	linkedList.Insert(12)
	linkedList.Insert(13)
	linkedList.Insert(14)
	linkedList.Insert(15)
	linkedList.InsertAt(4, 16)
	linkedList.InsertAt(5, 17)
	linkedList.InsertAt(6, 18)
	linkedList.InsertAt(7, 19)

	fmt.Println("************* Print *************")
	linkedList.Print()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", linkedList.Search(12))
	fmt.Println("Position of 14 value: ", linkedList.Search(14))
	fmt.Println("Position of 15 value: ", linkedList.Search(15))
	fmt.Println("Position of 100 value: ", linkedList.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", linkedList.GetAt(0))
	fmt.Println("Get At 3rd position: ", linkedList.GetAt(2))
	fmt.Println("Get At 4th position: ", linkedList.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", linkedList.GetAt(-5))

	fmt.Println("\n************* DeleteAt *************")
	linkedList.DeleteAt(3)

	fmt.Println("************* Print *************")
	linkedList.Print()

	fmt.Println("\n************* DeleteByValue *************")
	linkedList.DeleteByValue(14)
	linkedList.DeleteByValue(19)

	fmt.Println("************* Print *************")
	linkedList.Print()
}
