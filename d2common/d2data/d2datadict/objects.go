package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// An ObjectRecord represents the settings for one type of object from objects.txt
type ObjectRecord struct {
	Name        string
	Description string
	Id          int
	Token       string // refers to what graphics this object uses

	SpawnMax        int     // unused?
	Selectable      [8]bool // is this mode selectable
	TrapProbability int     // unused

	SizeX int
	SizeY int

	NTgtFX int // unknown
	NTgtFY int // unknown
	NTgtBX int // unknown
	NTgtBY int // unknown

	FrameCount     [8]int  // how many frames does this mode have, 0 = skip
	FrameDelta     [8]int  // what rate is the animation played at (256 = 100% speed)
	CycleAnimation [8]bool // probably whether animation loops
	LightDiameter  [8]int
	BlocksLight    [8]bool
	HasCollision   [8]bool
	IsAttackable   bool // do we kick it when interacting
	StartFrame     [8]int

	EnvEffect       bool // unknown
	IsDoor          bool
	BlockVisibility bool // only works with IsDoor
	Orientation     int  // unknown (1=sw, 2=nw, 3=se, 4=ne)
	Trans           int  // controls palette mapping

	OrderFlag        [8]int  //  0 = object, 1 = floor, 2 = wall
	PreOperate       bool    // unknown
	HasAnimationMode [8]bool // 'Mode' in source, true if this mode is used

	XOffset int // in pixels offset
	YOffset int
	Draw    bool // if false, object isn't drawn (shadow is still drawn and player can still select though)

	LightRed   byte // if lightdiameter is set, rgb of the light
	LightGreen byte
	LightBlue  byte

	SelHD bool // whether these DCC components are selectable
	SelTR bool
	SelLG bool
	SelRA bool
	SelLA bool
	SelRH bool
	SelLH bool
	SelSH bool
	SelS  [8]bool

	TotalPieces int // selectable DCC components count
	SubClass    int // subclass of object:
	// 1 = shrine
	// 2 = obelisk
	// 4 = portal
	// 8 = container
	// 16 = arcane sanctuary gateway
	// 32 = well
	// 64 = waypoint
	// 128 = secret jails door

	XSpace int // unknown
	YSpace int

	NameOffset int // pixels to offset the name from the animation pivot

	MonsterOk      bool // unknown
	OperateRange   int  // distance object can be used from, might be unused
	ShrineFunction int  // unused
	Restore        bool // if true, object is stored in memory and will be retained if you leave and re-enter the area

	Parm           [8]int // unknown
	Act            int    // what acts this object can appear in (15 = all three)
	Lockable       bool
	Gore           bool // unknown, something with corpses
	Sync           bool // unknown
	Flicker        bool // light flickers if true
	Damage         int  // amount of damage done by this (used depending on operatefn)
	Beta           bool // if true, appeared in the beta?
	Overlay        bool // unknown
	CollisionSubst bool // unknown, controls some kind of special collision checking?

	Left   int // unknown, clickable bounding box?
	Top    int
	Width  int
	Height int

	OperateFn int // what function is called when the player clicks on the object
	// (todo: we should enumerate all the functions somewhere, but probably not here
	//        b/c it's a very long list)
	PopulateFn int // what function is used to spawn this object?
	// (see above todo)
	InitFn int // what function is run when the object is initialized?
	// (see above todo)
	ClientFn int // controls special audio-visual functions
	// (see above todo)

	RestoreVirgins bool // if true, only restores unused objects (see Restore)
	BlockMissile   bool // if true, missiles collide with this
	DrawUnder      bool // if true, drawn as a floor tile is
	OpenWarp       bool // needs clarification, controls whether highlighting shows
	// 'To ...' or 'trap door' when highlighting, not sure which is T/F
	AutoMap int // controls how this object appears on the map
	// 0 = it doesn't, rest of modes need to be analyzed
}

