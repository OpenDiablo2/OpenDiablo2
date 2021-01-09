package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

//nolint:funlen // Makes no sense to split
func objectDetailsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ObjectDetails)

	i := 0

	for d.Next() {
		record := &ObjectDetailsRecord{
			Index:       i,
			Name:        d.String("Name"),
			Description: d.String("description - not loaded"),
			id:          d.Number("Id"),
			token:       d.String("Token"),

			SpawnMax: d.Number("SpawnMax"),
			Selectable: [8]bool{
				d.Number("Selectable0") == 1,
				d.Number("Selectable1") == 1,
				d.Number("Selectable2") == 1,
				d.Number("Selectable3") == 1,
				d.Number("Selectable4") == 1,
				d.Number("Selectable5") == 1,
				d.Number("Selectable6") == 1,
				d.Number("Selectable7") == 1,
			},
			TrapProbability: d.Number("TrapProb"),

			SizeX: d.Number("SizeX"),
			SizeY: d.Number("SizeY"),

			NTgtFX: d.Number("nTgtFX"),
			NTgtFY: d.Number("nTgtFY"),
			NTgtBX: d.Number("nTgtBX"),
			NTgtBY: d.Number("nTgtBY"),

			FrameCount: [8]int{
				d.Number("FrameCnt0"),
				d.Number("FrameCnt1"),
				d.Number("FrameCnt2"),
				d.Number("FrameCnt3"),
				d.Number("FrameCnt4"),
				d.Number("FrameCnt5"),
				d.Number("FrameCnt6"),
				d.Number("FrameCnt7"),
			},
			FrameDelta: [8]int{
				d.Number("FrameDelta0"),
				d.Number("FrameDelta1"),
				d.Number("FrameDelta2"),
				d.Number("FrameDelta3"),
				d.Number("FrameDelta4"),
				d.Number("FrameDelta5"),
				d.Number("FrameDelta6"),
				d.Number("FrameDelta7"),
			},
			CycleAnimation: [8]bool{
				d.Number("CycleAnim0") == 1,
				d.Number("CycleAnim1") == 1,
				d.Number("CycleAnim2") == 1,
				d.Number("CycleAnim3") == 1,
				d.Number("CycleAnim4") == 1,
				d.Number("CycleAnim5") == 1,
				d.Number("CycleAnim6") == 1,
				d.Number("CycleAnim7") == 1,
			},
			LightDiameter: [8]int{
				d.Number("Lit0"),
				d.Number("Lit1"),
				d.Number("Lit2"),
				d.Number("Lit3"),
				d.Number("Lit4"),
				d.Number("Lit5"),
				d.Number("Lit6"),
				d.Number("Lit7"),
			},
			BlocksLight: [8]bool{
				d.Number("BlocksLight0") == 1,
				d.Number("BlocksLight1") == 1,
				d.Number("BlocksLight2") == 1,
				d.Number("BlocksLight3") == 1,
				d.Number("BlocksLight4") == 1,
				d.Number("BlocksLight5") == 1,
				d.Number("BlocksLight6") == 1,
				d.Number("BlocksLight7") == 1,
			},
			HasCollision: [8]bool{
				d.Number("HasCollision0") == 1,
				d.Number("HasCollision1") == 1,
				d.Number("HasCollision2") == 1,
				d.Number("HasCollision3") == 1,
				d.Number("HasCollision4") == 1,
				d.Number("HasCollision5") == 1,
				d.Number("HasCollision6") == 1,
				d.Number("HasCollision7") == 1,
			},
			IsAttackable: d.Number("IsAttackable0") == 1,
			StartFrame: [8]int{
				d.Number("Start0"),
				d.Number("Start1"),
				d.Number("Start2"),
				d.Number("Start3"),
				d.Number("Start4"),
				d.Number("Start5"),
				d.Number("Start6"),
				d.Number("Start7"),
			},

			EnvEffect:       d.Number("EnvEffect") == 1,
			IsDoor:          d.Number("IsDoor") == 1,
			BlockVisibility: d.Number("BlocksVis") == 1,
			Orientation:     d.Number("Orientation"),
			Trans:           d.Number("Trans"),

			OrderFlag: [8]int{
				d.Number("OrderFlag0"),
				d.Number("OrderFlag1"),
				d.Number("OrderFlag2"),
				d.Number("OrderFlag3"),
				d.Number("OrderFlag4"),
				d.Number("OrderFlag5"),
				d.Number("OrderFlag6"),
				d.Number("OrderFlag7"),
			},
			PreOperate: d.Number("PreOperate") == 1,
			HasAnimationMode: [8]bool{
				d.Number("Mode0") == 1,
				d.Number("Mode1") == 1,
				d.Number("Mode2") == 1,
				d.Number("Mode3") == 1,
				d.Number("Mode4") == 1,
				d.Number("Mode5") == 1,
				d.Number("Mode6") == 1,
				d.Number("Mode7") == 1,
			},

			XOffset: d.Number("Yoffset"),
			YOffset: d.Number("Xoffset"),
			Draw:    d.Number("Draw") == 1,

			LightRed:   uint8(d.Number("Red")),
			LightGreen: uint8(d.Number("Green")),
			LightBlue:  uint8(d.Number("Blue")),

			SelHD: d.Number("HD") == 1,
			SelTR: d.Number("TR") == 1,
			SelLG: d.Number("LG") == 1,
			SelRA: d.Number("RA") == 1,
			SelLA: d.Number("LA") == 1,
			SelRH: d.Number("RH") == 1,
			SelLH: d.Number("LH") == 1,
			SelSH: d.Number("SH") == 1,
			SelS: [8]bool{
				d.Number("S1") == 1,
				d.Number("S2") == 1,
				d.Number("S3") == 1,
				d.Number("S4") == 1,
				d.Number("S5") == 1,
				d.Number("S6") == 1,
				d.Number("S7") == 1,
				d.Number("S8") == 1,
			},

			TotalPieces: d.Number("TotalPieces"),
			SubClass:    d.Number("SubClass"),

			XSpace: d.Number("Xspace"),
			YSpace: d.Number("Yspace"),

			NameOffset: d.Number("NameOffset"),

			MonsterOk:      uint8(d.Number("MonsterOK")) == 1,
			OperateRange:   d.Number("OperateRange"),
			ShrineFunction: d.Number("ShrineFunction"),
			Restore:        uint8(d.Number("Restore")) == 1,

			Parm: [8]int{
				d.Number("Parm0"),
				d.Number("Parm1"),
				d.Number("Parm2"),
				d.Number("Parm3"),
				d.Number("Parm4"),
				d.Number("Parm5"),
				d.Number("Parm6"),
				d.Number("Parm7"),
			},
			Act:            d.Number("Act"),
			Lockable:       uint8(d.Number("Lockable")) == 1,
			Gore:           uint8(d.Number("Gore")) == 1,
			Sync:           uint8(d.Number("Sync")) == 1,
			Flicker:        uint8(d.Number("Flicker")) == 1,
			Damage:         d.Number("Damage"),
			Beta:           uint8(d.Number("Beta")) == 1,
			Overlay:        uint8(d.Number("Overlay")) == 1,
			CollisionSubst: uint8(d.Number("CollisionSubst")) == 1,

			Left:   d.Number("Left"),
			Top:    d.Number("Top"),
			Width:  d.Number("Width"),
			Height: d.Number("Height"),

			OperateFn:  d.Number("OperateFn"),
			PopulateFn: d.Number("PopulateFn"),
			InitFn:     d.Number("InitFn"),
			ClientFn:   d.Number("ClientFn"),

			RestoreVirgins: uint8(d.Number("RestoreVirgins")) == 1,
			BlockMissile:   uint8(d.Number("BlockMissile")) == 1,
			DrawUnder:      uint8(d.Number("DrawUnder")) == 1,
			OpenWarp:       uint8(d.Number("OpenWarp")) == 1,

			AutoMap: d.Number("AutoMap"),
		}

		records[i] = record
		i++
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d objects", len(records))

	r.Object.Details = records

	return nil
}
