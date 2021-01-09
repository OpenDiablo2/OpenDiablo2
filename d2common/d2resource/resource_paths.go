package d2resource

// Paths of the resources inside the mpq files.
const (
	// --- Language

	LocalLanguage      = "/data/local/use"
	LanguageFontToken  = "{LANG_FONT}" //nolint:gosec // this is just a format string
	LanguageTableToken = "{LANG}"

	// --- Screens ---

	LoadingScreen = "/data/global/ui/Loading/loadingscreen.dc6"

	// --- Main Menu ---

	TrademarkScreen       = "/data/global/ui/FrontEnd/trademarkscreenEXP.dc6"
	GameSelectScreen      = "/data/global/ui/FrontEnd/gameselectscreenEXP.dc6"
	TCPIPBackground       = "/data/global/ui/FrontEnd/TCPIPscreen.dc6"
	Diablo2LogoFireLeft   = "/data/global/ui/FrontEnd/D2logoFireLeft.DC6"
	Diablo2LogoFireRight  = "/data/global/ui/FrontEnd/D2logoFireRight.DC6"
	Diablo2LogoBlackLeft  = "/data/global/ui/FrontEnd/D2logoBlackLeft.DC6"
	Diablo2LogoBlackRight = "/data/global/ui/FrontEnd/D2logoBlackRight.DC6"

	// --- Credits ---

	CreditsBackground = "/data/global/ui/CharSelect/creditsbckgexpand.dc6"
	CreditsText       = "/data/local/ui/" + LanguageTableToken + "/ExpansionCredits.txt"

	// --- Cinematics ---

	CinematicsBackground = "/data/global/ui/FrontEnd/CinematicsSelectionEXP.dc6"

	// --- Video Paths ---

	Act1Intro = "/data/local/video/" + LanguageTableToken + "/d2intro640x292.bik"
	Act2Intro = "/data/local/video/" + LanguageTableToken + "/act02start640x292.bik"
	Act3Intro = "/data/local/video/" + LanguageTableToken + "/act03start640x292.bik"
	Act4Intro = "/data/local/video/" + LanguageTableToken + "/act04start640x292.bik"
	Act4Outro = "/data/local/video/" + LanguageTableToken + "/act04end640x292.bik"
	Act5Intro = "/data/local/video/" + LanguageTableToken + "/d2x_intro_640x292.bik"
	Act5Outro = "/data/local/video/" + LanguageTableToken + "/d2x_out_640x292.bik"

	// --- Character Select Screen ---

	CharacterSelectBackground = "/data/global/ui/FrontEnd/charactercreationscreenEXP.dc6"
	CharacterSelectCampfire   = "/data/global/ui/FrontEnd/fire.DC6"

	CharacterSelectBarbarianUnselected         = "/data/global/ui/FrontEnd/barbarian/banu1.DC6"
	CharacterSelectBarbarianUnselectedH        = "/data/global/ui/FrontEnd/barbarian/banu2.DC6"
	CharacterSelectBarbarianSelected           = "/data/global/ui/FrontEnd/barbarian/banu3.DC6"
	CharacterSelectBarbarianForwardWalk        = "/data/global/ui/FrontEnd/barbarian/bafw.DC6"
	CharacterSelectBarbarianForwardWalkOverlay = "/data/global/ui/FrontEnd/barbarian/BAFWs.DC6"
	CharacterSelectBarbarianBackWalk           = "/data/global/ui/FrontEnd/barbarian/babw.DC6"

	CharacterSelectSorceressUnselected         = "/data/global/ui/FrontEnd/sorceress/SONU1.DC6"
	CharacterSelectSorceressUnselectedH        = "/data/global/ui/FrontEnd/sorceress/SONU2.DC6"
	CharacterSelectSorceressSelected           = "/data/global/ui/FrontEnd/sorceress/SONU3.DC6"
	CharacterSelectSorceressSelectedOverlay    = "/data/global/ui/FrontEnd/sorceress/SONU3s.DC6"
	CharacterSelectSorceressForwardWalk        = "/data/global/ui/FrontEnd/sorceress/SOFW.DC6"
	CharacterSelectSorceressForwardWalkOverlay = "/data/global/ui/FrontEnd/sorceress/SOFWs.DC6"
	CharacterSelectSorceressBackWalk           = "/data/global/ui/FrontEnd/sorceress/SOBW.DC6"
	CharacterSelectSorceressBackWalkOverlay    = "/data/global/ui/FrontEnd/sorceress/SOBWs.DC6"

	CharacterSelectNecromancerUnselected         = "/data/global/ui/FrontEnd/necromancer/NENU1.DC6"
	CharacterSelectNecromancerUnselectedH        = "/data/global/ui/FrontEnd/necromancer/NENU2.DC6"
	CharacterSelectNecromancerSelected           = "/data/global/ui/FrontEnd/necromancer/NENU3.DC6"
	CharacterSelectNecromancerSelectedOverlay    = "/data/global/ui/FrontEnd/necromancer/NENU3s.DC6"
	CharacterSelectNecromancerForwardWalk        = "/data/global/ui/FrontEnd/necromancer/NEFW.DC6"
	CharacterSelectNecromancerForwardWalkOverlay = "/data/global/ui/FrontEnd/necromancer/NEFWs.DC6"
	CharacterSelectNecromancerBackWalk           = "/data/global/ui/FrontEnd/necromancer/NEBW.DC6"
	CharacterSelectNecromancerBackWalkOverlay    = "/data/global/ui/FrontEnd/necromancer/NEBWs.DC6"

	CharacterSelectPaladinUnselected         = "/data/global/ui/FrontEnd/paladin/PANU1.DC6"
	CharacterSelectPaladinUnselectedH        = "/data/global/ui/FrontEnd/paladin/PANU2.DC6"
	CharacterSelectPaladinSelected           = "/data/global/ui/FrontEnd/paladin/PANU3.DC6"
	CharacterSelectPaladinForwardWalk        = "/data/global/ui/FrontEnd/paladin/PAFW.DC6"
	CharacterSelectPaladinForwardWalkOverlay = "/data/global/ui/FrontEnd/paladin/PAFWs.DC6"
	CharacterSelectPaladinBackWalk           = "/data/global/ui/FrontEnd/paladin/PABW.DC6"

	CharacterSelectAmazonUnselected         = "/data/global/ui/FrontEnd/amazon/AMNU1.DC6"
	CharacterSelectAmazonUnselectedH        = "/data/global/ui/FrontEnd/amazon/AMNU2.DC6"
	CharacterSelectAmazonSelected           = "/data/global/ui/FrontEnd/amazon/AMNU3.DC6"
	CharacterSelectAmazonForwardWalk        = "/data/global/ui/FrontEnd/amazon/AMFW.DC6"
	CharacterSelectAmazonForwardWalkOverlay = "/data/global/ui/FrontEnd/amazon/AMFWs.DC6"
	CharacterSelectAmazonBackWalk           = "/data/global/ui/FrontEnd/amazon/AMBW.DC6"

	CharacterSelectAssassinUnselected  = "/data/global/ui/FrontEnd/assassin/ASNU1.DC6"
	CharacterSelectAssassinUnselectedH = "/data/global/ui/FrontEnd/assassin/ASNU2.DC6"
	CharacterSelectAssassinSelected    = "/data/global/ui/FrontEnd/assassin/ASNU3.DC6"
	CharacterSelectAssassinForwardWalk = "/data/global/ui/FrontEnd/assassin/ASFW.DC6"
	CharacterSelectAssassinBackWalk    = "/data/global/ui/FrontEnd/assassin/ASBW.DC6"

	CharacterSelectDruidUnselected  = "/data/global/ui/FrontEnd/druid/DZNU1.dc6"
	CharacterSelectDruidUnselectedH = "/data/global/ui/FrontEnd/druid/DZNU2.dc6"
	CharacterSelectDruidSelected    = "/data/global/ui/FrontEnd/druid/DZNU3.DC6"
	CharacterSelectDruidForwardWalk = "/data/global/ui/FrontEnd/druid/DZFW.DC6"
	CharacterSelectDruidBackWalk    = "/data/global/ui/FrontEnd/druid/DZBW.DC6"

	// -- Character Selection

	CharacterSelectionBackground = "/data/global/ui/CharSelect/characterselectscreenEXP.dc6"
	CharacterSelectionSelectBox  = "/data/global/ui/CharSelect/charselectbox.dc6"
	PopUpOkCancel                = "/data/global/ui/FrontEnd/PopUpOKCancel.dc6"

	// --- Game ---

	GamePanels          = "/data/global/ui/PANEL/800ctrlpnl7.dc6"
	GameGlobeOverlap    = "/data/global/ui/PANEL/overlap.DC6"
	HealthManaIndicator = "/data/global/ui/PANEL/hlthmana.DC6"
	AddSkillButton      = "/data/global/ui/PANEL/level.DC6"
	MoveGoldDialog      = "/data/global/ui/menu/dialogbackground.DC6"
	WPTabs              = "/data/global/ui/menu/expwaygatetabs.dc6"
	WPBg                = "/data/global/ui/menu/waygatebackground.dc6"
	WPIcons             = "/data/global/ui/menu/waygateicons.dc6"
	UpDownArrows        = "/data/global/ui/BIGMENU/numberarrows.dc6"

	// --- Escape Menu ---
	// main
	EscapeOptions      = "/data/local/ui/" + LanguageTableToken + "/options.dc6"
	EscapeExit         = "/data/local/ui/" + LanguageTableToken + "/exit.dc6"
	EscapeReturnToGame = "/data/local/ui/" + LanguageTableToken + "/returntogame.dc6"
	// options
	EscapeOptSoundOptions   = "/data/local/ui/" + LanguageTableToken + "/soundoptions.dc6"
	EscapeOptVideoOptions   = "/data/local/ui/" + LanguageTableToken + "/videoOptions.dc6"
	EscapeOptAutoMapOptions = "/data/local/ui/" + LanguageTableToken + "/automapOptions.dc6"
	EscapeOptCfgOptions     = "/data/local/ui/" + LanguageTableToken + "/cfgOptions.dc6"
	EscapeOptPrevious       = "/data/local/ui/" + LanguageTableToken + "/previous.dc6"

	// sound options
	EscapeSndOptSoundVolume = "/data/local/ui/" + LanguageTableToken + "/sound.dc6"
	EscapeSndOptMusicVolume = "/data/local/ui/" + LanguageTableToken + "/music.dc6"
	EscapeSndOpt3DBias      = "/data/local/ui/" + LanguageTableToken + "/3dbias.dc6"
	// EscapeSndOptHWAcceleration =
	// EscapeSndOptENVEffects =
	EscapeSndOptNPCSpeech             = "/data/local/ui/" + LanguageTableToken + "/npcspeech.dc6"
	EscapeSndOptNPCSpeechAudioAndText = "/data/local/ui/" + LanguageTableToken + "/audiotext.dc6"
	EscapeSndOptNPCSpeechAudioOnly    = "/data/local/ui/" + LanguageTableToken + "/audioonly.dc6"
	EscapeSndOptNPCSpeechTextOnly     = "/data/local/ui/" + LanguageTableToken + "/textonly.dc6"

	// video options
	EscapeVidOptRes          = "/data/local/ui/" + LanguageTableToken + "/resolution.dc6"
	EscapeVidOptLightQuality = "/data/local/ui/" + LanguageTableToken + "/lightquality.dc6"
	EscapeVidOptBlendShadow  = "/data/local/ui/" + LanguageTableToken + "/blendshadow.dc6"
	EscapeVidOptPerspective  = "/data/local/ui/" + LanguageTableToken + "/prespective.dc6"
	EscapeVidOptGamma        = "/data/local/ui/" + LanguageTableToken + "/gamma.dc6"
	EscapeVidOptContrast     = "/data/local/ui/" + LanguageTableToken + "/contrast.dc6"

	// auto map
	EscapeAutoMapOptSize   = "/data/local/ui/" + LanguageTableToken + "/automapmode.dc6"
	EscapeAutoMapOptFade   = "/data/local/ui/" + LanguageTableToken + "/automapfade.dc6"
	EscapeAutoMapOptCenter = "/data/local/ui/" + LanguageTableToken + "/automapcenter.dc6"
	EscapeAutoMapOptNames  = "/data/local/ui/" + LanguageTableToken + "/automappartynames.dc6"

	// automap size
	EscapeAutoMapOptFullScreen = "/data/local/ui/" + LanguageTableToken + "/full.dc6"
	EscapeAutoMapOptMiniMap    = "/data/local/ui/" + LanguageTableToken + "/mini.dc6"

	// resolutions
	EscapeVideoOptRes640x480 = "/data/local/ui/" + LanguageTableToken + "/640x480.dc6"
	EscapeVideoOptRes800x600 = "/data/local/ui/" + LanguageTableToken + "/800x800.dc6"

	EscapeOn            = "/data/local/ui/" + LanguageTableToken + "/smallon.dc6"
	EscapeOff           = "/data/local/ui/" + LanguageTableToken + "/smalloff.dc6"
	EscapeYes           = "/data/local/ui/" + LanguageTableToken + "/smallyes.dc6"
	EscapeNo            = "/data/local/ui/" + LanguageTableToken + "/smallno.dc6"
	EscapeSlideBar      = "/data/global/ui/widgets/optbarc.dc6"
	EscapeSlideBarSkull = "/data/global/ui/widgets/optskull.dc6"

	// --- Help Overlay ---

	// HelpBorder = "/data/global/ui/MENU/helpborder.DC6"
	HelpBorder       = "/data/global/ui/MENU/800helpborder.DC6"
	HelpYellowBullet = "/data/global/ui/MENU/helpyellowbullet.DC6"
	HelpWhiteBullet  = "/data/global/ui/MENU/helpwhitebullet.DC6"

	// Box pieces, used in all in game boxes like npc interaction menu on click,
	// the chat window and the key binding menu
	BoxPieces = "/data/global/ui/MENU/boxpieces.DC6"

	// TextSlider contains the pieces to build a scrollbar in the
	// menus, such as the one in the configure keys menu
	TextSlider = "/data/global/ui/MENU/textslid.DC6"

	// Issue #685 - used in the mini-panel
	GameSmallMenuButton = "/data/global/ui/PANEL/menubutton.DC6"
	SkillIcon           = "/data/global/ui/PANEL/Skillicon.DC6"

	// --- Quest Log---
	QuestLogBg              = "/data/global/ui/MENU/questbackground.dc6"
	QuestLogDone            = "/data/global/ui/MENU/questdone.dc6"
	QuestLogTabs            = "/data/global/ui/MENU/expquesttabs.dc6"
	QuestLogQDescrBtn       = "/data/global/ui/MENU/questlast.dc6"
	QuestLogSocket          = "/data/global/ui/MENU/questsockets.dc6"
	QuestLogAQuestAnimation = "/data/global/ui/MENU/a%dq%d.dc6"

	// --- Mouse Pointers ---

	CursorDefault = "/data/global/ui/CURSOR/ohand.DC6"

	// --- Fonts & Locale (strings) ---
	Font6                = "/data/local/FONT/" + LanguageFontToken + "/font6"
	Font8                = "/data/local/FONT/" + LanguageFontToken + "/font8"
	Font16               = "/data/local/FONT/" + LanguageFontToken + "/font16"
	Font24               = "/data/local/FONT/" + LanguageFontToken + "/font24"
	Font30               = "/data/local/FONT/" + LanguageFontToken + "/font30"
	Font42               = "/data/local/FONT/" + LanguageFontToken + "/font42"
	FontFormal12         = "/data/local/FONT/" + LanguageFontToken + "/fontformal12"
	FontFormal11         = "/data/local/FONT/" + LanguageFontToken + "/fontformal11"
	FontFormal10         = "/data/local/FONT/" + LanguageFontToken + "/fontformal10"
	FontExocet10         = "/data/local/FONT/" + LanguageFontToken + "/fontexocet10"
	FontExocet8          = "/data/local/FONT/" + LanguageFontToken + "/fontexocet8"
	FontSucker           = "/data/local/FONT/" + LanguageFontToken + "/ReallyTheLastSucker"
	FontRediculous       = "/data/local/FONT/" + LanguageFontToken + "/fontridiculous"
	ExpansionStringTable = "/data/local/lng/" + LanguageTableToken + "/expansionstring.tbl"
	StringTable          = "/data/local/lng/" + LanguageTableToken + "/string.tbl"
	PatchStringTable     = "/data/local/lng/" + LanguageTableToken + "/patchstring.tbl"

	// --- UI ---

	WideButtonBlank   = "/data/global/ui/FrontEnd/WideButtonBlank.dc6"
	MediumButtonBlank = "/data/global/ui/FrontEnd/MediumButtonBlank.dc6"
	CancelButton      = "/data/global/ui/FrontEnd/CancelButtonBlank.dc6"
	NarrowButtonBlank = "/data/global/ui/FrontEnd/NarrowButtonBlank.dc6"
	ShortButtonBlank  = "/data/global/ui/CharSelect/ShortButtonBlank.dc6"
	TextBox2          = "/data/global/ui/FrontEnd/textbox2.dc6"
	TallButtonBlank   = "/data/global/ui/CharSelect/TallButtonBlank.dc6"
	Checkbox          = "/data/global/ui/FrontEnd/clickbox.dc6"
	Scrollbar         = "/data/global/ui/PANEL/scrollbar.dc6"

	PopUpLarge     = "/data/global/ui/FrontEnd/PopUpLarge.dc6"
	PopUpLargest   = "/data/global/ui/FrontEnd/PopUpLargest.dc6"
	PopUpWide      = "/data/global/ui/FrontEnd/PopUpWide.dc6"
	PopUpOk        = "/data/global/ui/FrontEnd/PopUpOk.dc6"
	PopUpOk2       = "/data/global/ui/FrontEnd/PopUpOk.dc6"
	PopUpOkCancel2 = "/data/global/ui/FrontEnd/PopUpOkCancel2.dc6"
	PopUp340x224   = "/data/global/ui/FrontEnd/PopUp_340x224.dc6"

	// --- GAME UI ---

	PentSpin        = "/data/global/ui/CURSOR/pentspin.DC6"
	Minipanel       = "/data/global/ui/PANEL/minipanel.DC6"
	MinipanelSmall  = "/data/global/ui/PANEL/minipanel_s.dc6"
	MinipanelButton = "/data/global/ui/PANEL/minipanelbtn.DC6"

	Frame                   = "/data/global/ui/PANEL/800borderframe.dc6"
	InventoryCharacterPanel = "/data/global/ui/PANEL/invchar6.DC6"
	InventoryWeaponsTab     = "/data/global/ui/PANEL/invchar6Tab.DC6"
	SkillsPanelAmazon       = "/data/global/ui/SPELLS/skltree_a_back.DC6"
	SkillsPanelBarbarian    = "/data/global/ui/SPELLS/skltree_b_back.DC6"
	SkillsPanelDruid        = "/data/global/ui/SPELLS/skltree_d_back.DC6"
	SkillsPanelAssassin     = "/data/global/ui/SPELLS/skltree_i_back.DC6"
	SkillsPanelNecromancer  = "/data/global/ui/SPELLS/skltree_n_back.DC6"
	SkillsPanelPaladin      = "/data/global/ui/SPELLS/skltree_p_back.DC6"
	SkillsPanelSorcerer     = "/data/global/ui/SPELLS/skltree_s_back.DC6"

	GenericSkills     = "/data/global/ui/SPELLS/Skillicon.DC6"
	AmazonSkills      = "/data/global/ui/SPELLS/AmSkillicon.DC6"
	BarbarianSkills   = "/data/global/ui/SPELLS/BaSkillicon.DC6"
	DruidSkills       = "/data/global/ui/SPELLS/DrSkillicon.DC6"
	AssassinSkills    = "/data/global/ui/SPELLS/AsSkillicon.DC6"
	NecromancerSkills = "/data/global/ui/SPELLS/NeSkillicon.DC6"
	PaladinSkills     = "/data/global/ui/SPELLS/PaSkillicon.DC6"
	SorcererSkills    = "/data/global/ui/SPELLS/SoSkillicon.DC6"

	RunButton      = "/data/global/ui/PANEL/runbutton.dc6"
	MenuButton     = "/data/global/ui/PANEL/menubutton.DC6"
	GoldCoinButton = "/data/global/ui/panel/goldcoinbtn.dc6"
	BuySellButton  = "/data/global/ui/panel/buysellbtn.dc6"

	ArmorPlaceholder      = "/data/global/ui/PANEL/inv_armor.DC6"
	BeltPlaceholder       = "/data/global/ui/PANEL/inv_belt.DC6"
	BootsPlaceholder      = "/data/global/ui/PANEL/inv_boots.DC6"
	HelmGlovePlaceholder  = "/data/global/ui/PANEL/inv_helm_glove.DC6"
	RingAmuletPlaceholder = "/data/global/ui/PANEL/inv_ring_amulet.DC6"
	WeaponsPlaceholder    = "/data/global/ui/PANEL/inv_weapons.DC6"

	// --- Data ---

	LevelPreset        = "/data/global/excel/LvlPrest.txt"
	LevelType          = "/data/global/excel/LvlTypes.txt"
	ObjectType         = "/data/global/excel/objtype.txt"
	LevelWarp          = "/data/global/excel/LvlWarp.txt"
	LevelDetails       = "/data/global/excel/Levels.txt"
	LevelMaze          = "/data/global/excel/LvlMaze.txt"
	LevelSubstitutions = "/data/global/excel/LvlSub.txt"

	ObjectDetails         = "/data/global/excel/Objects.txt"
	ObjectMode            = "/data/global/excel/ObjMode.txt"
	SoundSettings         = "/data/global/excel/Sounds.txt"
	ItemStatCost          = "/data/global/excel/ItemStatCost.txt"
	ItemRatio             = "/data/global/excel/itemratio.txt"
	ItemTypes             = "/data/global/excel/ItemTypes.txt"
	QualityItems          = "/data/global/excel/qualityitems.txt"
	LowQualityItems       = "/data/global/excel/lowqualityitems.txt"
	Overlays              = "/data/global/excel/Overlay.txt"
	Runes                 = "/data/global/excel/runes.txt"
	Sets                  = "/data/global/excel/Sets.txt"
	SetItems              = "/data/global/excel/SetItems.txt"
	AutoMagic             = "/data/global/excel/automagic.txt"
	BodyLocations         = "/data/global/excel/bodylocs.txt"
	Events                = "/data/global/excel/events.txt"
	Properties            = "/data/global/excel/Properties.txt"
	Hireling              = "/data/global/excel/hireling.txt"
	HirelingDescription   = "/data/global/excel/HireDesc.txt"
	DifficultyLevels      = "/data/global/excel/difficultylevels.txt"
	AutoMap               = "/data/global/excel/AutoMap.txt"
	CubeRecipes           = "/data/global/excel/cubemain.txt"
	CubeModifier          = "/data/global/excel/CubeMod.txt"
	CubeType              = "/data/global/excel/CubeType.txt"
	Skills                = "/data/global/excel/skills.txt"
	SkillDesc             = "/data/global/excel/skilldesc.txt"
	SkillCalc             = "/data/global/excel/skillcalc.txt"
	MissileCalc           = "/data/global/excel/misscalc.txt"
	TreasureClass         = "/data/global/excel/TreasureClass.txt"
	TreasureClassEx       = "/data/global/excel/TreasureClassEx.txt"
	States                = "/data/global/excel/states.txt"
	SoundEnvirons         = "/data/global/excel/soundenviron.txt"
	Shrines               = "/data/global/excel/shrines.txt"
	MonProp               = "/data/global/excel/Monprop.txt"
	ElemType              = "/data/global/excel/ElemTypes.txt"
	PlrMode               = "/data/global/excel/PlrMode.txt"
	PetType               = "/data/global/excel/pettype.txt"
	NPC                   = "/data/global/excel/npc.txt"
	MonsterUniqueModifier = "/data/global/excel/monumod.txt"
	MonsterEquipment      = "/data/global/excel/monequip.txt"
	UniqueAppellation     = "/data/global/excel/UniqueAppellation.txt"
	MonsterLevel          = "/data/global/excel/monlvl.txt"
	MonsterSound          = "/data/global/excel/monsounds.txt"
	MonsterSequence       = "/data/global/excel/monseq.txt"
	PlayerClass           = "/data/global/excel/PlayerClass.txt"
	PlayerType            = "/data/global/excel/PlrType.txt"
	Composite             = "/data/global/excel/Composit.txt"
	HitClass              = "/data/global/excel/HitClass.txt"
	ObjectGroup           = "/data/global/excel/objgroup.txt"
	CompCode              = "/data/global/excel/compcode.txt"
	Belts                 = "/data/global/excel/belts.txt"
	Gamble                = "/data/global/excel/gamble.txt"
	Colors                = "/data/global/excel/colors.txt"
	StorePage             = "/data/global/excel/StorePage.txt"

	// --- Animations ---

	ObjectData          = "/data/global/objects"
	AnimationData       = "/data/global/animdata.d2"
	PlayerAnimationBase = "/data/global/CHARS"
	MissileData         = "/data/global/missiles"
	ItemGraphics        = "/data/global/items"

	// --- Inventory Data ---

	Inventory   = "/data/global/excel/inventory.txt"
	Weapons     = "/data/global/excel/weapons.txt"
	Armor       = "/data/global/excel/armor.txt"
	ArmorType   = "/data/global/excel/ArmType.txt"
	WeaponClass = "/data/global/excel/WeaponClass.txt"
	Books       = "/data/global/excel/books.txt"
	Misc        = "/data/global/excel/misc.txt"
	UniqueItems = "/data/global/excel/UniqueItems.txt"
	Gems        = "/data/global/excel/gems.txt"

	// --- Affixes ---

	MagicPrefix = "/data/global/excel/MagicPrefix.txt"
	MagicSuffix = "/data/global/excel/MagicSuffix.txt"
	RarePrefix  = "/data/global/excel/RarePrefix.txt" // these are for item names
	RareSuffix  = "/data/global/excel/RareSuffix.txt"

	// --- Monster Prefix/Suffixes (?) ---
	UniquePrefix = "/data/global/excel/UniquePrefix.txt"
	UniqueSuffix = "/data/global/excel/UniqueSuffix.txt"

	// --- Character Data ---

	Experience = "/data/global/excel/experience.txt"
	CharStats  = "/data/global/excel/charstats.txt"

	// --- Music ---

	BGMTitle                    = "/data/global/music/introedit.wav"
	BGMOptions                  = "/data/global/music/Common/options.wav"
	BGMAct1AndarielAction       = "/data/global/music/Act1/andarielaction.wav"
	BGMAct1BloodRavenResolution = "/data/global/music/Act1/bloodravenresolution.wav"
	BGMAct1Caves                = "/data/global/music/Act1/caves.wav"
	BGMAct1Crypt                = "/data/global/music/Act1/crypt.wav"
	BGMAct1DenOfEvilAction      = "/data/global/music/Act1/denofevilaction.wav"
	BGMAct1Monastery            = "/data/global/music/Act1/monastery.wav"
	BGMAct1Town1                = "/data/global/music/Act1/town1.wav"
	BGMAct1Tristram             = "/data/global/music/Act1/tristram.wav"
	BGMAct1Wild                 = "/data/global/music/Act1/wild.wav"
	BGMAct2Desert               = "/data/global/music/Act2/desert.wav"
	BGMAct2Harem                = "/data/global/music/Act2/harem.wav"
	BGMAct2HoradricAction       = "/data/global/music/Act2/horadricaction.wav"
	BGMAct2Lair                 = "/data/global/music/Act2/lair.wav"
	BGMAct2RadamentResolution   = "/data/global/music/Act2/radamentresolution.wav"
	BGMAct2Sanctuary            = "/data/global/music/Act2/sanctuary.wav"
	BGMAct2Sewer                = "/data/global/music/Act2/sewer.wav"
	BGMAct2TaintedSunAction     = "/data/global/music/Act2/taintedsunaction.wav"
	BGMAct2Tombs                = "/data/global/music/Act2/tombs.wav"
	BGMAct2Town2                = "/data/global/music/Act2/town2.wav"
	BGMAct2Valley               = "/data/global/music/Act2/valley.wav"
	BGMAct3Jungle               = "/data/global/music/Act3/jungle.wav"
	BGMAct3Kurast               = "/data/global/music/Act3/kurast.wav"
	BGMAct3KurastSewer          = "/data/global/music/Act3/kurastsewer.wav"
	BGMAct3MefDeathAction       = "/data/global/music/Act3/mefdeathaction.wav"
	BGMAct3OrbAction            = "/data/global/music/Act3/orbaction.wav"
	BGMAct3Spider               = "/data/global/music/Act3/spider.wav"
	BGMAct3Town3                = "/data/global/music/Act3/town3.wav"
	BGMAct4Diablo               = "/data/global/music/Act4/diablo.wav"
	BGMAct4DiabloAction         = "/data/global/music/Act4/diabloaction.wav"
	BGMAct4ForgeAction          = "/data/global/music/Act4/forgeaction.wav"
	BGMAct4IzualAction          = "/data/global/music/Act4/izualaction.wav"
	BGMAct4Mesa                 = "/data/global/music/Act4/mesa.wav"
	BGMAct4Town4                = "/data/global/music/Act4/town4.wav"
	BGMAct5Baal                 = "/data/global/music/Act5/baal.wav"
	BGMAct5Siege                = "/data/global/music/Act5/siege.wav"
	BGMAct5Shenk                = "/data/global/music/Act5/shenkmusic.wav"
	BGMAct5XTown                = "/data/global/music/Act5/xtown.wav"
	BGMAct5XTemple              = "/data/global/music/Act5/xtemple.wav"
	BGMAct5IceCaves             = "/data/global/music/Act5/icecaves.wav"
	BGMAct5Nihlathak            = "/data/global/music/Act5/nihlathakmusic.wav"

	// --- Sound Effects ---

	SFXCursorSelect        = "cursor_select"
	SFXButtonClick         = "cursor_button_click"
	SFXAmazonDeselect      = "cursor_amazon_deselect"
	SFXAmazonSelect        = "cursor_amazon_select"
	SFXAssassinDeselect    = "Cursor/intro/assassin deselect.wav"
	SFXAssassinSelect      = "Cursor/intro/assassin select.wav"
	SFXBarbarianDeselect   = "cursor_barbarian_deselect"
	SFXBarbarianSelect     = "cursor_barbarian_select"
	SFXDruidDeselect       = "Cursor/intro/druid deselect.wav"
	SFXDruidSelect         = "Cursor/intro/druid select.wav"
	SFXNecromancerDeselect = "cursor_necromancer_deselect"
	SFXNecromancerSelect   = "cursor_necromancer_select"
	SFXPaladinDeselect     = "cursor_paladin_deselect"
	SFXPaladinSelect       = "cursor_paladin_select"
	SFXSorceressDeselect   = "cursor_sorceress_deselect"
	SFXSorceressSelect     = "cursor_sorceress_select"

	// --- Enemy Data ---

	MonStats         = "/data/global/excel/monstats.txt"
	MonStats2        = "/data/global/excel/monstats2.txt"
	MonPreset        = "/data/global/excel/monpreset.txt"
	MonType          = "/data/global/excel/Montype.txt"
	SuperUniques     = "/data/global/excel/SuperUniques.txt"
	MonMode          = "/data/global/excel/monmode.txt"
	MonsterPlacement = "/data/global/excel/MonPlace.txt"
	MonsterAI        = "/data/global/excel/monai.txt"

	// --- Skill Data ---

	Missiles = "/data/global/excel/Missiles.txt"

	// --- Palettes ---

	PaletteAct1      = "/data/global/palette/act1/pal.dat"
	PaletteAct2      = "/data/global/palette/act2/pal.dat"
	PaletteAct3      = "/data/global/palette/act3/pal.dat"
	PaletteAct4      = "/data/global/palette/act4/pal.dat"
	PaletteAct5      = "/data/global/palette/act5/pal.dat"
	PaletteEndGame   = "/data/global/palette/endgame/pal.dat"
	PaletteEndGame2  = "/data/global/palette/endgame2/pal.dat"
	PaletteFechar    = "/data/global/palette/fechar/pal.dat"
	PaletteLoading   = "/data/global/palette/loading/pal.dat"
	PaletteMenu0     = "/data/global/palette/menu0/pal.dat"
	PaletteMenu1     = "/data/global/palette/menu1/pal.dat"
	PaletteMenu2     = "/data/global/palette/menu2/pal.dat"
	PaletteMenu3     = "/data/global/palette/menu3/pal.dat"
	PaletteMenu4     = "/data/global/palette/menu4/pal.dat"
	PaletteSky       = "/data/global/palette/sky/pal.dat"
	PaletteStatic    = "/data/global/palette/static/pal.dat"
	PaletteTrademark = "/data/global/palette/trademark/pal.dat"
	PaletteUnits     = "/data/global/palette/units/pal.dat"

	// --- Palette Transforms ---

	PaletteTransformAct1      = "/data/global/palette/act1/Pal.pl2"
	PaletteTransformAct2      = "/data/global/palette/act2/Pal.pl2"
	PaletteTransformAct3      = "/data/global/palette/act3/Pal.pl2"
	PaletteTransformAct4      = "/data/global/palette/act4/Pal.pl2"
	PaletteTransformAct5      = "/data/global/palette/act5/Pal.pl2"
	PaletteTransformEndGame   = "/data/global/palette/endgame/Pal.pl2"
	PaletteTransformEndGame2  = "/data/global/palette/endgame2/Pal.pl2"
	PaletteTransformFechar    = "/data/global/palette/fechar/Pal.pl2"
	PaletteTransformLoading   = "/data/global/palette/loading/Pal.pl2"
	PaletteTransformMenu0     = "/data/global/palette/menu0/Pal.pl2"
	PaletteTransformMenu1     = "/data/global/palette/menu1/Pal.pl2"
	PaletteTransformMenu2     = "/data/global/palette/menu2/Pal.pl2"
	PaletteTransformMenu3     = "/data/global/palette/menu3/Pal.pl2"
	PaletteTransformMenu4     = "/data/global/palette/menu4/Pal.pl2"
	PaletteTransformSky       = "/data/global/palette/sky/Pal.pl2"
	PaletteTransformTrademark = "/data/global/palette/trademark/Pal.pl2"
)
