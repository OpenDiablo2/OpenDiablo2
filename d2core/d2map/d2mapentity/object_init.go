package d2mapentity

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Finds an init function for the given object
func initObject(ob *Object) (bool, error) {
	funcs := map[int]func(*Object) error{
		8:  initTorch,
		14: initTorch,
		17: initWaypoint,
		34: initTorchRnd,
	}

	fun, ok := funcs[ob.objectRecord.InitFn]
	if !ok {
		return false, nil
	}

	if err := fun(ob); err != nil {
		return false, err
	}

	return true, nil
}

// Initializes torch/brazier type objects
func initTorch(ob *Object) error {
	if ob.objectRecord.HasAnimationMode[d2enum.ObjectAnimationModeOpened] {
		return ob.setMode(d2enum.ObjectAnimationModeOpened, 0, true)
	}

	return nil
}

func initWaypoint(ob *Object) error {
	// Turn these on unconditionally for now, they look nice :)
	if ob.objectRecord.HasAnimationMode[d2enum.ObjectAnimationModeOpened] {
		return ob.setMode(d2enum.ObjectAnimationModeOpened, 0, true)
	}

	return nil
}

// Randomly spawns in either NU or OP
func initTorchRnd(ob *Object) error {
	const coinToss = 2
	n := rand.Intn(coinToss) // nolint:gosec // not concerned with crypto-strong randomness

	if n > 0 {
		return ob.setMode(d2enum.ObjectAnimationModeNeutral, 0, true)
	}

	return ob.setMode(d2enum.ObjectAnimationModeOperating, 0, true)
}
