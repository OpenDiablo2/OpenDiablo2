package d2input

import "github.com/gravestench/akara"

func NewInputVector() *InputVector {
	v := &InputVector{
		KeyVector:         akara.NewBitSet(),
		ModifierVector:    akara.NewBitSet(),
		MouseButtonVector: akara.NewBitSet(),
	}

	return v.Clear()
}

type InputVector struct {
	KeyVector         *akara.BitSet
	ModifierVector    *akara.BitSet
	MouseButtonVector *akara.BitSet
}

func (iv *InputVector) SetKey(key Key) *InputVector {
	return iv.SetKeys([]Key{key})
}

func (iv *InputVector) SetKeys(keys []Key) *InputVector {
	if len(keys) == 0 {
		return iv
	}

	for _, key := range keys {
		iv.KeyVector.Set(int(key), true)
	}

	return iv
}

func (iv *InputVector) SetModifier(mod Modifier) *InputVector {
	return iv.SetModifiers([]Modifier{mod})
}

func (iv *InputVector) SetModifiers(mods []Modifier) *InputVector {
	if len(mods) == 0 {
		return iv
	}

	for _, key := range mods {
		iv.ModifierVector.Set(int(key), true)
	}

	return iv
}

func (iv *InputVector) SetMouseButton(button MouseButton) *InputVector {
	return iv.SetMouseButtons([]MouseButton{button})
}

func (iv *InputVector) SetMouseButtons(buttons []MouseButton) *InputVector {
	if len(buttons) == 0 {
		return iv
	}

	for _, key := range buttons {
		iv.MouseButtonVector.Set(int(key), true)
	}

	return iv
}

func (iv *InputVector) Contains(other *InputVector) bool {
	keys := iv.KeyVector.ContainsAll(other.KeyVector)
	buttons := iv.MouseButtonVector.ContainsAll(other.MouseButtonVector)

	// We do Equals here, because we dont want CTRL+X and CTRL+ALT+X to fire at the same time
	mods := iv.ModifierVector.Equals(other.ModifierVector)

	return keys && mods && buttons
}

func (iv *InputVector) Intersects(other *InputVector) bool {
	keys := iv.KeyVector.Intersects(other.KeyVector)
	mods := iv.ModifierVector.Intersects(other.ModifierVector)
	buttons := iv.MouseButtonVector.Intersects(other.MouseButtonVector)

	return keys || mods || buttons
}

func (iv *InputVector) Clear() *InputVector {
	iv.KeyVector.Clear()
	iv.ModifierVector.Clear()
	iv.MouseButtonVector.Clear()

	return iv
}
