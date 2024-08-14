package coll

// define node
type Node struct {
	data any
	key  interface{}
	next *Node
}

// constructor creates a new node
func NewNode(k interface{}, v any) *Node {
	return &Node{key: k, data: v}
}

// return the node
func (n *Node) Get() *struct {
	k interface{}
	v any
} {
	if n == nil {
		return nil
	}
	return &struct {
		k interface{}
		v any
	}{n.key, n.data}
}

// return the key of the node
func (n *Node) GetKey() interface{} {
	if n == nil {
		return 0
	}
	return n.key
}

// return the value of the node
func (n *Node) GetValue() any {
	if n == nil {
		return nil
	}
	return n.data
}
