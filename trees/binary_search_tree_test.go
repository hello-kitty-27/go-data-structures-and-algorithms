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
