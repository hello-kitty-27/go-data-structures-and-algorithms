package trees

// Comparable is an interface that data to be stored in a tree can implement
// so that the nodes of the tree can be compared to each other (for example,
// in a binary search tree to determine which way to search)
type Comparable interface {
	Compare(Comparable) int
}

// comparableInteger is a concrete implementation of Comparable
// to use for testing the data structures within this package
type comparableInteger struct {
	value int
}

func (ci comparableInteger) Compare(other Comparable) int {
	otherCi := other.(comparableInteger)
	return ci.value - otherCi.value
}