//nolint:funlen // Makes no sense to split
// CreateObjectRecord parses a row from objects.txt into an object record
func createObjectRecord(props []string) ObjectRecord {
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := ObjectRecord{
		Name:        props[inc()],
		Description: props[inc()],
		Id:          d2common.StringToInt(props[inc()]),
		Token:       props[inc()],

		SpawnMax: d2common.StringToInt(props[inc()]),
		Selectable: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},
		TrapProbability: d2common.StringToInt(props[inc()]),

		SizeX: d2common.StringToInt(props[inc()]),
		SizeY: d2common.StringToInt(props[inc()]),

		NTgtFX: d2common.StringToInt(props[inc()]),
		NTgtFY: d2common.StringToInt(props[inc()]),
		NTgtBX: d2common.StringToInt(props[inc()]),
		NTgtBY: d2common.StringToInt(props[inc()]),

		FrameCount: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},
		FrameDelta: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},
		CycleAnimation: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},
		LightDiameter: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},
		BlocksLight: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},
		HasCollision: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},
		IsAttackable: d2common.StringToUint8(props[inc()]) == 1,
		StartFrame: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},

		EnvEffect:       d2common.StringToUint8(props[inc()]) == 1,
		IsDoor:          d2common.StringToUint8(props[inc()]) == 1,
		BlockVisibility: d2common.StringToUint8(props[inc()]) == 1,
		Orientation:     d2common.StringToInt(props[inc()]),
		Trans:           d2common.StringToInt(props[inc()]),

		OrderFlag: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},
		PreOperate: d2common.StringToUint8(props[inc()]) == 1,
		HasAnimationMode: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},

		XOffset: d2common.StringToInt(props[inc()]),
		YOffset: d2common.StringToInt(props[inc()]),
		Draw:    d2common.StringToUint8(props[inc()]) == 1,

		LightRed:   d2common.StringToUint8(props[inc()]),
		LightGreen: d2common.StringToUint8(props[inc()]),
		LightBlue:  d2common.StringToUint8(props[inc()]),

		SelHD: d2common.StringToUint8(props[inc()]) == 1,
		SelTR: d2common.StringToUint8(props[inc()]) == 1,
		SelLG: d2common.StringToUint8(props[inc()]) == 1,
		SelRA: d2common.StringToUint8(props[inc()]) == 1,
		SelLA: d2common.StringToUint8(props[inc()]) == 1,
		SelRH: d2common.StringToUint8(props[inc()]) == 1,
		SelLH: d2common.StringToUint8(props[inc()]) == 1,
		SelSH: d2common.StringToUint8(props[inc()]) == 1,
		SelS: [8]bool{
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
			d2common.StringToUint8(props[inc()]) == 1,
		},

		TotalPieces: d2common.StringToInt(props[inc()]),
		SubClass:    d2common.StringToInt(props[inc()]),

		XSpace: d2common.StringToInt(props[inc()]),
		YSpace: d2common.StringToInt(props[inc()]),

		NameOffset: d2common.StringToInt(props[inc()]),

		MonsterOk:      d2common.StringToUint8(props[inc()]) == 1,
		OperateRange:   d2common.StringToInt(props[inc()]),
		ShrineFunction: d2common.StringToInt(props[inc()]),
		Restore:        d2common.StringToUint8(props[inc()]) == 1,

		Parm: [8]int{
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
			d2common.StringToInt(props[inc()]),
		},
		Act:            d2common.StringToInt(props[inc()]),
		Lockable:       d2common.StringToUint8(props[inc()]) == 1,
		Gore:           d2common.StringToUint8(props[inc()]) == 1,
		Sync:           d2common.StringToUint8(props[inc()]) == 1,
		Flicker:        d2common.StringToUint8(props[inc()]) == 1,
		Damage:         d2common.StringToInt(props[inc()]),
		Beta:           d2common.StringToUint8(props[inc()]) == 1,
		Overlay:        d2common.StringToUint8(props[inc()]) == 1,
		CollisionSubst: d2common.StringToUint8(props[inc()]) == 1,

		Left:   d2common.StringToInt(props[inc()]),
		Top:    d2common.StringToInt(props[inc()]),
		Width:  d2common.StringToInt(props[inc()]),
		Height: d2common.StringToInt(props[inc()]),

		OperateFn:  d2common.StringToInt(props[inc()]),
		PopulateFn: d2common.StringToInt(props[inc()]),
		InitFn:     d2common.StringToInt(props[inc()]),
		ClientFn:   d2common.StringToInt(props[inc()]),

		RestoreVirgins: d2common.StringToUint8(props[inc()]) == 1,
		BlockMissile:   d2common.StringToUint8(props[inc()]) == 1,
		DrawUnder:      d2common.StringToUint8(props[inc()]) == 1,
		OpenWarp:       d2common.StringToUint8(props[inc()]) == 1,

		AutoMap: d2common.StringToInt(props[inc()]),
	}

	return result
}

//nolint:gochecknoglobals // Currently global by design, only written once
var Objects map[int]*ObjectRecord

func LoadObjects(file []byte) {
	Objects = make(map[int]*ObjectRecord)
	data := strings.Split(string(file), "\r\n")[1:]

	for _, line := range data {
		if line == "" {
			continue
		}

		props := strings.Split(line, "\t")

		if props[2] == "" {
			continue // skip a line that doesn't have an id
		}

		rec := createObjectRecord(props)
		Objects[rec.Id] = &rec
	}

	log.Printf("Loaded %d objects", len(Objects))
}
