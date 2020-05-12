package main

type BPlussTree struct {
	Root *Node
	len	int
}

type Record struct {
	Value []byte
}

type Node struct {
	Pointers []interface{}
	Keys     []int
	Parent   *Node
	IsLeaf   bool
	NumKeys  int
	Next     *Node
}



func NewPlusTree() *BPlussTree {
	return &BPlussTree{
		len: 0,
	}
}



func (t *BPlussTree) CreateNewBPTree (key int, pointer *Record) error {
	t.Root  = makeLeaf()
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = pointer
	t.Root.Pointers[order-1] = nil
	t.Root.Parent = nil
	t.Root.NumKeys += 1
	return nil
}


func makeLeaf() *Node {
	leaf := makeNode()
	leaf.IsLeaf = true
	return leaf
}

func makeRecord(value []byte) *Record {
	return &Record{
		Value: value,
	}
}

func getLeftIndex(parent, left *Node) int {
	indexLeaf := 0
	for indexLeaf <= parent.NumKeys && parent.Pointers[indexLeaf] != left {
		indexLeaf += 1
	}
	return indexLeaf
}


func makeNode() *Node {
	newNode := new(Node)

	newNode.Keys = make([]int, order-1)

	newNode.Pointers = make([]interface{}, order)

	newNode.IsLeaf = false
	newNode.NumKeys = 0
	newNode.Parent = nil
	newNode.Next = nil
	return newNode
}
