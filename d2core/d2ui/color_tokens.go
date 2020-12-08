package d2ui

import "fmt"

// ColorToken is a string which is used inside of label strings to set font color.
type ColorToken string

// Color token formatting and pattern matching utility strings
const (
	ColorTokenFmt   = `%s%s`
	ColorTokenMatch = `\[[^\]]+\]` // nolint:gosec // has nothing to to with credentials
	ColorStrMatch   = ColorTokenMatch + `[^\[]+`
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
	ColorGrey100Alpha   = 0x69_69_69_ff
	ColorWhite100Alpha  = 0xff_ff_ff_ff
	ColorBlue100Alpha   = 0x69_69_ff_ff
	ColorYellow100Alpha = 0xff_ff_64_ff
	ColorGreen100Alpha  = 0x00_ff_00_ff
	ColorGold100Alpha   = 0xc7_b3_77_ff
	ColorOrange100Alpha = 0xff_a8_00_ff
	ColorRed100Alpha    = 0xff_77_77_ff
	ColorBlack100Alpha  = 0x00_00_00_ff
)

// ColorTokenize formats the string with the given color token
func ColorTokenize(s string, t ColorToken) string {
	return fmt.Sprintf(ColorTokenFmt, t, s)
}
