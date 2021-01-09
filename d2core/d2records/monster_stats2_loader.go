package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonStats2 loads MonStats2Records from monstats2.txt
//nolint:funlen //just a big data loader
func monsterStats2Loader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonStats2)

	for d.Next() {
		resurrectMode, err := monsterAnimationModeFromString(d.String("ResurrectMode"))
		if err != nil {
			return err
		}

		record := &MonStats2Record{
			Key:             d.String("Id"),
			Height:          d.Number("Height"),
			OverlayHeight:   d.Number("OverlayHeight"),
			PixelHeight:     d.Number("pixHeight"),
			SizeX:           d.Number("SizeX"),
			SizeY:           d.Number("SizeY"),
			SpawnMethod:     d.Number("spawnCol"),
			MeleeRng:        d.Number("MeleeRng"),
			BaseWeaponClass: d.String("BaseW"),
			HitClass:        d.Number("HitClass"),
			EquipmentOptions: [16][]string{
				d.List("HDv"),
				d.List("TRv"),
				d.List("LGv"),
				d.List("Rav"),
				d.List("Lav"),
				d.List("RHv"),
				d.List("LHv"),
				d.List("SHv"),
				d.List("S1v"),
				d.List("S2v"),
				d.List("S3v"),
				d.List("S4v"),
				d.List("S5v"),
				d.List("S6v"),
				d.List("S7v"),
				d.List("S8v"),
			},
			HasComponent: [16]bool{
				d.Bool("HD"),
				d.Bool("TR"),
				d.Bool("LG"),
				d.Bool("RA"),
				d.Bool("LA"),
				d.Bool("RH"),
				d.Bool("LH"),
				d.Bool("SH"),
				d.Bool("S1"),
				d.Bool("S2"),
				d.Bool("S3"),
				d.Bool("S4"),
				d.Bool("S5"),
				d.Bool("S6"),
				d.Bool("S7"),
				d.Bool("S8"),
			},
			TotalPieces: d.Number("TotalPieces"),
			HasAnimationMode: [16]bool{
				d.Bool("mDT"),
				d.Bool("mNU"),
				d.Bool("mWL"),
				d.Bool("mGH"),
				d.Bool("mA1"),
				d.Bool("mA2"),
				d.Bool("mBL"),
				d.Bool("mSC"),
				d.Bool("mS1"),
				d.Bool("mS2"),
				d.Bool("mS3"),
				d.Bool("mS4"),
				d.Bool("mDD"),
				d.Bool("mKB"),
				d.Bool("mSQ"),
				d.Bool("mRN"),
			},
			DirectionsPerMode: [16]int{
				d.Number("dDT"),
				d.Number("dNU"),
				d.Number("dWL"),
				d.Number("dGH"),
				d.Number("dA1"),
				d.Number("dA2"),
				d.Number("dBL"),
				d.Number("dSC"),
				d.Number("dS1"),
				d.Number("dS2"),
				d.Number("dS3"),
				d.Number("dS4"),
				d.Number("dDD"),
				d.Number("dKB"),
				d.Number("dSQ"),
				d.Number("dRN"),
			},
			A1mv:               d.Bool("A1mv"),
			A2mv:               d.Bool("A2mv"),
			SCmv:               d.Bool("SCmv"),
			S1mv:               d.Bool("S1mv"),
			S2mv:               d.Bool("S2mv"),
			S3mv:               d.Bool("S3mv"),
			S4mv:               d.Bool("S4mv"),
			NoGfxHitTest:       d.Bool("noGfxHitTest"),
			BoxTop:             d.Number("htTop"),
			BoxLeft:            d.Number("htLeft"),
			BoxWidth:           d.Number("htWidth"),
			BoxHeight:          d.Number("htHeight"),
			Restore:            d.Number("restore"),
			AutomapCel:         d.Number("automapCel"),
			NoMap:              d.Bool("noMap"),
			NoOvly:             d.Bool("noOvly"),
			IsSelectable:       d.Bool("isSel"),
			AllySelectable:     d.Bool("alSel"),
			shiftSel:           d.Bool("shiftSel"),
			NotSelectable:      d.Bool("noSel"),
			IsCorpseSelectable: d.Bool("corpseSel"),
			IsAttackable:       d.Bool("isAtt"),
			IsRevivable:        d.Bool("revive"),
			IsCritter:          d.Bool("critter"),
			IsSmall:            d.Bool("small"),
			IsLarge:            d.Bool("large"),
			IsSoft:             d.Bool("soft"),
			IsInert:            d.Bool("inert"),
			objCol:             d.Bool("objCol"),
			IsCorpseCollidable: d.Bool("deadCol"),
			IsCorpseWalkable:   d.Bool("unflatDead"),
			HasShadow:          d.Bool("Shadow"),
			NoUniqueShift:      d.Bool("noUniqueShift"),
			CompositeDeath:     d.Bool("compositeDeath"),
			LocalBlood:         d.Number("localBlood"),
			Bleed:              d.Number("Bleed"),
			Light:              d.Number("Light"),
			LightR:             d.Number("light-r"),
			LightG:             d.Number("light-g"),
			lightB:             d.Number("light-b"),
			NormalPalette:      d.Number("Utrans"),
			NightmarePalette:   d.Number("Utrans(N)"),
			HellPalatte:        d.Number("Utrans(H)"),
			Heart:              d.String("Heart"),
			BodyPart:           d.String("BodyPart"),
			InfernoLen:         d.Number("InfernoLen"),
			InfernoAnim:        d.Number("InfernoAnim"),
			InfernoRollback:    d.Number("InfernoRollback"),
			ResurrectMode:      resurrectMode,
			ResurrectSkill:     d.String("ResurrectSkill"),
		}

		records[record.Key] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Logger.Infof("Loaded %d MonStats2 records", len(records))

	r.Monster.Stats2 = records

	return nil
}

//nolint:gochecknoglobals // better for lookup
var monsterAnimationModeLookup = map[string]d2enum.MonsterAnimationMode{
	d2enum.MonsterAnimationModeNeutral.String():  d2enum.MonsterAnimationModeNeutral,
	d2enum.MonsterAnimationModeSkill1.String():   d2enum.MonsterAnimationModeSkill1,
	d2enum.MonsterAnimationModeSequence.String(): d2enum.MonsterAnimationModeSequence,
}

func monsterAnimationModeFromString(s string) (d2enum.MonsterAnimationMode, error) {
	v, ok := monsterAnimationModeLookup[s]
	if !ok {
		return d2enum.MonsterAnimationModeNeutral, fmt.Errorf("unhandled MonsterAnimationMode %q", s)
	}

	return v, nil
}
