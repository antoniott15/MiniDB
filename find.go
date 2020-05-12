package main

func (t *BPlussTree) Search(key int) (*Record, error) {
	i := 0
	c,err := t.findLeaf(key)
	if err != nil {
		return nil, keyNotFound
	}
	for i = 0; i < c.NumKeys; i++ {
		if c.Keys[i] == key {
			break
		}
	}
	if i == c.NumKeys {
		return nil, keyNotFound
	}

	r, _ := c.Pointers[i].(*Record)

	return r, nil
}




func (t *BPlussTree) findLeaf(key int) (*Node, error) {
	i := 0
	if t.Root == nil {
		return nil,treeIsEmpty
	}
	for !t.Root.IsLeaf {
		i = 0
		for i < t.Root.NumKeys {
			if key >= t.Root.Keys[i] {
				i += 1
			} else {
				break
			}
		}
		t.Root, _ = t.Root.Pointers[i].(*Node)
	}
	return t.Root, nil
}
