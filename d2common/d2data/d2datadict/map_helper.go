package d2datadict

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

func mapHeaders(line string) map[string]int {
	m := make(map[string]int)
	r := strings.Split(line, "\t")

	for index, header := range r {
		m[header] = index
	}

	return m
}

func mapLoadInt(r *[]string, mapping map[string]int, field string) int {
	index, ok := (mapping)[field]
	if ok {
		return d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty((*r)[index])))
	}

	return 0
}

func mapLoadString(r *[]string, mapping map[string]int, field string) string {
	index, ok := (mapping)[field]
	if ok {
		return d2util.AsterToEmpty((*r)[index])
	}

	return ""
}

func mapLoadBool(r *[]string, mapping map[string]int, field string) bool {
	return mapLoadInt(r, mapping, field) == 1
}

func mapLoadUint8(r *[]string, mapping map[string]int, field string) uint8 {
	index, ok := (mapping)[field]
	if ok {
		return d2util.StringToUint8(d2util.EmptyToZero(d2util.AsterToEmpty((*r)[index])))
	}

	return 0
}
