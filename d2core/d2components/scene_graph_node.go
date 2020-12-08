package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2scene"
)

// static check that SceneGraphNode implements Component
var _ akara.Component = &SceneGraphNode{}

// SceneGraphNode represents an entities x,y axis scale as a vector
type SceneGraphNode struct {
	*d2scene.Node
}

// New creates a new SceneGraphNode instance. By default, the scale is (1,1)
func (*SceneGraphNode) New() akara.Component {
	return &SceneGraphNode{
		Node: d2scene.NewNode(),
	}
}

// SceneGraphNodeFactory is a wrapper for the generic component factory that returns SceneGraphNode component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a SceneGraphNode.
type SceneGraphNodeFactory struct {
	*akara.ComponentFactory
}

// Add adds a SceneGraphNode component to the given entity and returns it
func (m *SceneGraphNodeFactory) Add(id akara.EID) *SceneGraphNode {
	return m.ComponentFactory.Add(id).(*SceneGraphNode)
}

// Get returns the SceneGraphNode component for the given entity, and a bool for whether or not it exists
func (m *SceneGraphNodeFactory) Get(id akara.EID) (*SceneGraphNode, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*SceneGraphNode), found
}
