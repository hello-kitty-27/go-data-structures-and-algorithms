package trees

import (
	"testing"
)

func TestPut(t *testing.T) {
	var emptyBst BinarySearchTree
	bst := &emptyBst

	bst.Put(comparableInteger{50})

	if !(bst.Contains(comparableInteger{50}) && bst.Size() == 1) {
		t.Error("Put() fails on an empty BST (50)")
	}

	bst.Put(comparableInteger{25})

	if !(bst.Contains(comparableInteger{25}) && bst.Size() == 2 && bst.root.left.data.(comparableInteger).value == 25) {
		t.Error("Put() fails when pushing the second node (25)")
	}

	bst.Put(comparableInteger{22})

	if !(bst.Contains(comparableInteger{22}) && bst.Size() == 3 && bst.root.left.left.data.(comparableInteger).value == 22) {
		t.Error("Put() fails when pushing the third node (22)")
	}

	bst.Put(comparableInteger{22})

	if !(bst.Size() == 4 && bst.root.left.left.left.data.(comparableInteger).value == 22) {
		t.Error("Put() fails when pushing the fourth node (22)")
	}

	bst.Put(comparableInteger{23})

	if !(bst.Size() == 5 && bst.root.left.left.right.data.(comparableInteger).value == 23) {
		t.Error("Put() fails when pushing the fifth node (23)")
	}

	bst.Put(comparableInteger{75})

	if !(bst.Size() == 6 && bst.root.right.data.(comparableInteger).value == 75) {
		t.Error("Put() fails when pushing the sixth node (23)")
	}

	bst.Put(comparableInteger{60})

	if !(bst.Size() == 7 && bst.root.right.left.data.(comparableInteger).value == 60) {
		t.Error("Put() fails when pushing the seventh node (60)")
	}

	bst.Put(comparableInteger{61})

	if !(bst.Size() == 8 && bst.root.right.left.right.data.(comparableInteger).value == 61) {
		t.Error("Put() fails when pushing the eighth node (61)")
	}
}

func TestRemove(t *testing.T) {
	var emptyBst BinarySearchTree
	bst := &emptyBst

	bst.Put(comparableInteger{44})
	bst.Put(comparableInteger{17})
	bst.Put(comparableInteger{8})
	bst.Put(comparableInteger{32})
	bst.Put(comparableInteger{28})
	bst.Put(comparableInteger{21})
	bst.Put(comparableInteger{29})
	bst.Put(comparableInteger{88})
	bst.Put(comparableInteger{65})
	bst.Put(comparableInteger{54})
	bst.Put(comparableInteger{82})
	bst.Put(comparableInteger{76})
	bst.Put(comparableInteger{68})
	bst.Put(comparableInteger{80})
	bst.Put(comparableInteger{97})
	bst.Put(comparableInteger{93})

	// Remove a node with only a left child
	bst.Remove(comparableInteger{32})

	// Check the whole left subtree to verify that it's correct, along
	// with the size of the tree after the remove
	if !(bst.root.data.(comparableInteger).value == 44 &&
		bst.root.left.data.(comparableInteger).value == 17 &&
		bst.root.left.left.data.(comparableInteger).value == 8 &&
		bst.root.left.left.left == nil &&
		bst.root.left.left.right == nil &&
		bst.root.left.right.data.(comparableInteger).value == 28 &&
		bst.root.left.right.left.data.(comparableInteger).value == 21 &&
		bst.root.left.right.left.left == nil &&
		bst.root.left.right.left.right == nil &&
		bst.root.left.right.right.data.(comparableInteger).value == 29 &&
		bst.root.left.right.right.left == nil &&
		bst.root.left.right.right.right == nil &&
		bst.Size() == 15) {
		t.Error("Remove() fails when removing a node with only one child (32)")
	}

	// Remove a node with two children
	bst.Remove(comparableInteger{88})

	// Like above, check the whole right subtree and the size of the tree
	if !(bst.root.data.(comparableInteger).value == 44 &&
		bst.root.right.data.(comparableInteger).value == 82 &&
		bst.root.right.left.data.(comparableInteger).value == 65 &&
		bst.root.right.left.left.data.(comparableInteger).value == 54 &&
		bst.root.right.left.left.left == nil &&
		bst.root.right.left.left.right == nil &&
		bst.root.right.left.right.data.(comparableInteger).value == 76 &&
		bst.root.right.left.right.left.data.(comparableInteger).value == 68 &&
		bst.root.right.left.right.left.left == nil &&
		bst.root.right.left.right.left.right == nil &&
		bst.root.right.left.right.right.data.(comparableInteger).value == 80 &&
		bst.root.right.left.right.right.left == nil &&
		bst.root.right.left.right.right.right == nil &&
		bst.root.right.right.data.(comparableInteger).value == 97 &&
		bst.root.right.right.right == nil &&
		bst.root.right.right.left.data.(comparableInteger).value == 93 &&
		bst.root.right.right.left.left == nil &&
		bst.root.right.right.left.right == nil &&
		bst.Size() == 14) {
		t.Error("Remove() fails when removing a node with two children (88)")
	}

	bst.Put(comparableInteger{100})
	bst.Put(comparableInteger{107})
	bst.Put(comparableInteger{99})
	bst.Put(comparableInteger{100})
	bst.Put(comparableInteger{100})

	// Remove a node with a duplicate to verify that only the first occurrence
	// of the data is removed
	bst.Remove(comparableInteger{100})

	if !(bst.root.data.(comparableInteger).value == 44 &&
		bst.root.right.data.(comparableInteger).value == 82 &&
		bst.root.right.left.data.(comparableInteger).value == 65 &&
		bst.root.right.left.left.data.(comparableInteger).value == 54 &&
		bst.root.right.left.left.left == nil &&
		bst.root.right.left.left.right == nil &&
		bst.root.right.left.right.data.(comparableInteger).value == 76 &&
		bst.root.right.left.right.left.data.(comparableInteger).value == 68 &&
		bst.root.right.left.right.left.left == nil &&
		bst.root.right.left.right.left.right == nil &&
		bst.root.right.left.right.right.data.(comparableInteger).value == 80 &&
		bst.root.right.left.right.right.left == nil &&
		bst.root.right.left.right.right.right == nil &&
		bst.root.right.right.data.(comparableInteger).value == 97 &&
		bst.root.right.right.right.data.(comparableInteger).value == 100 &&
		bst.root.right.right.right.left.data.(comparableInteger).value == 99 &&
		bst.root.right.right.right.left.left == nil &&
		bst.root.right.right.right.left.right.data.(comparableInteger).value == 100 &&
		bst.root.right.right.right.left.right.left == nil &&
		bst.root.right.right.right.left.right.right == nil &&
		bst.root.right.right.right.right.data.(comparableInteger).value == 107 &&
		bst.root.right.right.right.right.left == nil &&
		bst.root.right.right.right.right.right == nil &&
		bst.Size() == 18) {
		t.Error("Remove() fails when removing a node with a duplicate (100)")
	}
}

func TestIsEmpty(t *testing.T) {
	var emptyBst BinarySearchTree
	bst := &emptyBst

	if !bst.IsEmpty() {
		t.Error("IsEmpty() fails on a newly initialized, empty BST")
	}

	bst.Put(comparableInteger{1})

	if bst.IsEmpty() {
		t.Error("IsEmpty() fails after adding one node to an empty BST")
	}
}
