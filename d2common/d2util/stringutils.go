package d2util

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// AsterToEmpty converts strings beginning with "*" to "", for use when handling columns where an asterix can be used to comment out entries
func AsterToEmpty(text string) string {
	if strings.HasPrefix(text, "*") {
		return ""
	}

	return text
}

// EmptyToZero converts empty strings to "0" and leaves non-empty strings as is,
// for use before converting numerical data which equates empty to zero
func EmptyToZero(text string) string {
	if text == "" || text == " " {
		return "0"
	}

	return text
}

// StringToInt converts a string to an integer
func StringToInt(text string) int {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	return result
}

// SafeStringToInt converts a string to an integer, or returns -1 on falure

// StringToUint converts a string to a uint32
func StringToUint(text string) uint {
	result, err := strconv.ParseUint(text, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(result)
}

// StringToUint8 converts a string to an uint8
func StringToUint8(text string) uint8 {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	if result < 0 || result > 255 {
		panic(fmt.Sprintf("value %d out of range of byte", result))
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
		panic(fmt.Sprintf("value %d out of range of a signed byte", result))
	}

	return int8(result)
}

// Utf16BytesToString converts a utf16 byte array to string
func Utf16BytesToString(b []byte) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("must have even length byte slice")
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

// SplitIntoLinesWithMaxWidth splits the given string into lines considering the given maxChars
func SplitIntoLinesWithMaxWidth(fullSentence string, maxChars int) []string {
	lines := make([]string, 0)
	line := ""
	totalLength := 0
	words := strings.Split(fullSentence, " ")

	if len(words[0]) > maxChars {
		// mostly happened within CJK characters (no whitespace)
		return splitCjkIntoChunks(fullSentence, maxChars)
	}

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

func splitCjkIntoChunks(str string, chars int) []string {
	chunks := make([]string, chars/len(str))
	i, count := 0, 0

	for j, ch := range str {
		if ch < unicode.MaxLatin1 {
			count++
		} else {
			// assume we're truncating CJK characters
			count += 2
		}

		if count >= chars {
			chunks = append(chunks, str[i:j])
			i, count = j, 0
		}
	}

	return append(chunks, str[i:])
}
