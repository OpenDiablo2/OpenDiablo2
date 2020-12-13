package d2input

import "github.com/gravestench/akara"

// NewInputVector creates a new input vector
func NewInputVector() *InputVector {
	v := &InputVector{
		KeyVector:         akara.NewBitSet(),
		ModifierVector:    akara.NewBitSet(),
		MouseButtonVector: akara.NewBitSet(),
	}

	return v.Clear()
}

// InputVector represents the state of keys, modifiers, and mouse buttons.
// It can be used to compare input states, and is intended to be used as such:
// 		* whatever manages system input keeps a "current" input vector and updates it
//  	* things that are listening for certain inputs will be compared using `Contains` and `Intersects` methods
type InputVector struct {
	KeyVector         *akara.BitSet
	ModifierVector    *akara.BitSet
	MouseButtonVector *akara.BitSet
}

// SetKey sets the corresponding key bit in the keys bitset
func (iv *InputVector) SetKey(key Key) *InputVector {
	return iv.SetKeys([]Key{key})
}

// SetKeys sets multiple key bits in the keys bitset
func (iv *InputVector) SetKeys(keys []Key) *InputVector {
	if len(keys) == 0 {
		return iv
	}

	for _, key := range keys {
		iv.KeyVector.Set(key, true)
	}

	return iv
}

// SetModifier sets the corresponding modifier bit in the modifier bitset
func (iv *InputVector) SetModifier(mod Modifier) *InputVector {
	return iv.SetModifiers([]Modifier{mod})
}

// SetModifiers sets multiple modifier bits in the modifier bitset
func (iv *InputVector) SetModifiers(mods []Modifier) *InputVector {
	if len(mods) == 0 {
		return iv
	}

	for _, key := range mods {
		iv.ModifierVector.Set(key, true)
	}

	return iv
}

// SetMouseButton sets the corresponding mouse button bit in the mouse button bitset
func (iv *InputVector) SetMouseButton(button MouseButton) *InputVector {
	return iv.SetMouseButtons([]MouseButton{button})
}

// SetMouseButtons sets multiple mouse button bits in the mouse button bitset
func (iv *InputVector) SetMouseButtons(buttons []MouseButton) *InputVector {
	if len(buttons) == 0 {
		return iv
	}

	for _, key := range buttons {
		iv.MouseButtonVector.Set(key, true)
	}

	return iv
}

// Contains returns true if this input vector is a superset of the given input vector
func (iv *InputVector) Contains(other *InputVector) bool {
	keys := iv.KeyVector.ContainsAll(other.KeyVector)
	buttons := iv.MouseButtonVector.ContainsAll(other.MouseButtonVector)

	// We do Equals here, because we dont want CTRL+X and CTRL+ALT+X to fire at the same time
	mods := iv.ModifierVector.Equals(other.ModifierVector)

	return keys && mods && buttons
}

// Intersects returns true if this input vector shares any bits with the given input vector
func (iv *InputVector) Intersects(other *InputVector) bool {
	keys := iv.KeyVector.Intersects(other.KeyVector)
	mods := iv.ModifierVector.Intersects(other.ModifierVector)
	buttons := iv.MouseButtonVector.Intersects(other.MouseButtonVector)

	return keys || mods || buttons
}

// Clear sets all bits in this input vector to 0
func (iv *InputVector) Clear() *InputVector {
	iv.KeyVector.Clear()
	iv.ModifierVector.Clear()
	iv.MouseButtonVector.Clear()

	return iv
}
