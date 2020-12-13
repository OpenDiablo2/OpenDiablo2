package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// NewNode creates and initializes a new scene graph node
func NewNode() *Node {
	n := &Node{
		World:    d2math.NewMatrix4(nil),
		Local:    d2math.NewMatrix4(nil),
		children: make([]*Node, 0),
	}

	return n
}

// Node is a scene graph node.
type Node struct {
	parent   *Node
	World    *d2math.Matrix4
	Local    *d2math.Matrix4
	children []*Node
}

// SetParent sets the parent of this scene graph node
func (n *Node) SetParent(p *Node) *Node {
	n.parent.removeChild(n)

	n.parent = p

	if p != nil {
		n.parent.children = append(n.parent.children, n)
	}

	return n
}

func (n *Node) removeChild(m *Node) *Node {
	if n == nil {
		return nil
	}

	if m == nil {
		return n
	}

	for idx := len(n.children)-1; idx >= 0; idx-- {
		if n.children[idx] != m {
			continue
		}

		n.children = append(n.children[:idx], n.children[idx+1:]...)
	}

	return n
}

// UpdateWorldMatrix updates this node's World matrix using the (optional) parent World matrix
func (n *Node) UpdateWorldMatrix(args ...*d2math.Matrix4) *Node {
	// this is a hack so that we can just call `node.UpdateWorldMatrix()`
	parentWorldMatrix := (*d2math.Matrix4)(nil)
	if len(args) > 0 {
		parentWorldMatrix = args[0]
	}

	n.World = parentWorldMatrix

	for idx := range n.children {
		n.children[idx].UpdateWorldMatrix(n.GetWorldMatrix())
	}

	return n
}

// GetWorldMatrix applies the local transform to the world matrix and returns it
func (n *Node) GetWorldMatrix() *d2math.Matrix4 {
	if n.World == nil {
		return n.Local.Clone()
	}

	return n.World.Clone().Multiply(n.Local)
}
