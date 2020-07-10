package d2enum

//go:generate stringer -linecomment -type ObjectAnimationMode -output object_animation_mode_string.go
//go:generate string2enum -samepkg -linecomment -type ObjectAnimationMode -output object_animation_mode_string2enum.go

// ObjectAnimationMode represents object animation modes
type ObjectAnimationMode int

// Object animation modes
const (
	ObjectAnimationModeNeutral   ObjectAnimationMode = iota // NU
	ObjectAnimationModeOperating                            // OP
	ObjectAnimationModeOpened                               // ON
	ObjectAnimationModeSpecial1                             // S1
	ObjectAnimationModeSpecial2                             // S2
	ObjectAnimationModeSpecial3                             // S3
	ObjectAnimationModeSpecial4                             // S4
	ObjectAnimationModeSpecial5                             // S5
)
