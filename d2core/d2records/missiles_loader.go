package d2records

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
)

// nolint:funlen // cant reduce
func missilesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Missiles)
	r.missilesByName = make(missilesByName)

	for d.Next() {
		record := &MissileRecord{
			Name: d.String("Missile"),
			Id:   d.Number("Id"),

			ClientMovementFunc:  d.Number("pCltDoFunc"),
			ClientCollisionFunc: d.Number("pCltHitFunc"),
			ServerMovementFunc:  d.Number("pSrvDoFunc"),
			ServerCollisionFunc: d.Number("pSrvHitFunc"),
			ServerDamageFunc:    d.Number("pSrvDmgFunc"),

			ServerMovementCalc: MissileCalc{
				Calc: "SrvCalc1",
				Desc: "*srv calc 1 desc",
				Params: []MissileCalcParam{
					{
						d.Number("Param1"),
						d.String("*param1 desc"),
					},
					{
						d.Number("Param2"),
						d.String("*param2 desc"),
					},
					{
						d.Number("Param3"),
						d.String("*param3 desc"),
					},
					{
						d.Number("Param4"),
						d.String("*param4 desc"),
					},
					{
						d.Number("Param5"),
						d.String("*param5 desc"),
					},
				},
			},

			ClientMovementCalc: MissileCalc{
				Calc: "CltCalc1",
				Desc: "*client calc 1 desc",
				Params: []MissileCalcParam{
					{
						d.Number("CltParam1"),
						d.String("*client param1 desc"),
					},
					{
						d.Number("CltParam2"),
						d.String("*client param2 desc"),
					},
					{
						d.Number("CltParam3"),
						d.String("*client param3 desc"),
					},
					{
						d.Number("CltParam4"),
						d.String("*client param4 desc"),
					},
					{
						d.Number("CltParam5"),
						d.String("*client param5 desc"),
					},
				},
			},

			ServerCollisionCalc: MissileCalc{
				Calc: "SHitCalc1",
				Desc: "*server hit calc 1 desc",
				Params: []MissileCalcParam{
					{
						d.Number("sHitPar1"),
						d.String("*server hit param1 desc"),
					},
					{
						d.Number("sHitPar2"),
						d.String("*server hit param2 desc"),
					},
					{
						d.Number("sHitPar3"),
						d.String("*server hit param3 desc"),
					},
				},
			},

			ClientCollisionCalc: MissileCalc{
				Calc: "CHitCalc1",
				Desc: "*client hit calc 1 desc",
				Params: []MissileCalcParam{
					{
						d.Number("cHitPar1"),
						d.String("*client hit param1 desc"),
					},
					{
						d.Number("cHitPar2"),
						d.String("*client hit param2 desc"),
					},
					{
						d.Number("cHitPar3"),
						d.String("*client hit param3 desc"),
					},
				},
			},

			ServerDamageCalc: MissileCalc{
				Calc: "DmgCalc1",
				Desc: "*damage calc 1",
				Params: []MissileCalcParam{
					{
						d.Number("dParam1"),
						d.String("*damage param1 desc"),
					},
					{
						d.Number("dParam2"),
						d.String("*damage param2 desc"),
					},
				},
			},

			Velocity:           d.Number("Vel"),
			MaxVelocity:        d.Number("MaxVel"),
			LevelVelocityBonus: d.Number("VelLev"),
			Accel:              d.Number("Accel"),
			Range:              d.Number("Range"),
			LevelRangeBonus:    d.Number("LevRange"),

			Light: MissileLight{
				Diameter: d.Number("Light"),
				Flicker:  d.Number("Flicker"),
				Red:      uint8(d.Number("Red")),
				Green:    uint8(d.Number("Green")),
				Blue:     uint8(d.Number("Blue")),
			},

			Animation: MissileAnimation{
				StepsBeforeVisible: d.Number("InitSteps"),
				StepsBeforeActive:  d.Number("Activate"),
				LoopAnimation:      d.Number("LoopAnim") > 0,
				CelFileName:        d.String("CelFile"),
				AnimationRate:      d.Number("animrate"),
				AnimationLength:    d.Number("AnimLen"),
				AnimationSpeed:     d.Number("AnimSpeed"),
				StartingFrame:      d.Number("RandStart"),
				HasSubLoop:         d.Number("SubLoop") > 0,
				SubStartingFrame:   d.Number("SubStart"),
				SubEndingFrame:     d.Number("SubStop"),
			},

			Collision: MissileCollision{
				CollisionType:          d.Number("CollideType"),
				DestroyedUponCollision: d.Number("CollideKill") > 0,
				FriendlyFire:           d.Number("CollideFriend") > 0,
				LastCollide:            d.Number("LastCollide") > 0,
				Collision:              d.Number("Collision") > 0,
				ClientCollision:        d.Number("ClientCol") > 0,
				ClientSend:             d.Number("ClientSend") > 0,
				UseCollisionTimer:      d.Number("NextHit") > 0,
				TimerFrames:            d.Number("NextDelay"),
			},

			XOffset: d.Number("xoffset"),
			YOffset: d.Number("yoffset"),
			ZOffset: d.Number("zoffset"),
			Size:    d.Number("Size"),

			DestroyedByTP:      d.Number("SrcTown") > 0,
			DestroyedByTPFrame: d.Number("CltSrcTown"),
			CanDestroy:         d.Number("CanDestroy") > 0,

			UseAttackRating: d.Number("ToHit") > 0,
			AlwaysExplode:   d.Number("AlwaysExplode") > 0,

			ClientExplosion:     d.Number("Explosion") > 0,
			TownSafe:            d.Number("Town") > 0,
			IgnoreBossModifiers: d.Number("NoUniqueMod") > 0,
			IgnoreMultishot:     d.Number("NoMultiShot") > 0,
			HolyFilterType:      d.Number("Holy"),
			CanBeSlowed:         d.Number("CanSlow") > 0,
			TriggersHitEvents:   d.Number("ReturnFire") > 0,
			TriggersGetHit:      d.Number("GetHit") > 0,
			SoftHit:             d.Number("SoftHit") > 0,
			KnockbackPercent:    d.Number("KnockBack"),

			TransparencyMode: d.Number("Trans"),

			UseQuantity:      d.Number("Qty") > 0,
			AffectedByPierce: d.Number("Pierce") > 0,
			SpecialSetup:     d.Number("SpecialSetup") > 0,

			MissileSkill: d.Number("MissileSkill") > 0,
			SkillName:    d.String("Skill"),

			ResultFlags: d.Number("ResultFlags"),
			HitFlags:    d.Number("HitFlags"),

			HitShift:               d.Number("HitShift"),
			ApplyMastery:           d.Number("ApplyMastery") > 0,
			SourceDamage:           d.Number("SrcDamage"),
			HalfDamageForTwoHander: d.Number("Half2HSrc") > 0,
			SourceMissDamage:       d.Number("SrcMissDmg"),

			Damage: MissileDamage{
				MinDamage: d.Number("MinDamage"),
				MinLevelDamage: [5]int{
					d.Number("MinLevDam1"),
					d.Number("MinLevDam2"),
					d.Number("MinLevDam3"),
					d.Number("MinLevDam4"),
					d.Number("MinLevDam5"),
				},
				MaxDamage: d.Number("MaxDamage"),
				MaxLevelDamage: [5]int{
					d.Number("MaxLevDam1"),
					d.Number("MaxLevDam2"),
					d.Number("MaxLevDam3"),
					d.Number("MaxLevDam4"),
					d.Number("MaxLevDam5"),
				},
				DamageSynergyPerCalc: d2calculation.CalcString(d.String("DmgSymPerCalc")),
			},
			ElementalDamage: MissileElementalDamage{
				ElementType: d.String("EType"),
				Damage: MissileDamage{
					MinDamage: d.Number("MinEDamage"),
					MinLevelDamage: [5]int{
						d.Number("MinELevDam1"),
						d.Number("MinELevDam2"),
						d.Number("MinELevDam3"),
						d.Number("MinELevDam4"),
						d.Number("MinELevDam5"),
					},
					MaxDamage: d.Number("MaxEDamage"),
					MaxLevelDamage: [5]int{
						d.Number("MaxELevDam1"),
						d.Number("MaxELevDam2"),
						d.Number("MaxELevDam3"),
						d.Number("MaxELevDam4"),
						d.Number("MaxELevDam5"),
					},
					DamageSynergyPerCalc: d2calculation.CalcString(d.String("EDmgSymPerCalc")),
				},
				Duration: d.Number("ELen"),
				LevelDuration: [3]int{
					d.Number("ELevLen1"),
					d.Number("ELevLen2"),
					d.Number("ELevLen3"),
				},
			},

			HitClass:            d.Number("HitClass"),
			NumDirections:       d.Number("NumDirections"),
			LocalBlood:          d.Number("LocalBlood"),
			DamageReductionRate: d.Number("DamageRate"),

			TravelSound:      d.String("TravelSound"),
			HitSound:         d.String("HitSound"),
			ProgSound:        d.String("ProgSound"),
			ProgOverlay:      d.String("ProgOverlay"),
			ExplosionMissile: d.String("ExplosionMissile"),

			SubMissile: [3]string{
				d.String("SubMissile1"),
				d.String("SubMissile2"),
				d.String("SubMissile3"),
			},
			HitSubMissile: [4]string{
				d.String("HitSubMissile1"),
				d.String("HitSubMissile2"),
				d.String("HitSubMissile3"),
				d.String("HitSubMissile4"),
			},
			ClientSubMissile: [3]string{
				d.String("CltSubMissile1"),
				d.String("CltSubMissile2"),
				d.String("CltSubMissile3"),
			},
			ClientHitSubMissile: [4]string{
				d.String("CltHitSubMissile1"),
				d.String("CltHitSubMissile2"),
				d.String("CltHitSubMissile3"),
				d.String("CltHitSubMissile4"),
			},
		}

		records[record.Id] = record
		r.missilesByName[sanitizeMissilesKey(record.Name)] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Missile Records", len(records))

	r.Missiles = records

	return nil
}

func sanitizeMissilesKey(missileName string) string {
	return strings.ToLower(strings.ReplaceAll(missileName, " ", ""))
}
