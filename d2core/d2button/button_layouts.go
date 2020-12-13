package d2button

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom/rectangle"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// ButtonLayout defines the type of buttons
type ButtonLayout struct {
	SpritePath    string
	PalettePath   string
	FontPath      string
	ClickableRect *rectangle.Rectangle
	XSegments     int
	YSegments     int
	BaseFrame     int
	DisabledFrame int
	DisabledColor uint32
	TextOffset    int
	FixedWidth    int
	FixedHeight   int
	LabelColor       uint32
	Toggleable       bool
	AllowFrameChange bool
	HasImage         bool
	Tooltip          int
	TooltipXOffset   int
	TooltipYOffset   int
}

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
func GetLayout(t ButtonType) ButtonLayout {
	layouts := GetLayouts()

	return layouts[t]
}

func GetLayouts() map[ButtonType]ButtonLayout {
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

	return map[ButtonType]ButtonLayout{
		ButtonTypeWide: {
			XSegments:        buttonWideSegmentsX,
			YSegments:        buttonWideSegmentsY,
			DisabledFrame:    buttonWideDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonWideTextOffset,
			SpritePath:       d2resource.WideButtonBlank,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.CancelButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.ShortButtonBlank,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.MediumButtonBlank,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.TallButtonBlank,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.CancelButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.RunButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.GoldCoinButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
		ButtonTypeLeftArrow: {
			XSegments:        buttonBuySellSegmentsX,
			YSegments:        buttonBuySellSegmentsY,
			DisabledFrame:    buttonBuySellDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.BuySellButton,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.QuestLogQDescrBtn,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.WPTabs,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.QuestLogDone,
			PalettePath:      d2resource.PaletteUnits,
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
			SpritePath:       d2resource.SkillsPanelAmazon,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MenuButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
			SpritePath:       d2resource.MinipanelButton,
			PalettePath:      d2resource.PaletteSky,
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
