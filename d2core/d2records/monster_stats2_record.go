package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// MonStats2 stores all of the MonStat2Records
type MonStats2 map[string]*MonStat2Record

// MonStat2Record is a representation of a row from monstats2.txt
type MonStat2Record struct {
	Key                string
	BaseWeaponClass    string
	ResurrectSkill     string
	Heart              string
	BodyPart           string
	EquipmentOptions   [16][]string
	DirectionsPerMode  [16]int
	ResurrectMode      d2enum.MonsterAnimationMode
	OverlayHeight      int
	SizeX              int
	SizeY              int
	BoxTop             int
	BoxLeft            int
	BoxWidth           int
	BoxHeight          int
	SpawnMethod        int
	MeleeRng           int
	HitClass           int
	TotalPieces        int
	Height             int
	Restore            int
	AutomapCel         int
	LocalBlood         int
	Bleed              int
	Light              int
	LightR             int
	LightG             int
	lightB             int
	NormalPalette      int
	NightmarePalette   int
	HellPalatte        int
	InfernoLen         int
	InfernoAnim        int
	InfernoRollback    int
	PixelHeight        int
	HasAnimationMode   [16]bool
	HasComponent       [16]bool
	NoGfxHitTest       bool
	A1mv               bool
	A2mv               bool
	SCmv               bool
	S1mv               bool
	S2mv               bool
	S3mv               bool
	S4mv               bool
	NoMap              bool
	NoOvly             bool
	IsSelectable       bool
	AllySelectable     bool
	NotSelectable      bool
	shiftSel           bool
	IsCorpseSelectable bool
	IsAttackable       bool
	IsRevivable        bool
	IsCritter          bool
	IsSmall            bool
	IsLarge            bool
	IsSoft             bool
	IsInert            bool
	objCol             bool
	IsCorpseCollidable bool
	IsCorpseWalkable   bool
	HasShadow          bool
	NoUniqueShift      bool
	CompositeDeath     bool
}
