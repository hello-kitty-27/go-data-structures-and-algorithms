package trees

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

func (bst *BinarySearchTree) Remove(data Comparable) {
	// TODO
	// bst.size -= 1
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}
