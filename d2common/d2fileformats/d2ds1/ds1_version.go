package d2ds1

type ds1version int

const (
	v3  ds1version = 3
	v4  ds1version = 4
	v7  ds1version = 7
	v8  ds1version = 8
	v9  ds1version = 9
	v10 ds1version = 10
	v12 ds1version = 12
	v13 ds1version = 13
	v14 ds1version = 14
	v15 ds1version = 15
	v16 ds1version = 16
	v18 ds1version = 18
)

func (v ds1version) hasUnknown1Bytes() bool {
	// just after the header will be some meaningless (?) bytes
	return v >= v9 && v <= v13
}

func (v ds1version) hasUnknown2Bytes() bool {
	return v >= v18
}

func (v ds1version) specifiesAct() bool {
	// in the header
	return v >= v8
}

func (v ds1version) specifiesSubstitutionType() bool {
	// in the header
	return v >= v10
}

func (v ds1version) specifiesWalls() bool {
	// just after header, specifies number of Walls
	return v >= v4
}

func (v ds1version) specifiesFloors() bool {
	// just after header, specifies number of Floors
	return v >= v16
}

func (v ds1version) hasFileList() bool {
	return v >= v3
}

func (v ds1version) hasObjects() bool {
	return v >= v3
}

func (v ds1version) hasSubstitutions() bool {
	return v >= v12
}
