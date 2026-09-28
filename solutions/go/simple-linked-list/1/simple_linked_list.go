package simplelinkedlist

import "errors"

// Define the List and Element types here.

type List struct {
	headNode *Element
	tailNode *Element
	size     int
}

type Element struct {
	val  int
	next *Element
}

const getSingleTargetElementOpType = 1
const getAllElementsValArrayOpType = 2

var errListEmpty = errors.New("list is empty")

func New(elements []int) *List {
	if len(elements) < 1 {
		listPtr := List{
			headNode: nil,
			tailNode: nil,
			size:     0,
		}

		return &listPtr
	}

	head := Element{
		val:  elements[0],
		next: nil,
	}

	trav := &head

	for i := 1; i < len(elements); i++ {
		temp := Element{
			val:  elements[i],
			next: nil,
		}

		trav.next = &temp
		trav = trav.next
	}

	listPtr := List{
		headNode: &head,
		tailNode: trav,
		size:     len(elements),
	}

	return &listPtr
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	nextNode := Element{
		val:  element,
		next: nil,
	}

	if l.size > 0 {
		l.tailNode.next = &nextNode
	} else {
		l.headNode = &nextNode
	}

	l.tailNode = &nextNode
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errListEmpty
	}

	nodeVal := l.tailNode.val

	if l.size == 1 {
		l.headNode = nil
		l.tailNode = nil
		l.size--
		return nodeVal, nil
	}

	if l.size == 2 {
		l.headNode.next = nil
		l.tailNode = l.headNode
		l.size--
		return nodeVal, nil
	}

	targetNode, _ := traverseToIndex(l, l.size-1, getSingleTargetElementOpType)

	if targetNode == nil {
		return 0, errors.New("how is this even possible")
	}

	l.tailNode = targetNode
	l.tailNode.next = nil
	l.size--
	return nodeVal, nil
}

func (l *List) Peek() (int, error) {
	if l.size == 0 {
		return 0, errListEmpty
	}

	return l.tailNode.val, nil
}

func (l *List) Array() []int {
	_, nodeValsArr := traverseToIndex(l, l.size, getAllElementsValArrayOpType)
	return nodeValsArr
}

func (l *List) Reverse() *List {
	if l.size < 2 {
		return l
	}

	prev := l.headNode
	curr := prev.next
	nextNode := curr.next

	prev.next = nil
	laterTail := prev

	for curr != nil {
		curr.next = prev
		prev = curr
		curr = nextNode
		if nextNode != nil {
			nextNode = nextNode.next
		}
	}

	l.tailNode = laterTail
	l.headNode = prev
	return l
}

func traverseToIndex(l *List, id, opType int) (*Element, []int) {
	if id < 1 || id > l.size {
		return nil, nil
	}

	trav := l.headNode
	var valArr []int
	var targetId int

	if opType == getAllElementsValArrayOpType {
		targetId = id
	} else {
		targetId = id - 1
	}

	for i := 1; i <= targetId; i++ {
		if opType == getAllElementsValArrayOpType {
			valArr = append(valArr, trav.val)
		}

		trav = trav.next
	}

	if opType == getSingleTargetElementOpType {
		return trav, nil
	}

	return nil, valArr
}
