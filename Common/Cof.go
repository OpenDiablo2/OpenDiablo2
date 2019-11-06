package Common

import "strings"

type AnimationMode int

const (
	AnimationModePlayerDeath       AnimationMode = 0
	AnimationModePlayerNeutral     AnimationMode = 1
	AnimationModePlayerWalk        AnimationMode = 2
	AnimationModePlayerRun         AnimationMode = 3
	AnimationModePlayerGetHit      AnimationMode = 4
	AnimationModePlayerTownNeutral AnimationMode = 5
	AnimationModePlayerTownWalk    AnimationMode = 6
	AnimationModePlayerAttack1     AnimationMode = 7
	AnimationModePlayerAttack2     AnimationMode = 8
	AnimationModePlayerBlock       AnimationMode = 9
	AnimationModePlayerCast        AnimationMode = 10
	AnimationModePlayerThrow       AnimationMode = 11
	AnimationModePlayerKick        AnimationMode = 12
	AnimationModePlayerSkill1      AnimationMode = 13
	AnimationModePlayerSkill2      AnimationMode = 14
	AnimationModePlayerSkill3      AnimationMode = 15
	AnimationModePlayerSkill4      AnimationMode = 16
	AnimationModePlayerDead        AnimationMode = 17
	AnimationModePlayerSequence    AnimationMode = 18
	AnimationModePlayerKnockBack   AnimationMode = 19
	AnimationModeMonsterDeath      AnimationMode = 20
	AnimationModeMonsterNeutral    AnimationMode = 21
	AnimationModeMonsterWalk       AnimationMode = 22
	AnimationModeMonsterGetHit     AnimationMode = 23
	AnimationModeMonsterAttack1    AnimationMode = 24
	AnimationModeMonsterAttack2    AnimationMode = 25
	AnimationModeMonsterBlock      AnimationMode = 26
	AnimationModeMonsterCast       AnimationMode = 27
	AnimationModeMonsterSkill1     AnimationMode = 28
	AnimationModeMonsterSkill2     AnimationMode = 29
	AnimationModeMonsterSkill3     AnimationMode = 30
	AnimationModeMonsterSkill4     AnimationMode = 31
	AnimationModeMonsterDead       AnimationMode = 32
	AnimationModeMonsterKnockback  AnimationMode = 33
	AnimationModeMonsterSequence   AnimationMode = 34
	AnimationModeMonsterRun        AnimationMode = 35
	AnimationModeObjectNeutral     AnimationMode = 36
	AnimationModeObjectOperating   AnimationMode = 37
	AnimationModeObjectOpened      AnimationMode = 38
	AnimationModeObjectSpecial1    AnimationMode = 39
	AnimationModeObjectSpecial2    AnimationMode = 40
	AnimationModeObjectSpecial3    AnimationMode = 41
	AnimationModeObjectSpecial4    AnimationMode = 42
	AnimationModeObjectSpecial5    AnimationMode = 43
)

