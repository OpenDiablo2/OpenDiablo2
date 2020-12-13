package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode()

	if n == nil {
		t.Error("node not created")
	}
}

func TestNode_SetParent(t *testing.T) {
	a := NewNode()
	b := NewNode()

	a.SetParent(b)

	if a.parent != b || len(b.children) != 1 {
		t.Error("Node parent not set")
	}

	c := NewNode()

	a.SetParent(c)

	if a.parent != c || len(b.children) != 0 {
		t.Error("Node child not removed")
	}

	b.SetParent(c)

	if b.parent != c || len(c.children) != 2 {
		t.Error("Node child not removed")
	}

	a.SetParent(nil)
	b.SetParent(nil)
	c.SetParent(nil)

	if a.parent != nil || b.parent != nil || c.parent != nil {
		t.Error("Node child not removed")
	}
}

func TestNode_UpdateWorldMatrix(t *testing.T) {
	world := NewNode()
	a := NewNode()

	a.SetParent(world)

	world.Local.SetXYZ(10, 20, 30)
	world.UpdateWorldMatrix()

	ax, ay, az := d2math.NewVector3(0, 0, 0).ApplyMatrix4(a.GetWorldMatrix()).XYZ()

	if ax != 20 && ay != 40 && az != 30 {
		t.Error("error updating world matrix")
	}
}
