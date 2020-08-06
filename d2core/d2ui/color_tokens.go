package d2ui

import "fmt"

// ColorToken is a string which is used inside of label strings to set font color.
type ColorToken string

const (
	colorTokenFmt   = `%s%s`
	colorTokenMatch = `\[[^\]]+\]` // nolint:gosec // has nothing to to with credentials
	colorStrMatch   = colorTokenMatch + `[^\[]+`
)

// Color tokens for colored labels
const (
	ColorTokenGrey   ColorToken = "[grey]"
	ColorTokenRed    ColorToken = "[red]"
	ColorTokenWhite  ColorToken = "[white]"
	ColorTokenBlue   ColorToken = "[blue]"
	ColorTokenYellow ColorToken = "[yellow]"
	ColorTokenGreen  ColorToken = "[green]"
	ColorTokenGold   ColorToken = "[gold]"
	ColorTokenOrange ColorToken = "[orange]"
	ColorTokenBlack  ColorToken = "[black]"
)

// Color tokens for specific use-cases
const (
	ColorTokenSocketedItem  = ColorTokenGrey
	ColorTokenNormalItem    = ColorTokenWhite
	ColorTokenMagicItem     = ColorTokenBlue
	ColorTokenRareItem      = ColorTokenYellow
	ColorTokenSetItem       = ColorTokenGreen
	ColorTokenUniqueItem    = ColorTokenGold
	ColorTokenCraftedItem   = ColorTokenOrange
	ColorTokenServer        = ColorTokenRed
	ColorTokenButton        = ColorTokenBlack
	ColorTokenCharacterName = ColorTokenGold
	ColorTokenCharacterDesc = ColorTokenWhite
	ColorTokenCharacterType = ColorTokenGreen
)

const (
	colorGrey100Alpha   = 0x69_69_69_ff
	colorWhite100Alpha  = 0xff_ff_ff_ff
	colorBlue100Alpha   = 0x69_69_ff_ff
	colorYellow100Alpha = 0xff_ff_64_ff
	colorGreen100Alpha  = 0x00_ff_00_ff
	colorGold100Alpha   = 0xc7_b3_77_ff
	colorOrange100Alpha = 0xff_a8_00_ff
	colorRed100Alpha    = 0xff_77_77_ff
	colorBlack100Alpha  = 0x00_00_00_ff
)

// ColorTokenize formats the string with the given color token
func ColorTokenize(s string, t ColorToken) string {
	return fmt.Sprintf(colorTokenFmt, t, s)
}
