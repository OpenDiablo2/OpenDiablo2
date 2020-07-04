package d2object

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Finds an init function for the given object
func initObject(ob *Object) bool {
	funcs := map[int]func(*Object){
		8:  initTorch,
		14: initTorch,
		17: initWaypoint,
	}

	fun, ok := funcs[ob.objectRecord.InitFn]
	if !ok {
		return false
	}

	fun(ob)

	return true
}

// Initializes torch/brazier type objects
func initTorch(ob *Object) {
	if ob.objectRecord.HasAnimationMode[d2enum.AnimationModeObjectOperating] {
		ob.setMode("ON", 0)
	}
}

func initWaypoint(ob *Object) {
	// Turn these on unconditionally for now, they look nice :)
	if ob.objectRecord.HasAnimationMode[d2enum.AnimationModeObjectOperating] {
		ob.setMode("ON", 0)
	}
}
