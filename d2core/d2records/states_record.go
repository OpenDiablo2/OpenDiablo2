package d2records

// States contains the state records
type States map[string]*StateRecord

// StateRecord describes a body location that items can be equipped to
type StateRecord struct {
	CltEvent      string
	Overlay1      string
	Overlay2      string
	Overlay3      string
	Overlay4      string
	PgOverlay     string
	CastOverlay   string
	RemOverlay    string
	Stat          string
	Missile       string
	Skill         string
	ItemType      string
	ItemTrans     string
	OnSound       string
	OffSound      string
	State         string
	SetFunc       int
	RemFunc       int
	ColorPri      int
	ColorShift    int
	LightR        int
	LightG        int
	LightB        int
	GfxType       int
	GfxClass      int
	Group         int
	CltEventFunc  int
	CltActiveFunc int
	SrvActiveFunc int
	Curse         bool
	NoSend        bool
	Transform     bool
	Aura          bool
	Cureable      bool
	RemHit        bool
	Active        bool
	Restrict      bool
	Disguise      bool
	Blue          bool
	AttBlue       bool
	DmgBlue       bool
	ArmBlue       bool
	RfBlue        bool
	RlBlue        bool
	RcBlue        bool
	RpBlue        bool
	AttRed        bool
	DmgRed        bool
	ArmRed        bool
	RfRed         bool
	RlRed         bool
	RcRed         bool
	RpRed         bool
	StamBarBlue   bool
	Exp           bool
	PlrStayDeath  bool
	MonStayDeath  bool
	BossStayDeath bool
	Hide          bool
	Shatter       bool
	UDead         bool
	Life          bool
	Green         bool
	Pgsv          bool
	NoOverlays    bool
	NoClear       bool
	BossInv       bool
	MeleeOnly     bool
	NotOnDead     bool
}
