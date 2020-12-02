package p1

import (
	"errors"
	"fmt"
)

type Value struct {
	Name string
}
type Node struct {
	Value      // Embedded struct
	next, prev *Node
}
type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

// Pop last item from list
func (l *List) Length() int {
	if l.head == nil {
		return 0
	} else {
		cnt := 0
		for n := l.First(); n != nil; n = n.Next() {
			cnt++
		}
		return cnt
	}
}

// Get node by index
func (l *List) Get(index int) *Node {
	n := l.First()
	for i := 0; i < index && n != nil; i++ {
		n.Next()
	}
	return n
}

var errOutOfRange = errors.New("ERROR - index out of range")

// Create new node with value
func (l *List) Insert(index int, v Value) (*List, error) {
	length := l.Length()
	if index > length {
		return nil, errOutOfRange
	} else if index == length {
		return l.Push(v), nil
	} else {
		n := &Node{Value: v}
		if l.head == nil {
			l.head = n // First node
			l.tail = n
		} else {
			node2add := l.Get(index)
			if node2add == nil {
				return nil, errOutOfRange
			} else {
				prev_node := node2add.prev
				n.prev = prev_node
				n.next = node2add
				prev_node.prev = n
			}
		}
	}
	return l, nil
}

// Create new node with value
func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n // First node
	} else {
		l.tail.next = n // Add after prev last node
		n.prev = l.tail // Link back to prev last node
	}
	l.tail = n // reset tail to newly added node
	return l
}

func (l *List) Find(name string) *Node {
	found := false
	var ret *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Value.Name == name {
			found = true
			ret = n
		}
	}
	return ret
}
func (l *List) Delete(name string) bool {
	success := false
	node2del := l.Find(name)
	if node2del != nil {
		fmt.Println("Delete - FOUND: ", name)
		prev_node := node2del.prev
		next_node := node2del.next
		// Remove this node
		prev_node.next = node2del.next
		next_node.prev = node2del.prev
		success = true
	}
	return success
}

var errEmpty = errors.New("ERROR - List is empty")

// Pop last item from list
func (l *List) Pop() (v Value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return v, err
}
