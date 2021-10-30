package d2records

// ObjectDetails stores all of the ObjectDetailRecords
type ObjectDetails map[int]*ObjectDetailRecord

// ObjectDetailRecord represents the settings for one type of object from objects.txt
type ObjectDetailRecord struct {
	token            string
	Name             string
	Description      string
	LightDiameter    [8]int
	StartFrame       [8]int
	OrderFlag        [8]int
	Parm             [8]int
	FrameCount       [8]int
	FrameDelta       [8]int
	AutoMap          int
	id               int
	SpawnMax         int
	TrapProbability  int
	SizeX            int
	SizeY            int
	NTgtFX           int
	NTgtFY           int
	NTgtBX           int
	NTgtBY           int
	Orientation      int
	Trans            int
	XOffset          int
	YOffset          int
	TotalPieces      int
	SubClass         int
	XSpace           int
	YSpace           int
	NameOffset       int
	OperateRange     int
	ShrineFunction   int
	Act              int
	Damage           int
	Left             int
	Top              int
	Width            int
	Height           int
	OperateFn        int
	PopulateFn       int
	InitFn           int
	ClientFn         int
	Index            int
	CycleAnimation   [8]bool
	Selectable       [8]bool
	BlocksLight      [8]bool
	HasCollision     [8]bool
	HasAnimationMode [8]bool
	SelS             [8]bool
	SelLH            bool
	EnvEffect        bool
	IsDoor           bool
	BlockVisibility  bool
	PreOperate       bool
	Draw             bool
	SelHD            bool
	SelTR            bool
	SelLG            bool
	SelRA            bool
	SelLA            bool
	SelRH            bool
	IsAttackable     bool
	SelSH            bool
	MonsterOk        bool
	Restore          bool
	Lockable         bool
	Gore             bool
	Sync             bool
	Flicker          bool
	Beta             bool
	Overlay          bool
	CollisionSubst   bool
	RestoreVirgins   bool
	BlockMissile     bool
	DrawUnder        bool
	OpenWarp         bool
	LightRed         byte
	LightGreen       byte
	LightBlue        byte
}
