package d2ui

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// ButtonType defines the type of button
type ButtonType int

// ButtonType constants
const (
	ButtonTypeWide     ButtonType = 1
	ButtonTypeMedium   ButtonType = 2
	ButtonTypeNarrow   ButtonType = 3
	ButtonTypeCancel   ButtonType = 4
	ButtonTypeTall     ButtonType = 5
	ButtonTypeShort    ButtonType = 6
	ButtonTypeOkCancel ButtonType = 7

	// Game UI

	ButtonTypeSkill              ButtonType = 7
	ButtonTypeRun                ButtonType = 8
	ButtonTypeMenu               ButtonType = 9
	ButtonTypeGoldCoin           ButtonType = 10
	ButtonTypeClose              ButtonType = 11
	ButtonTypeSecondaryInvHand   ButtonType = 12
	ButtonTypeMinipanelCharacter ButtonType = 13
	ButtonTypeMinipanelInventory ButtonType = 14
	ButtonTypeMinipanelSkill     ButtonType = 15
	ButtonTypeMinipanelAutomap   ButtonType = 16
	ButtonTypeMinipanelMessage   ButtonType = 17
	ButtonTypeMinipanelQuest     ButtonType = 18
	ButtonTypeMinipanelMen       ButtonType = 19
	ButtonTypeSquareClose        ButtonType = 20
	ButtonTypeSquareOk           ButtonType = 21
	ButtonTypeSkillTreeTab       ButtonType = 22
	ButtonTypeQuestDescr         ButtonType = 23
	ButtonTypeMinipanelOpenClose ButtonType = 24
	ButtonTypeMinipanelParty     ButtonType = 25
	ButtonTypeBuy                ButtonType = 26
	ButtonTypeSell               ButtonType = 27
	ButtonTypeRepair             ButtonType = 28
	ButtonTypeRepairAll          ButtonType = 29
	ButtonTypeUpArrow            ButtonType = 30
	ButtonTypeDownArrow          ButtonType = 31
	ButtonTypeLeftArrow          ButtonType = 32
	ButtonTypeRightArrow         ButtonType = 33
	ButtonTypeQuery              ButtonType = 34
	ButtonTypeSquelchChat        ButtonType = 35
	ButtonTypeTabBlank           ButtonType = 36
	ButtonTypeBlankQuestBtn      ButtonType = 37

	ButtonNoFixedWidth  int = -1
	ButtonNoFixedHeight int = -1
)

const (
	buttonStatePressed = iota + 1
	buttonStateToggled
	buttonStatePressedToggled
)

const (
	buyButtonBaseFrame         = 2  // base frame offset of the "buy" button dc6
	sellButtonBaseFrame        = 4  // base frame offset of the "sell" button dc6
	repairButtonBaseFrame      = 6  // base frame offset of the "repair" button dc6
	queryButtonBaseFrame       = 8  // base frame offset of the "query" button dc6
	closeButtonBaseFrame       = 10 // base frame offset of the "close" button dc6
	leftArrowButtonBaseFrame   = 12 // base frame offset of the "leftArrow" button dc6
	rightArrowButtonBaseFrame  = 14 // base frame offset of the "rightArrow" button dc6
	okButtonBaseFrame          = 16 // base frame offset of the "ok" button dc6
	repairAllButtonBaseFrame   = 18 // base frame offset of the "repair all" button dc6
	squelchChatButtonBaseFrame = 20 // base frame offset of the "?" button dc6
)

const (
	greyAlpha100     = 0x646464ff
	lightGreyAlpha75 = 0x808080c3
	whiteAlpha100    = 0xffffffff
)

// ButtonLayout defines the type of buttons
type ButtonLayout struct {
	ResourceName     string
	PaletteName      string
	FontPath         string
	ClickableRect    *image.Rectangle
	XSegments        int
	YSegments        int
	BaseFrame        int
	DisabledFrame    int
	DisabledColor    uint32
	TextOffset       int
	FixedWidth       int
	FixedHeight      int
	LabelColor       uint32
	Toggleable       bool
	AllowFrameChange bool
	HasImage         bool
	Tooltip          int
	TooltipXOffset   int
	TooltipYOffset   int
}

const (
	buttonTooltipNone int = iota
	buttonTooltipClose
	buttonTooltipOk
	buttonTooltipBuy
	buttonTooltipSell
	buttonTooltipRepair
	buttonTooltipRepairAll
	buttonTooltipLeftArrow
	buttonTooltipRightArrow
	buttonTooltipQuery
	buttonTooltipSquelchChat
)

