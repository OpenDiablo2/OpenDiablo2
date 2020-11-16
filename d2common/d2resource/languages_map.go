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
