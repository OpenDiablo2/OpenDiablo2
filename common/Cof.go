package common

import "strings"

// Tools used for go:generate.
//
//    go get golang.org/x/tools/cmd/stringer
//    go get github.com/mewspring/tools/cmd/string2enum

//go:generate stringer -linecomment -type AnimationMode

type AnimationMode int

const (
	AnimationModePlayerDeath       AnimationMode = 0 // DT
	AnimationModePlayerNeutral     AnimationMode = 1 // NU
	AnimationModePlayerWalk        AnimationMode = 2 // WL
	AnimationModePlayerRun         AnimationMode = 3 // RN
	AnimationModePlayerGetHit      AnimationMode = 4 // GH
	AnimationModePlayerTownNeutral AnimationMode = 5 // TN
	AnimationModePlayerTownWalk    AnimationMode = 6 // TW
	AnimationModePlayerAttack1     AnimationMode = 7 // A1
	AnimationModePlayerAttack2     AnimationMode = 8 // A2
	AnimationModePlayerBlock       AnimationMode = 9 // BL
	AnimationModePlayerCast        AnimationMode = 10 // SC
	AnimationModePlayerThrow       AnimationMode = 11 // TH
	AnimationModePlayerKick        AnimationMode = 12 // KK
	AnimationModePlayerSkill1      AnimationMode = 13 // S1
	AnimationModePlayerSkill2      AnimationMode = 14 // S2
	AnimationModePlayerSkill3      AnimationMode = 15 // S3
	AnimationModePlayerSkill4      AnimationMode = 16 // S4
	AnimationModePlayerDead        AnimationMode = 17 // DD
	AnimationModePlayerSequence    AnimationMode = 18 // GH
	AnimationModePlayerKnockBack   AnimationMode = 19 // GH
	AnimationModeMonsterDeath      AnimationMode = 20 // DT
	AnimationModeMonsterNeutral    AnimationMode = 21 // NU
	AnimationModeMonsterWalk       AnimationMode = 22 // WL
	AnimationModeMonsterGetHit     AnimationMode = 23 // GH
	AnimationModeMonsterAttack1    AnimationMode = 24 // A1
	AnimationModeMonsterAttack2    AnimationMode = 25 // A2
	AnimationModeMonsterBlock      AnimationMode = 26 // BL
	AnimationModeMonsterCast       AnimationMode = 27 // SC
	AnimationModeMonsterSkill1     AnimationMode = 28 // S1
	AnimationModeMonsterSkill2     AnimationMode = 29 // S2
	AnimationModeMonsterSkill3     AnimationMode = 30 // S3
	AnimationModeMonsterSkill4     AnimationMode = 31 // S4
	AnimationModeMonsterDead       AnimationMode = 32 // DD
	AnimationModeMonsterKnockback  AnimationMode = 33 // GH
	AnimationModeMonsterSequence   AnimationMode = 34 // xx
	AnimationModeMonsterRun        AnimationMode = 35 // RN
	AnimationModeObjectNeutral     AnimationMode = 36 // NU
	AnimationModeObjectOperating   AnimationMode = 37 // OP
	AnimationModeObjectOpened      AnimationMode = 38 // ON
	AnimationModeObjectSpecial1    AnimationMode = 39 // S1
	AnimationModeObjectSpecial2    AnimationMode = 40 // S2
	AnimationModeObjectSpecial3    AnimationMode = 41 // S3
	AnimationModeObjectSpecial4    AnimationMode = 42 // S4
	AnimationModeObjectSpecial5    AnimationMode = 43 // S5
)

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

//go:generate stringer -linecomment -type WeaponClass
//go:generate string2enum -samepkg -linecomment -type WeaponClass

type WeaponClass int

const (
	WeaponClassNone                 WeaponClass = 0 //
	WeaponClassHandToHand           WeaponClass = 1 // hth
	WeaponClassBow                  WeaponClass = 2 // bow
	WeaponClassOneHandSwing         WeaponClass = 3 // 1hs
	WeaponClassOneHandThrust        WeaponClass = 4 // 1ht
	WeaponClassStaff                WeaponClass = 5 // stf
	WeaponClassTwoHandSwing         WeaponClass = 6 // 2hs
	WeaponClassTwoHandThrust        WeaponClass = 7 // 2ht
	WeaponClassCrossbow             WeaponClass = 8 // xbw
	WeaponClassLeftJabRightSwing    WeaponClass = 9 // 1js
	WeaponClassLeftJabRightThrust   WeaponClass = 10 // 1jt
	WeaponClassLeftSwingRightSwing  WeaponClass = 11 // 1ss
	WeaponClassLeftSwingRightThrust WeaponClass = 12 // 1st
	WeaponClassOneHandToHand        WeaponClass = 13 // ht1
	WeaponClassTwoHandToHand        WeaponClass = 14 // ht2
)

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
		layer.WeaponClass = WeaponClassFromString(strings.TrimSpace(strings.ReplaceAll(string(weaponClassStr), string(0), "")))
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
