package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{
        left: nil,
        data: i,
        right: nil,
    }
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	node := bst

    for node != nil {
        if node.data > i && node.left != nil {
            node = node.left
        } else if node.data < i && node.right != nil {
            node = node.right
        } else {
            break
        }
    }

    if node.data >= i {
        node.left = NewBst(i)
    } else {
        node.right = NewBst(i)
    }
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	return inOrder(bst)
}

func inOrder(node *BinarySearchTree) []int {
    ret := []int{}

    if node == nil {
        return ret
    }

    ret = append(ret, inOrder(node.left)...)
    ret = append(ret, node.data)
    ret = append(ret, inOrder(node.right)...)

    return ret
}
