package d2animdata

// AnimationDataRecord represents a single record from the AnimData.d2 file
type AnimationDataRecord struct {
	name               string
	framesPerDirection uint32
	speed              uint16
	events             map[int]AnimationEvent
}

// FramesPerDirection returns frames per direction value
func (r *AnimationDataRecord) FramesPerDirection() int {
	return int(r.framesPerDirection)
}

// SetFramesPerDirection sets frames per direction value
func (r *AnimationDataRecord) SetFramesPerDirection(fpd uint32) {
	r.framesPerDirection = fpd
}

// Speed returns animation's speed
func (r *AnimationDataRecord) Speed() int {
	return int(r.speed)
}

// SetSpeed sets record's speed
func (r *AnimationDataRecord) SetSpeed(s uint16) {
	r.speed = s
}

// FPS returns the frames per second for this animation record
func (r *AnimationDataRecord) FPS() float64 {
	speedf := float64(r.speed)
	divisorf := float64(speedDivisor)
	basef := float64(speedBaseFPS)

	return basef * speedf / divisorf
}

// FrameDurationMS returns the duration in milliseconds that a frame is displayed
func (r *AnimationDataRecord) FrameDurationMS() float64 {
	return milliseconds / r.FPS()
}

// Events returns events map
func (r *AnimationDataRecord) Events() map[int]AnimationEvent {
	return r.events
}

// Event returns specific event
func (r *AnimationDataRecord) Event(idx int) AnimationEvent {
	event, found := r.events[idx]
	if found {
		return event
	}

	return AnimationEventNone
}

// SetEvent sets event on specific index to given
func (r *AnimationDataRecord) SetEvent(index int, event AnimationEvent) {
	r.events[index] = event
}
