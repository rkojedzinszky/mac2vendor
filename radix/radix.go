package radix

// Readonly interface for read only access
type Readonly interface {
	Get(string) interface{}
}

// Node element of a tree
type Node struct {
	nodes map[rune]*Node
	value interface{}
}

// New creates a new Node
func New() *Node {
	return &Node{
		nodes: make(map[rune]*Node),
	}
}

// Add a string with a value
func (n *Node) Add(s string, value interface{}) {
	for _, chr := range s {
		if _, ok := n.nodes[chr]; !ok {
			n.nodes[chr] = New()
		}

		n = n.nodes[chr]
	}

	n.value = value
}

// Get returns the longest match
func (n *Node) Get(s string) interface{} {
	last := n

	for _, chr := range s {
		var ok bool
		if n, ok = n.nodes[chr]; !ok {
			break
		}

		if n.value != nil {
			last = n
		}
	}

	return last.value
}

// GetReadonlyRadix returns a readonly radix tree
func (n *Node) Readonly() Readonly {
	return n
}