var AnimationModeStr = map[AnimationMode]string{
	AnimationModePlayerDeath:       "DT",
	AnimationModePlayerNeutral:     "NU",
	AnimationModePlayerWalk:        "WL",
	AnimationModePlayerRun:         "RN",
	AnimationModePlayerGetHit:      "GH",
	AnimationModePlayerTownNeutral: "TN",
	AnimationModePlayerTownWalk:    "TW",
	AnimationModePlayerAttack1:     "A1",
	AnimationModePlayerAttack2:     "A2",
	AnimationModePlayerBlock:       "BL",
	AnimationModePlayerCast:        "SC",
	AnimationModePlayerThrow:       "TH",
	AnimationModePlayerKick:        "KK",
	AnimationModePlayerSkill1:      "S1",
	AnimationModePlayerSkill2:      "S2",
	AnimationModePlayerSkill3:      "S3",
	AnimationModePlayerSkill4:      "S4",
	AnimationModePlayerDead:        "DD",
	AnimationModePlayerSequence:    "GH",
	AnimationModePlayerKnockBack:   "GH",
	AnimationModeMonsterDeath:      "DT",
	AnimationModeMonsterNeutral:    "NU",
	AnimationModeMonsterWalk:       "WL",
	AnimationModeMonsterGetHit:     "GH",
	AnimationModeMonsterAttack1:    "A1",
	AnimationModeMonsterAttack2:    "A2",
	AnimationModeMonsterBlock:      "BL",
	AnimationModeMonsterCast:       "SC",
	AnimationModeMonsterSkill1:     "S1",
	AnimationModeMonsterSkill2:     "S2",
	AnimationModeMonsterSkill3:     "S3",
	AnimationModeMonsterSkill4:     "S4",
	AnimationModeMonsterDead:       "DD",
	AnimationModeMonsterKnockback:  "GH",
	AnimationModeMonsterSequence:   "xx",
	AnimationModeMonsterRun:        "RN",
	AnimationModeObjectNeutral:     "NU",
	AnimationModeObjectOperating:   "OP",
	AnimationModeObjectOpened:      "ON",
	AnimationModeObjectSpecial1:    "S1",
	AnimationModeObjectSpecial2:    "S2",
	AnimationModeObjectSpecial3:    "S3",
	AnimationModeObjectSpecial4:    "S4",
	AnimationModeObjectSpecial5:    "S5",
}

type CompositeType int

const (
	CompositeTypeHead      CompositeType = 0
	CompositeTypeTorso     CompositeType = 1
	CompositeTypeLegs      CompositeType = 2
	CompositeTypeRightArm  CompositeType = 3
	CompositeTypeLeftArm   CompositeType = 4
	CompositeTypeRightHand CompositeType = 5
	CompositeTypeLeftHand  CompositeType = 6
	CompositeTypeShield    CompositeType = 7
	CompositeTypeSpecial1  CompositeType = 8
	CompositeTypeSpecial2  CompositeType = 9
	CompositeTypeSpecial3  CompositeType = 10
	CompositeTypeSpecial4  CompositeType = 11
	CompositeTypeSpecial5  CompositeType = 12
	CompositeTypeSpecial6  CompositeType = 13
	CompositeTypeSpecial7  CompositeType = 14
	CompositeTypeSpecial8  CompositeType = 15
	CompositeTypeMax       CompositeType = 16
)

type DrawEffect int

const (
	DrawEffectPctTransparency75  = 0 //75 % transparency (colormaps 561-816 in a .pl2)
	DrawEffectPctTransparency50  = 1 //50 % transparency (colormaps 305-560 in a .pl2)
	DrawEffectPctTransparency25  = 2 //25 % transparency (colormaps 49-304 in a .pl2)
	DrawEffectScreen             = 3 //Screen (colormaps 817-1072 in a .pl2)
	DrawEffectLuminance          = 4 //luminance (colormaps 1073-1328 in a .pl2)
	DrawEffectBringAlphaBlending = 5 //bright alpha blending (colormaps 1457-1712 in a .pl2)
)

type WeaponClass int

const (
	WeaponClassNone                 WeaponClass = 0
	WeaponClassHandToHand           WeaponClass = 1
	WeaponClassBow                  WeaponClass = 2
	WeaponClassOneHandSwing         WeaponClass = 3
	WeaponClassOneHandThrust        WeaponClass = 4
	WeaponClassStaff                WeaponClass = 5
	WeaponClassTwoHandSwing         WeaponClass = 6
	WeaponClassTwoHandThrust        WeaponClass = 7
	WeaponClassCrossbow             WeaponClass = 8
	WeaponClassLeftJabRightSwing    WeaponClass = 9
	WeaponClassLeftJabRightThrust   WeaponClass = 10
	WeaponClassLeftSwingRightSwing  WeaponClass = 11
	WeaponClassLeftSwingRightThrust WeaponClass = 12
	WeaponClassOneHandToHand        WeaponClass = 13
	WeaponClassTwoHandToHand        WeaponClass = 14
)

