package binarysearchtree

type BinarySearchTreeNode struct {
	Key   int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func (b *BinarySearchTreeNode) Add(key int) {
	if b.Key == key {
		return
	}

	newNode := &BinarySearchTreeNode{Key: key}

	if b.Key < key {
		b.addToRight(newNode)

		return
	}

	if b.Key > key {
		b.addToLeft(newNode)

		return
	}
}

func (b *BinarySearchTreeNode) Search(key int) bool {
	if b == nil {
		return false
	}

	if b.Key < key {
		return b.Right.Search(key)
	}

	if b.Key > key {
		return b.Left.Search(key)
	}

	return true
}

func (b *BinarySearchTreeNode) addToRight(newNode *BinarySearchTreeNode) {
	if b.Right == nil {
		b.Right = newNode

		return
	}

	b.Right.Add(newNode.Key)
}

func (b *BinarySearchTreeNode) addToLeft(newNode *BinarySearchTreeNode) {
	if b.Left == nil {
		b.Left = newNode

		return
	}

	b.Left.Add(newNode.Key)
}
