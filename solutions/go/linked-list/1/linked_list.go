package linkedlist

import (
	"errors"
	"reflect"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

var errListEmpty = errors.New("list is empty")

type List struct {
	headNode *Node
	tailNode *Node
	size     int
}

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

func NewList(elements ...any) *List {
	if len(elements) == 0 {
		return &List{}
	}

	head := Node{
		Value: elements[0],
		next:  nil,
		prev:  nil,
	}

	trav := &head

	for i := 1; i < len(elements); i++ {
		temp := Node{
			Value: elements[i],
			next:  nil,
			prev:  nil,
		}

		trav.next = &temp
		temp.prev = trav
		trav = trav.next
	}

	return &List{
		headNode: &head,
		tailNode: trav,
		size:     len(elements),
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	pushed := &Node{
		Value: v,
		next:  l.headNode,
		prev:  nil,
	}

	if l.size == 0 {
		l.tailNode = pushed
	} else {
		l.headNode.prev = pushed
	}

	l.headNode = pushed
	l.size++
}

func (l *List) Push(v any) {
	pushed := &Node{
		Value: v,
		next:  nil,
		prev:  l.tailNode,
	}

	if l.size == 0 {
		l.headNode = pushed
	} else {
		l.tailNode.next = pushed
	}

	l.tailNode = pushed
	l.size++
}

func (l *List) Shift() (any, error) {
	if l.size == 0 {
		return nil, errListEmpty
	}

	del := l.headNode.Value

	l.headNode = l.headNode.next
	l.size--

	if l.headNode == nil {
		l.tailNode = nil
	} else {
		l.headNode.prev = nil
	}

	return del, nil
}

func (l *List) Pop() (any, error) {
	if l.size == 0 {
		return nil, errListEmpty
	}

	del := l.tailNode.Value

	l.tailNode = l.tailNode.prev
	l.size--

	if l.tailNode == nil {
		l.headNode = nil
	} else {
		l.tailNode.next = nil
	}

	return del, nil
}

func (l *List) Reverse() {
	if l.size < 2 {
		return
	}

	prev := l.headNode
	curr := prev.next
	nextNode := curr.next
	prev.next = nil
	newTail := prev

	for curr != nil {
		curr.next = prev
		prev.prev = curr
		prev = curr
		curr = nextNode

		if nextNode != nil {
			nextNode = nextNode.next
		}
	}

	prev.prev = nil
	l.tailNode = newTail
	l.headNode = prev
}

func (l *List) First() *Node {
	return l.headNode
}

func (l *List) Last() *Node {
	return l.tailNode
}

func (l *List) Count() int {
	return l.size
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (l *List) Delete(v any) bool {
	if l.size == 0 {
		return false
	}

	if reflect.DeepEqual(l.headNode.Value, v) {
		if _, err := l.Shift(); err != nil {
			return false
		}

		return true
	}

	if reflect.DeepEqual(l.tailNode.Value, v) {
		if _, err := l.Pop(); err != nil {
			return false
		}

		return true
	}

	for tmp := l.headNode.next; tmp != nil && tmp != l.tailNode; tmp = tmp.next {
		if reflect.DeepEqual(tmp.Value, v) {
            tmp.prev.next = tmp.next
            tmp.next.prev = tmp.prev
			tmp.prev = nil
			tmp.next = nil
			l.size--
			return true
		}
	}

	return false
}