var WeaponClassStr = map[WeaponClass]string{
	WeaponClassNone:                 "",
	WeaponClassHandToHand:           "hth",
	WeaponClassBow:                  "bow",
	WeaponClassOneHandSwing:         "1hs",
	WeaponClassOneHandThrust:        "1ht",
	WeaponClassStaff:                "stf",
	WeaponClassTwoHandSwing:         "2hs",
	WeaponClassTwoHandThrust:        "2ht",
	WeaponClassCrossbow:             "xbw",
	WeaponClassLeftJabRightSwing:    "1js",
	WeaponClassLeftJabRightThrust:   "1jt",
	WeaponClassLeftSwingRightSwing:  "1ss",
	WeaponClassLeftSwingRightThrust: "1st",
	WeaponClassOneHandToHand:        "ht1",
	WeaponClassTwoHandToHand:        "ht2",
}

func GetWeaponClass(val string) WeaponClass {
	for weaponClass, weaponStr := range WeaponClassStr {
		if val != weaponStr {
			continue
		}
		return weaponClass
	}
	return WeaponClassNone
}

type AnimationFrame int

const (
	AnimationFrameNoEvent AnimationFrame = 0
	AnimationFrameAttack  AnimationFrame = 1
	AnimationFrameMissile AnimationFrame = 2
	AnimationFrameSound   AnimationFrame = 3
	AnimationFrameSkill   AnimationFrame = 4
)

type CofLayer struct {
	Type        CompositeType
	Shadow      byte
	Transparent bool
	DrawEffect  DrawEffect
	WeaponClass WeaponClass
}

type Cof struct {
	NumberOfDirections int
	FramesPerDirection int
	NumberOfLayers     int
	CofLayers          []*CofLayer
	CompositeLayers    map[CompositeType]int
	AnimationFrames    []AnimationFrame
	Priority           []CompositeType
}

func LoadCof(fileName string, fileProvider FileProvider) *Cof {
	result := &Cof{}
	fileData := fileProvider.LoadFile(fileName)
	streamReader := CreateStreamReader(fileData)
	result.NumberOfLayers = int(streamReader.GetByte())
	result.FramesPerDirection = int(streamReader.GetByte())
	result.NumberOfDirections = int(streamReader.GetByte())
	streamReader.SkipBytes(25) // Skip 25 unknown bytes...
	result.CofLayers = make([]*CofLayer, 0)
	result.CompositeLayers = make(map[CompositeType]int, 0)
	for i := 0; i < result.NumberOfLayers; i++ {
		layer := &CofLayer{}
		layer.Type = CompositeType(streamReader.GetByte())
		layer.Shadow = streamReader.GetByte()
		streamReader.SkipBytes(1) // Unknown
		layer.Transparent = streamReader.GetByte() != 0
		layer.DrawEffect = DrawEffect(streamReader.GetByte())
		weaponClassStr, _ := streamReader.ReadBytes(4)
		layer.WeaponClass = GetWeaponClass(strings.TrimSpace(strings.ReplaceAll(string(weaponClassStr), string(0), "")))
		result.CofLayers = append(result.CofLayers, layer)
		result.CompositeLayers[layer.Type] = i
	}
	animationFrameBytes, _ := streamReader.ReadBytes(result.FramesPerDirection)
	result.AnimationFrames = make([]AnimationFrame, result.FramesPerDirection)
	for i := range animationFrameBytes {
		result.AnimationFrames[i] = AnimationFrame(animationFrameBytes[i])
	}
	priorityLen := result.FramesPerDirection * result.NumberOfDirections * result.NumberOfLayers
	result.Priority = make([]CompositeType, priorityLen)
	priorityBytes, _ := streamReader.ReadBytes(priorityLen)
	for i := range priorityBytes {
		result.Priority[i] = CompositeType(priorityBytes[i])
	}
	return result
}
