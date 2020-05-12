package main

import log "github.com/sirupsen/logrus"

func (t *BPlussTree) Insert(key int, value []byte) error {
	var leaf *Node
	log.Infof("Inserting in %v", key)
	t.len +=  1
	if _, err := t.Search(key); err == nil {
		return keyExist
	}

	pointer := makeRecord(value)

	if t.Root == nil {
		return t.CreateNewBPTree(key, pointer)
	}

	leaf, err := t.findLeaf(key)
	if err != nil {
		return err
	}

	if leaf.NumKeys < order-1 {
		insertIntoLeaf(leaf, key, pointer)
		return nil
	}


	return t.insertIntoLeafAfterSplitting(leaf, key, pointer)
}


func (t *BPlussTree) insertIntoLeafAfterSplitting(leaf *Node, key int, pointer *Record) error {
	var (
		indexInsert int
		split       int
		newKey      int
	)

	newLeaf := makeLeaf()

	tempKey := make([]int, order)

	tempPointers := make([]interface{}, order)

	for indexInsert < order-1 && leaf.Keys[indexInsert] < key {
		indexInsert += 1
	}
	j := 0
	for i := 0; i < leaf.NumKeys; i++ {
		if j == indexInsert {
			j += 1
		}
		tempKey[j] = leaf.Keys[i]
		tempPointers[j] = leaf.Pointers[i]
		j += 1
	}

	tempKey[indexInsert] = key
	tempPointers[indexInsert] = pointer

	leaf.NumKeys = 0

	split = mid(order - 1)

	for i := 0; i < split; i++ {
		leaf.Pointers[i] = tempPointers[i]
		leaf.Keys[i] = tempKey[i]
		leaf.NumKeys += 1
	}

	j = 0
	for i := split; i < order; i++ {
		newLeaf.Pointers[j] = tempPointers[i]
		newLeaf.Keys[j] = tempKey[i]
		newLeaf.NumKeys += 1
		j += 1
	}

	newLeaf.Pointers[order-1] = leaf.Pointers[order-1]
	leaf.Pointers[order-1] = newLeaf

	for i := leaf.NumKeys; i < order-1; i++ {
		leaf.Pointers[i] = nil
	}
	for i := newLeaf.NumKeys; i < order-1; i++ {
		newLeaf.Pointers[i] = nil
	}

	newLeaf.Parent = leaf.Parent
	newKey = newLeaf.Keys[0]

	return t.insertIntoParent(leaf, newKey, newLeaf)
}

func insertIntoLeaf(leaf *Node, key int, pointer *Record) {
	var pointInsert int

	for pointInsert < leaf.NumKeys && leaf.Keys[pointInsert] < key {
		pointInsert += 1
	}

	for i := leaf.NumKeys; i > pointInsert; i-- {
		leaf.Keys[i] = leaf.Keys[i-1]
		leaf.Pointers[i] = leaf.Pointers[i-1]
	}
	leaf.Keys[pointInsert] = key
	leaf.Pointers[pointInsert] = pointer
	leaf.NumKeys += 1
	return
}

func (t *BPlussTree) insertIntoParent(left *Node, key int, right *Node) error {
	var indexLeaf int
	parent := left.Parent

	if parent == nil {
		return t.insertIntoNewRoot(left, key, right)
	}
	indexLeaf = getLeftIndex(parent, left)

	if parent.NumKeys < order-1 {
		insertIntoNode(parent, indexLeaf, key, right)
		return nil
	}

	return t.insertIntoNodeAfterSplitting(parent, indexLeaf, key, right)
}

func (t *BPlussTree) insertIntoNodeAfterSplitting(oldNode *Node, indexLeaf, key int, right *Node) error {
	var (
		i            int
		j            int
		split        int
		prime        int
		child        *Node
		tempKey      []int
		tempPointers []interface{}
	)

	tempPointers = make([]interface{}, order+1)

	tempKey = make([]int, order)

	for i = 0; i < oldNode.NumKeys+1; i++ {
		if j == indexLeaf+1 {
			j += 1
		}
		tempPointers[j] = oldNode.Pointers[i]
		j += 1
	}

	j = 0
	for i = 0; i < oldNode.NumKeys; i++ {
		if j == indexLeaf {
			j += 1
		}
		tempKey[j] = oldNode.Keys[i]
		j += 1
	}

	tempPointers[indexLeaf+1] = right
	tempKey[indexLeaf] = key

	split = mid(order)
	newNode := makeNode()

	oldNode.NumKeys = 0
	for i = 0; i < split-1; i++ {
		oldNode.Pointers[i] = tempPointers[i]
		oldNode.Keys[i] = tempKey[i]
		oldNode.NumKeys += 1
	}
	oldNode.Pointers[i] = tempPointers[i]
	prime = tempKey[split-1]
	j = 0
	for i += 1; i < order; i++ {
		newNode.Pointers[j] = tempPointers[i]
		newNode.Keys[j] = tempKey[i]
		newNode.NumKeys += 1
		j += 1
	}
	newNode.Pointers[j] = tempPointers[i]
	newNode.Parent = oldNode.Parent
	for i = 0; i <= newNode.NumKeys; i++ {
		child, _ = newNode.Pointers[i].(*Node)
		child.Parent = newNode
	}

	return t.insertIntoParent(oldNode, prime, newNode)
}

func insertIntoNode(n *Node, indexLeaf, key int, right *Node) {
	for i := n.NumKeys; i > indexLeaf; i-- {
		n.Pointers[i+1] = n.Pointers[i]
		n.Keys[i] = n.Keys[i-1]
	}
	n.Pointers[indexLeaf+1] = right
	n.Keys[indexLeaf] = key
	n.NumKeys += 1
}


func (t *BPlussTree) insertIntoNewRoot(left *Node, key int, right *Node) error {
	t.Root = makeNode()
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = left
	t.Root.Pointers[1] = right
	t.Root.NumKeys += 1
	t.Root.Parent = nil
	left.Parent = t.Root
	right.Parent = t.Root
	return nil
}
