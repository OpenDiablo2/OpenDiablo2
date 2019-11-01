package Common

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// StringToInt converts a string to an integer
func StringToInt(text string) int {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return result
}

// StringToUint8 converts a string to an uint8
func StringToUint8(text string) uint8 {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	if result < 0 || result > 255 {
		panic("value out of range of byte")
	}
	return uint8(result)
}

// StringToInt8 converts a string to an int8
func StringToInt8(text string) int8 {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	if result < -128 || result > 122 {
		panic("value out of range of a signed byte")
	}
	return int8(result)
}

func Utf16BytesToString(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

func CombineStrings(input []string) string {
	return strings.Join(input, "\n")
}

func SplitIntoLinesWithMaxWidth(fullSentence string, maxChars int) []string {
	lines := make([]string, 0)
	line := ""
	totalLength := 0
	words := strings.Split(fullSentence, " ")
	for _, word := range words {
		totalLength += 1 + len(word)
		if totalLength > maxChars {
			totalLength = len(word)
			lines = append(lines, line)
			line = ""
		} else {
			line += " "
		}
		line += word
	}

	if len(line) > 0 {
		lines = append(lines, line)
	}

	return lines
}