const (
	buttonBuySellTooltipXOffset = 15
	buttonBuySellTooltipYOffset = -2
)

const (
	buttonWideSegmentsX     = 2
	buttonWideSegmentsY     = 1
	buttonWideDisabledFrame = -1
	buttonWideTextOffset    = 1

	buttonShortSegmentsX     = 1
	buttonShortSegmentsY     = 1
	buttonShortDisabledFrame = -1
	buttonShortTextOffset    = -1

	buttonMediumSegmentsX = 1
	buttonMediumSegmentsY = 1

	buttonTallSegmentsX  = 1
	buttonTallSegmentsY  = 1
	buttonTallTextOffset = 5

	buttonCancelSegmentsX  = 1
	buttonCancelSegmentsY  = 1
	buttonCancelTextOffset = 1

	buttonOkCancelSegmentsX     = 1
	buttonOkCancelSegmentsY     = 1
	buttonOkCancelDisabledFrame = -1

	buttonUpDownArrowSegmentsX     = 1
	buttonUpDownArrowSegmentsY     = 1
	buttonUpDownArrowDisabledFrame = -1
	buttonUpArrowBaseFrame         = 0
	buttonDownArrowBaseFrame       = 2

	buttonBuySellSegmentsX     = 1
	buttonBuySellSegmentsY     = 1
	buttonBuySellDisabledFrame = 1

	buttonSkillTreeTabXSegments = 1
	buttonSkillTreeTabYSegments = 1

	buttonSkillTreeTabDisabledFrame = 7
	buttonSkillTreeTabBaseFrame     = 7
	buttonSkillTreeTabFixedWidth    = 93
	buttonSkillTreeTabFixedHeight   = 107

	buttonTabXSegments = 1
	buttonTabYSegments = 1

	buttonMinipanelOpenCloseBaseFrame = 0
	buttonMinipanelXSegments          = 1
	buttonMinipanelYSegments          = 1

	blankQuestButtonXSegments      = 1
	blankQuestButtonYSegments      = 1
	blankQuestButtonDisabledFrames = -1

	buttonMinipanelCharacterBaseFrame = 0
	buttonMinipanelInventoryBaseFrame = 2
	buttonMinipanelSkilltreeBaseFrame = 4
	buttonMinipanelPartyBaseFrame     = 6
	buttonMinipanelAutomapBaseFrame   = 8
	buttonMinipanelMessageBaseFrame   = 10
	buttonMinipanelQuestBaseFrame     = 12
	buttonMinipanelMenBaseFrame       = 14

	buttonRunSegmentsX     = 1
	buttonRunSegmentsY     = 1
	buttonRunDisabledFrame = -1

	buttonGoldCoinSegmentsX     = 1
	buttonGoldCoinSegmentsY     = 1
	buttonGoldCoinDisabledFrame = -1

	pressedButtonOffset = 2
)

