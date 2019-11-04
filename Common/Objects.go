package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
)

// An ObjectRecord represents the settings for one type of object from objects.txt
type ObjectRecord struct {
	Name string
	Description string
	Id int
	Token string // refers to what graphics this object uses

	SpawnMax int // unused? 
	Selectable [8]bool // is this mode selectable
	TrapProbability int // unused

	SizeX int
	SizeY int

	NTgtFX int // unknown
	NTgtFY int // unknown
	NTgtBX int // unknown
	NTgtBY int // unknown

	FrameCount [8]int // how many frames does this mode have, 0 = skip
	FrameDelta [8]int // what rate is the animation played at (256 = 100% speed)
	CycleAnimation [8]bool // probably whether animation loops
	LightDiameter [8]int
	BlocksLight [8]bool 
	HasCollision [8]bool
	IsAttackable bool // do we kick it when interacting
	StartFrame [8]int

	EnvEffect bool // unknown
	IsDoor bool
	BlockVisibility bool // only works with IsDoor
	Orientation int // unknown (1=sw, 2=nw, 3=se, 4=ne)
	Trans int // controls palette mapping

	OrderFlag [8]int //  0 = object, 1 = floor, 2 = wall
	PreOperate bool // unknown
	HasAnimationMode [8]bool // 'Mode' in source, true if this mode is used

	XOffset int // in pixels offset
	YOffset int
	Draw bool // if false, object isn't drawn (shadow is still drawn and player can still select though)

	LightRed byte // if lightdiameter is set, rgb of the light
	LightGreen byte
	LightBlue byte

	SelHD bool // whether these DCC components are selectable
	SelTR bool
	SelLG bool
	SelRA bool
	SelLA bool
	SelRH bool
	SelLH bool
	SelSH bool
	SelS [8]bool

	TotalPieces int // selectable DCC components count
	SubClass int // subclass of object:
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

	MonsterOk bool // unknown
	OperateRange int // distance object can be used from, might be unused
	ShrineFunction int // unused
	Restore bool // if true, object is stored in memory and will be retained if you leave and re-enter the area
	
	Parm [8]int // unknown
	Act int // what acts this object can appear in (15 = all three)
	Lockable bool
	Gore bool // unknown, something with corpses
	Sync bool // unknown
	Flicker bool // light flickers if true
	Damage int // amount of damage done by this (used depending on operatefn)
	Beta bool // if true, appeared in the beta?
	Overlay bool // unknown
	CollisionSubst bool // unknown, controls some kind of special collision checking?

	Left int // unknown, clickable bounding box?
	Top int
	Width int
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
	BlockMissile bool // if true, missiles collide with this
	DrawUnder bool // if true, drawn as a floor tile is
	OpenWarp bool // needs clarification, controls whether highlighting shows
	// 'To ...' or 'trap door' when highlighting, not sure which is T/F
	AutoMap int // controls how this object appears on the map
	// 0 = it doesn't, rest of modes need to be analyzed
}

// CreateObjectRecord parses a row from objects.txt into an object record
func createObjectRecord(props []string) ObjectRecord {
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := ObjectRecord{
		Name: props[inc()],
		Description: props[inc()],
		Id: StringToInt(props[inc()]),
		Token: props[inc()], 
	
		SpawnMax: StringToInt(props[inc()]), 
		Selectable: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
		TrapProbability: StringToInt(props[inc()]), 
	
		SizeX: StringToInt(props[inc()]),
		SizeY: StringToInt(props[inc()]),
	
		NTgtFX: StringToInt(props[inc()]), 
		NTgtFY: StringToInt(props[inc()]), 
		NTgtBX: StringToInt(props[inc()]), 
		NTgtBY: StringToInt(props[inc()]), 
	
		FrameCount: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
		FrameDelta: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
		CycleAnimation: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
		LightDiameter: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
		BlocksLight: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
		HasCollision: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
		IsAttackable: StringToUint8(props[inc()]) == 1, 
		StartFrame: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
	
		EnvEffect: StringToUint8(props[inc()]) == 1, 
		IsDoor: StringToUint8(props[inc()]) == 1,
		BlockVisibility: StringToUint8(props[inc()]) == 1, 
		Orientation: StringToInt(props[inc()]), 
		Trans: StringToInt(props[inc()]), 
	
		OrderFlag: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
		PreOperate: StringToUint8(props[inc()]) == 1, 
		HasAnimationMode: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
	
		XOffset: StringToInt(props[inc()]), 
		YOffset: StringToInt(props[inc()]),
		Draw: StringToUint8(props[inc()]) == 1, 
	
		LightRed: StringToUint8(props[inc()]),
		LightGreen: StringToUint8(props[inc()]),
		LightBlue: StringToUint8(props[inc()]),
	
		SelHD: StringToUint8(props[inc()]) == 1, 
		SelTR: StringToUint8(props[inc()]) == 1,
		SelLG: StringToUint8(props[inc()]) == 1,
		SelRA: StringToUint8(props[inc()]) == 1,
		SelLA: StringToUint8(props[inc()]) == 1,
		SelRH: StringToUint8(props[inc()]) == 1,
		SelLH: StringToUint8(props[inc()]) == 1,
		SelSH: StringToUint8(props[inc()]) == 1,
		SelS: [8]bool{
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
			StringToUint8(props[inc()]) == 1,
		},
	
		TotalPieces: StringToInt(props[inc()]), 
		SubClass: StringToInt(props[inc()]), 
		
		XSpace: StringToInt(props[inc()]), 
		YSpace: StringToInt(props[inc()]),
	
		NameOffset: StringToInt(props[inc()]), 
	
		MonsterOk: StringToUint8(props[inc()]) == 1, 
		OperateRange: StringToInt(props[inc()]), 
		ShrineFunction: StringToInt(props[inc()]), 
		Restore: StringToUint8(props[inc()]) == 1, 
		
		Parm: [8]int{
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
			StringToInt(props[inc()]),
		},
		Act: StringToInt(props[inc()]), 
		Lockable: StringToUint8(props[inc()]) == 1,
		Gore: StringToUint8(props[inc()]) == 1, 
		Sync: StringToUint8(props[inc()]) == 1, 
		Flicker: StringToUint8(props[inc()]) == 1, 
		Damage: StringToInt(props[inc()]), 
		Beta: StringToUint8(props[inc()]) == 1, 
		Overlay: StringToUint8(props[inc()]) == 1, 
		CollisionSubst: StringToUint8(props[inc()]) == 1, 
	
		Left: StringToInt(props[inc()]), 
		Top: StringToInt(props[inc()]),
		Width: StringToInt(props[inc()]),
		Height: StringToInt(props[inc()]),
	
		OperateFn: StringToInt(props[inc()]), 
		PopulateFn: StringToInt(props[inc()]), 
		InitFn: StringToInt(props[inc()]), 
		ClientFn: StringToInt(props[inc()]), 
		
		RestoreVirgins: StringToUint8(props[inc()]) == 1, 
		BlockMissile: StringToUint8(props[inc()]) == 1, 
		DrawUnder: StringToUint8(props[inc()]) == 1, 
		OpenWarp: StringToUint8(props[inc()]) == 1, 
		
		AutoMap: StringToInt(props[inc()]), 
	}
	return result
}

var Objects map[int]*ObjectRecord

func LoadObjects(fileProvider FileProvider) {
	Objects = make(map[int]*ObjectRecord)
	data := strings.Split(string(fileProvider.LoadFile(ResourcePaths.ObjectDetails)), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
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