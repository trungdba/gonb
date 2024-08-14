package coll

// define linked list
type Linkedlist struct {
	len  int
	root *Node
	curr *Node
	last *Node
}

// constructor creates a new linked list
func NewLinkedList() *Linkedlist {
	return &Linkedlist{}
}

// method to get the first node of the linked list
func (l *Linkedlist) First() *Node {
	if l.root == nil {
		return nil
	}
	l.curr = l.root
	return l.root
}

// method to get the next node of the linked list
func (l *Linkedlist) Next() *Node {
	if l.curr == nil {
		return nil
	}
	l.curr = l.curr.next
	return l.curr
}

// method to get the last node of the linked list
func (l *Linkedlist) Last() *Node {
	return l.last
}

// method to add a new node at the beginning of the linked list
func (l *Linkedlist) Prepend(n *Node) {
	if l.root == nil {
		l.root = n
		l.last = l.root
	} else {
		n.next = l.root
		l.root = n
	}
	l.len++
}

// method to add a new node at the tail of the linked list
func (l *Linkedlist) Append(n *Node) {
	if l.root == nil {
		l.root = n
	} else {
		curr := l.root
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
		l.last = curr.next
	}
	l.len++
}

// method to remove a node from the linked list
func (l *Linkedlist) Remove(key interface{}) {
	if l.root == nil {
		return
	}
	if l.root.key == key {
		l.root = l.root.next
	} else {
		curr := l.root
		for curr.next != nil && curr.next.key != key {
			curr = curr.next
		}
		if curr.next != nil {
			curr.next = curr.next.next
		}
	}
	l.len--
}

// method to remove all nodes from the linked list
func (l *Linkedlist) RemoveAll() {
	l.len = 0
	l.root = nil
	l.curr = nil
	l.last = nil
}

// method to replace an existing node in the linked list
func (l *Linkedlist) Replace(n *Node) {
	if l.root == nil {
		return
	}
	if l.root.key == n.key {
		n.next = l.root.next
		l.root = n
	} else {
		curr := l.root
		for curr.next != nil && curr.next.key != n.key {
			curr = curr.next
		}
		if curr.next != nil {
			n.next = curr.next.next
			curr.next = n
		}
	}
}

// method to get the node value from the linked list
func (l *Linkedlist) Get(key interface{}) interface{} {
	curr := l.root
	for curr != nil {
		if curr.key == key {
			return curr.data
		} else {
			curr = curr.next
		}
	}
	return nil
}

// method to get the node from the linked list
func (l *Linkedlist) GetNode(key interface{}) *struct {
	k interface{}
	v any
} {
	curr := l.root
	for curr != nil {
		if curr.key == key {
			return &struct {
				k interface{}
				v any
			}{curr.key, curr.data}
		} else {
			curr = curr.next
		}
	}
	return nil
}

// method to get all the nodes from the linked list
func (l *Linkedlist) GetAllNode() []struct {
	k interface{}
	v any
} {
	nodes := []struct {
		k interface{}
		v any
	}{}
	curr := l.root
	for curr != nil {
		nodes = append(nodes, struct {
			k interface{}
			v any
		}{curr.key, curr.data})
		curr = curr.next
	}
	return nodes
}

// method to check if a key belongs to the linked list
func (l *Linkedlist) ContainsKey(key interface{}) bool {
	curr := l.root
	for curr != nil {
		if curr.key == key {
			return true
		} else {
			curr = curr.next
		}
	}
	return false
}

// method to check if a value belongs to the linked list
func (l *Linkedlist) ContainsValue(value any) bool {
	curr := l.root
	for curr != nil {
		if curr.data == value {
			return true
		} else {
			curr = curr.next
		}
	}
	return false
}

// method to clear the linked list
func (l *Linkedlist) Clear() {
	l.len = 0
	l.root = nil
	l.curr = nil
	l.last = nil
	l = nil
}

// method to count the total nodes of the linked list
func (l *Linkedlist) Length() int {
	return l.len
}

// method to convert the linked list to a map
func (l *Linkedlist) Map() map[interface{}]any {
	newMap := make(map[interface{}]any)
	curr := l.root
	for curr != nil {
		newMap[curr.key] = curr.data
		curr = curr.next
	}
	return newMap
}

// method to convert the linked list to a array of keys
func (l *Linkedlist) Keys() []interface{} {
	var keys []interface{}
	curr := l.root
	for curr != nil {
		keys = append(keys, curr.key)
		curr = curr.next
	}
	return keys
}

// method to convert the linked list to a array of values
func (l *Linkedlist) Values() []any {
	var values []any
	curr := l.root
	for curr != nil {
		values = append(values, curr.data)
		curr = curr.next
	}
	return values
}
