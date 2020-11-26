package d2resource

func getLanguages() map[byte]string {
	return map[byte]string{
		0x00: "ENG", // (English)
		0x01: "ESP", // (Spanish)
		0x02: "DEU", // (German)
		0x03: "FRA", // (French)
		0x04: "POR", // (Portuguese)
		0x05: "ITA", // (Italian)
		0x06: "JPN", // (Japanese)
		0x07: "KOR", // (Korean)
		0x08: "SIN", //
		0x09: "CHI", // (Chinese)
		0x0A: "POL", // (Polish)
		0x0B: "RUS", // (Russian)
		0x0C: "ENG", // (English)
	}
}

// GetLanguageLiteral returns string representation of language code
func GetLanguageLiteral(code byte) string {
	languages := getLanguages()

	return languages[code]
}

// Source https://github.com/eezstreet/OpenD2/blob/065f6e466048482b28b9dbc6286908dc1e0d10f6/Shared/D2Shared.hpp#L36
func getCharsets() map[string]string {
	return map[string]string{
		"ENG": "LATIN",  // (English)
		"ESP": "LATIN",  // (Spanish)
		"DEU": "LATIN",  // (German)
		"FRA": "LATIN",  // (French)
		"POR": "LATIN",  // (Portuguese)
		"ITA": "LATIN",  // (Italian)
		"JPN": "JPN",    // (Japanese)
		"KOR": "KOR",    // (Korean)
		"SIN": "LATIN",  //
		"CHI": "CHI",    // (Chinese)
		"POL": "LATIN2", // (Polish)
		"RUS": "CYR",    // (Russian)
	}
}

// GetFontCharset returns string representation of font charset
func GetFontCharset(language string) string {
	charset := getCharsets()

	return charset[language]
}

// GetLabelModifier returns modifier for language
/* modifiers for labels (used in string tables)
modifier is something like that:
english table:       polish table:
key  | value         key  |  value
#1   | v1                 |
#4   | v2            #4   | v1
#5   | v3            #5   | v2
#8   | v4            #8   | v3
So, GetLabelModifier returns value of offset in locale languages table
*/
// some of values need to be set up. For now values with "checked" comment
// was tested and works fine in main menu.
func GetLabelModifier(language string) int {
	modifiers := map[string]int{
		"ENG": 0, // (English) // checked
		"ESP": 0, // (Spanish)
		"DEU": 0, // (German) // checked
		"FRA": 0, // (French)
		"POR": 0, // (Portuguese)
		"ITA": 0, // (Italian)
		"JPN": 0, // (Japanese)
		"KOR": 0, // (Korean)
		"SIN": 0, //
		"CHI": 0, // (Chinese)
		"POL": 1, // (Polish) // checked
		"RUS": 0, // (Russian)
	}

	return modifiers[language]
}