// nolint:funlen // cant reduce
func getButtonLayouts() map[ButtonType]ButtonLayout {
	return map[ButtonType]ButtonLayout{
		ButtonTypeWide: {
			XSegments:        buttonWideSegmentsX,
			YSegments:        buttonWideSegmentsY,
			DisabledFrame:    buttonWideDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonWideTextOffset,
			ResourceName:     d2resource.WideButtonBlank,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeCancel: {
			XSegments:        buttonCancelSegmentsX,
			YSegments:        buttonCancelSegmentsY,
			DisabledFrame:    0,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonCancelTextOffset,
			ResourceName:     d2resource.CancelButton,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeShort: {
			XSegments:        buttonShortSegmentsX,
			YSegments:        buttonShortSegmentsY,
			DisabledFrame:    buttonShortDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonShortTextOffset,
			ResourceName:     d2resource.ShortButtonBlank,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeMedium: {
			XSegments:        buttonMediumSegmentsX,
			YSegments:        buttonMediumSegmentsY,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.MediumButtonBlank,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeTall: {
			XSegments:        buttonTallSegmentsX,
			YSegments:        buttonTallSegmentsY,
			TextOffset:       buttonTallTextOffset,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.TallButtonBlank,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeOkCancel: {
			XSegments:        buttonOkCancelSegmentsX,
			YSegments:        buttonOkCancelSegmentsY,
			DisabledFrame:    buttonOkCancelDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.CancelButton,
			PaletteName:      d2resource.PaletteUnits,
			FontPath:         d2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeRun: {
			XSegments:        buttonRunSegmentsX,
			YSegments:        buttonRunSegmentsY,
			DisabledFrame:    buttonRunDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.RunButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       true,
			FontPath:         d2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeGoldCoin: {
			XSegments:        buttonGoldCoinSegmentsX,
			YSegments:        buttonGoldCoinSegmentsY,
			DisabledFrame:    buttonGoldCoinDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.GoldCoinButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       true,
			FontPath:         d2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeSquareClose: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        closeButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipClose,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeSquareOk: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        okButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipOk,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeBuy: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        buyButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipBuy,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeSell: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        sellButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipSell,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeRepair: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        repairButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipRepair,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeRepairAll: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        repairAllButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipRepairAll,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeUpArrow: {
			XSegments:        buttonUpDownArrowSegmentsX,
			YSegments:        buttonUpDownArrowSegmentsY,
			DisabledFrame:    buttonUpDownArrowDisabledFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonUpArrowBaseFrame,
			ResourceName:     d2resource.UpDownArrows,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
		},
		ButtonTypeDownArrow: {
			XSegments:        buttonUpDownArrowSegmentsX,
			YSegments:        buttonUpDownArrowSegmentsY,
			DisabledFrame:    buttonUpDownArrowDisabledFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonDownArrowBaseFrame,
			ResourceName:     d2resource.UpDownArrows,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
		},
		ButtonTypeLeftArrow: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        leftArrowButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipLeftArrow,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeRightArrow: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        rightArrowButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipRightArrow,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeQuery: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        queryButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipQuery,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeSquelchChat: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.BuySellButton,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			BaseFrame:        squelchChatButtonBaseFrame,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
			Tooltip:          buttonTooltipSquelchChat,
			TooltipXOffset:   buttonBuySellTooltipXOffset,
			TooltipYOffset:   buttonBuySellTooltipYOffset,
		},
		ButtonTypeQuestDescr: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.QuestLogQDescrBtn,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeTabBlank: {
			XSegments:        buttonTabXSegments,
			YSegments:        buttonTabYSegments,
			DisabledFrame:    0,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.WPTabs,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: false,
			HasImage:         false,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeBlankQuestBtn: {
			XSegments:        blankQuestButtonXSegments,
			YSegments:        blankQuestButtonYSegments,
			DisabledFrame:    blankQuestButtonDisabledFrames,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     d2resource.QuestLogDone,
			PaletteName:      d2resource.PaletteUnits,
			Toggleable:       true,
			FontPath:         d2resource.Font30,
			AllowFrameChange: false,
			HasImage:         false,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeSkillTreeTab: {
			XSegments:        buttonSkillTreeTabXSegments,
			YSegments:        buttonSkillTreeTabYSegments,
			DisabledFrame:    buttonSkillTreeTabDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			BaseFrame:        buttonSkillTreeTabBaseFrame,
			ResourceName:     d2resource.SkillsPanelAmazon,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: false,
			HasImage:         false,
			FixedWidth:       buttonSkillTreeTabFixedWidth,
			FixedHeight:      buttonSkillTreeTabFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelOpenClose: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelOpenCloseBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelOpenCloseBaseFrame,
			ResourceName:     d2resource.MenuButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       true,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelCharacter: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelCharacterBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelCharacterBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelInventory: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelInventoryBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelInventoryBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelSkill: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelSkilltreeBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelSkilltreeBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelParty: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelPartyBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelPartyBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelAutomap: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelAutomapBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelAutomapBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelMessage: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelMessageBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelMessageBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelQuest: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelQuestBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelQuestBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
		ButtonTypeMinipanelMen: {
			XSegments:        buttonMinipanelXSegments,
			YSegments:        buttonMinipanelYSegments,
			DisabledFrame:    buttonMinipanelMenBaseFrame,
			DisabledColor:    whiteAlpha100,
			BaseFrame:        buttonMinipanelMenBaseFrame,
			ResourceName:     d2resource.MinipanelButton,
			PaletteName:      d2resource.PaletteSky,
			Toggleable:       false,
			FontPath:         d2resource.Font16,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       whiteAlpha100,
		},
	}
}

// GetButtonLayout returns a button layout for the given button type
func GetButtonLayout(t ButtonType) ButtonLayout {
	return getButtonLayouts()[t]
}

var _ Widget = &Button{} // static check to ensure button implements widget

// Button defines a standard wide UI button
type Button struct {
	*BaseWidget
	buttonLayout          ButtonLayout
	normalSurface         d2interface.Surface
	pressedSurface        d2interface.Surface
	toggledSurface        d2interface.Surface
	pressedToggledSurface d2interface.Surface
	disabledSurface       d2interface.Surface
	onClick               func()
	enabled               bool
	pressed               bool
	toggled               bool
	tooltip               *Tooltip
}

// NewButton creates an instance of Button
func (ui *UIManager) NewButton(buttonType ButtonType, text string) *Button {
	base := NewBaseWidget(ui)
	base.SetVisible(true)

	btn := &Button{
		BaseWidget: base,
		enabled:    true,
		pressed:    false,
	}

	buttonLayout := getButtonLayouts()[buttonType]
	btn.buttonLayout = buttonLayout
	lbl := ui.NewLabel(buttonLayout.FontPath, d2resource.PaletteUnits)

	lbl.SetText(text)
	lbl.Color[0] = d2util.Color(buttonLayout.LabelColor)
	lbl.Alignment = HorizontalAlignCenter

	buttonSprite, err := ui.NewSprite(buttonLayout.ResourceName, buttonLayout.PaletteName)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	if buttonLayout.FixedWidth > 0 {
		btn.width = buttonLayout.FixedWidth
	} else {
		for i := 0; i < buttonLayout.XSegments; i++ {
			w, _, frameSizeErr := buttonSprite.GetFrameSize(i)
			if frameSizeErr != nil {
				ui.Error(frameSizeErr.Error())
				return nil
			}

			btn.width += w
		}
	}

	if buttonLayout.FixedHeight > 0 {
		btn.height = buttonLayout.FixedHeight
	} else {
		for i := 0; i < buttonLayout.YSegments; i++ {
			_, h, frameSizeErr := buttonSprite.GetFrameSize(i * buttonLayout.YSegments)
			if frameSizeErr != nil {
				ui.Error(frameSizeErr.Error())
				return nil
			}

			btn.height += h
		}
	}

	btn.normalSurface = ui.renderer.NewSurface(btn.width, btn.height)

	buttonSprite.SetPosition(0, 0)
	buttonSprite.SetEffect(d2enum.DrawEffectModulate)

	btn.createTooltip()

	ui.addWidget(btn) // important that this comes before prerenderStates!

	btn.prerenderStates(buttonSprite, &buttonLayout, lbl)

	return btn
}

type buttonStateDescriptor struct {
	baseFrame            int
	offsetX, offsetY     int
	prerenderdestination *d2interface.Surface
	fmtErr               string
}

func (v *Button) createTooltip() {
	var t *Tooltip
	// this is also related with https://github.com/OpenDiablo2/OpenDiablo2/issues/944
	// all strings starting with "#" could be wrong translated to another locales
	switch v.buttonLayout.Tooltip {
	case buttonTooltipNone:
		return
	case buttonTooltipClose:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("strClose"))
	case buttonTooltipOk:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateLabel(d2enum.OKLabel))
	case buttonTooltipBuy:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCPurchaseItems"))
	case buttonTooltipSell:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCSellItems"))
	case buttonTooltipRepair:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCRepairItems"))
	case buttonTooltipRepairAll:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("#128"))
	case buttonTooltipLeftArrow:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("KeyLeft"))
	case buttonTooltipRightArrow:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("KeyRight"))
	case buttonTooltipQuery:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("")) // need to be set up
	case buttonTooltipSquelchChat:
		t = v.manager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("strParty19")) // need to be verivied
	}

	t.SetVisible(false)
	v.SetTooltip(t)
}

func (v *Button) prerenderStates(btnSprite *Sprite, btnLayout *ButtonLayout, label *Label) {
	numButtonStates := btnSprite.GetFrameCount() / (btnLayout.XSegments * btnLayout.YSegments)

	// buttons always have a base image
	if v.buttonLayout.HasImage {
		btnSprite.RenderSegmented(v.normalSurface, btnLayout.XSegments,
			btnLayout.YSegments, btnLayout.BaseFrame)
	}

	_, labelHeight := label.GetSize()
	textY := half(v.height - labelHeight)
	xOffset := half(v.width)

	label.SetPosition(xOffset, textY)
	label.Render(v.normalSurface)

	if !btnLayout.HasImage || !btnLayout.AllowFrameChange {
		return
	}

	xSeg, ySeg, baseFrame := btnLayout.XSegments, btnLayout.YSegments, btnLayout.BaseFrame

	buttonStateConfigs := make([]*buttonStateDescriptor, 0)

	// pressed button
	if numButtonStates > buttonStatePressed {
		state := &buttonStateDescriptor{
			baseFrame + buttonStatePressed,
			xOffset - pressedButtonOffset, textY + pressedButtonOffset,
			&v.pressedSurface,
			"failed to render button pressedSurface, err: %v\n",
		}

		buttonStateConfigs = append(buttonStateConfigs, state)
	}

	// toggle button
	if numButtonStates > buttonStateToggled {
		buttonStateConfigs = append(buttonStateConfigs, &buttonStateDescriptor{
			baseFrame + buttonStateToggled,
			xOffset, textY,
			&v.toggledSurface,
			"failed to render button toggledSurface, err: %v\n",
		})
	}

	// pressed+toggled
	if numButtonStates > buttonStatePressedToggled {
		buttonStateConfigs = append(buttonStateConfigs, &buttonStateDescriptor{
			baseFrame + buttonStatePressedToggled,
			xOffset, textY,
			&v.pressedToggledSurface,
			"failed to render button pressedToggledSurface, err: %v\n",
		})
	}

	// disabled button
	if btnLayout.DisabledFrame != -1 {
		disabledState := &buttonStateDescriptor{
			btnLayout.DisabledFrame,
			xOffset, textY,
			&v.disabledSurface,
			"failed to render button disabledSurface, err: %v\n",
		}

		buttonStateConfigs = append(buttonStateConfigs, disabledState)
	}

	for stateIdx, w, h := 0, v.width, v.height; stateIdx < len(buttonStateConfigs); stateIdx++ {
		state := buttonStateConfigs[stateIdx]

		if stateIdx > 1 && btnLayout.ResourceName == d2resource.BuySellButton {
			// Without returning early, the button UI gets all subsequent (unrelated) frames
			// stacked on top. Only 2 frames from this sprite are applicable to the button
			// in question. The presentation is incorrect without this hack!
			continue
		}

		surface := v.manager.renderer.NewSurface(w, h)

		*state.prerenderdestination = surface

		btnSprite.RenderSegmented(*state.prerenderdestination, xSeg, ySeg, state.baseFrame)

		label.SetPosition(state.offsetX, state.offsetY)
		label.Render(*state.prerenderdestination)
	}
}

// OnActivated defines the callback handler for the activate event
func (v *Button) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate calls the on activated callback handler, if any
func (v *Button) Activate() {
	if v.onClick == nil {
		return
	}

	v.onClick()
}

// Render renders the button
func (v *Button) Render(target d2interface.Surface) {
	target.PushFilter(d2enum.FilterNearest)
	defer target.Pop()

	target.PushTranslation(v.x, v.y)
	defer target.Pop()

	switch {
	case !v.enabled:
		target.PushColor(d2util.Color(v.buttonLayout.DisabledColor))
		defer target.Pop()

		if v.toggled {
			target.Render(v.toggledSurface)
		} else {
			target.Render(v.disabledSurface)
		}
	case v.toggled && v.pressed:
		target.Render(v.pressedToggledSurface)
	case v.pressed:
		if v.buttonLayout.AllowFrameChange {
			target.Render(v.pressedSurface)
		} else {
			target.Render(v.normalSurface)
		}
	case v.toggled:
		target.Render(v.toggledSurface)
	default:
		target.Render(v.normalSurface)
	}
}

// Toggle negates the toggled state of the button
func (v *Button) Toggle() {
	v.toggled = !v.toggled
}

// GetToggled returns the toggled state of the button
func (v *Button) GetToggled() bool {
	return v.toggled
}

// Advance advances the button state
func (v *Button) Advance(_ float64) error {
	return nil
}

// GetEnabled returns the enabled state
func (v *Button) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Button) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetPressed returns the pressed state of the button
func (v *Button) GetPressed() bool {
	return v.pressed
}

// SetPressed sets the pressed state of the button
func (v *Button) SetPressed(pressed bool) {
	v.pressed = pressed
}

// SetVisible sets the pressed state of the button
func (v *Button) SetVisible(visible bool) {
	v.BaseWidget.SetVisible(visible)

	if v.isHovered() && !visible {
		v.hoverEnd()
	}
}

// SetPosition sets the position of the widget
func (v *Button) SetPosition(x, y int) {
	v.BaseWidget.SetPosition(x, y)

	if v.buttonLayout.Tooltip != buttonTooltipNone {
		v.tooltip.SetPosition(x+v.buttonLayout.TooltipXOffset, y+v.buttonLayout.TooltipYOffset)
	}
}

// SetTooltip adds a tooltip to the button
func (v *Button) SetTooltip(t *Tooltip) {
	v.tooltip = t
	v.OnHoverStart(func() { v.tooltip.SetVisible(true) })
	v.OnHoverEnd(func() { v.tooltip.SetVisible(false) })
}

func half(n int) int {
	return n / 2
}
