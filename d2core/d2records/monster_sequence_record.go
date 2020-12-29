package d2records

// MonsterSequences contains the MonsterSequenceRecords
type MonsterSequences map[string]*MonsterSequenceRecord

// MonsterSequenceRecord contains a record for a monster sequence
// Composed of multiple lines from monseq.txt with the same name in the first column.
// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=395]
type MonsterSequenceRecord struct {

	// Name of the sequence, referred to by monstats.txt
	Name string

	// Frames of this sequence
	Frames []*MonsterSequenceFrame
}

// MonsterSequenceFrame represents a single frame of a monster sequence
type MonsterSequenceFrame struct {
	// The animation mode for this frame (refers to MonMode.txt)
	Mode string

	// The frame of the animation mode used for this frame of the sequence
	Frame int

	// Direction of the frame, enumerated by d2enum.AnimationFrameDirection
	Direction int

	// Event triggered by this frame
	Event int
}
