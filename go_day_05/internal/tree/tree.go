package tree

type Node struct {
	HasToy bool
	Left   *Node
	Right  *Node
}

func New(value bool) *Node {
	return &Node{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}
