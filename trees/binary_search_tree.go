package trees

import "fmt"

type binarySearchTreeNode struct {
	data  Comparable
	left  *binarySearchTreeNode
	right *binarySearchTreeNode
}

type BinarySearchTree struct {
	root *binarySearchTreeNode
	size int
}

func (bst *BinarySearchTree) Contains(data Comparable) bool {
	return bst.root.contains(data)
}

func (bstn *binarySearchTreeNode) contains(data Comparable) bool {
	if nil == bstn {
		return false
	}

	if data.Compare(bstn.data) < 0 {
		return bstn.left.contains(data)
	} else if data.Compare(bstn.data) > 0 {
		return bstn.right.contains(data)
	}

	return true
}

func (bst *BinarySearchTree) Put(data Comparable) {
	bst.root = bst.root.put(data)
	bst.size += 1
}

func (bstn *binarySearchTreeNode) put(data Comparable) *binarySearchTreeNode {
	if nil == bstn {
		return &binarySearchTreeNode{data, nil, nil}
	}

	// Duplicate data goes to the left of the existing data node
	// when inserted
	if data.Compare(bstn.data) <= 0 {
		bstn.left = bstn.left.put(data)
	} else {
		bstn.right = bstn.right.put(data)
	}

	return bstn
}

func (bst *BinarySearchTree) Remove(data Comparable) error {
	// If the tree contains duplicate entries of "data", only the
	// first occurrence found starting from the root is removed

	if !bst.Contains(data) {
		return fmt.Errorf("%v cannot be removed because it is not in the BinarySearchTree", data)
	}

	bst.root = bst.root.remove(data, false)
	bst.size -= 1
	return nil
}

func (bstn *binarySearchTreeNode) remove(data Comparable, done bool) *binarySearchTreeNode {
	if nil == bstn {
		return nil
	}

	if data.Compare(bstn.data) < 0 || data.Compare(bstn.data) == 0 && done {
		bstn.left = bstn.left.remove(data, done)
	} else if data.Compare(bstn.data) > 0 {
		bstn.right = bstn.right.remove(data, done)
	} else {
		if bstn.left != nil && bstn.right != nil {
			var prev *binarySearchTreeNode
			var curr *binarySearchTreeNode = bstn.left

			for curr != nil && curr.right != nil {
				prev = curr
				curr = curr.right
			}

			if prev != nil && curr != nil {
				prev.right = curr.left
			}

			curr.left = bstn.left.remove(data, true)
			curr.right = bstn.right.remove(data, true)
			return curr
		} else if bstn.left != nil {
			return bstn.left
		} else if bstn.right != nil {
			return bstn.right
		} else {
			return nil
		}
	}

	return bstn
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}
