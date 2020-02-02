package d2datadict

import (
	"strings"

	dh "github.com/OpenDiablo2/OpenDiablo2/d2common"
)

func MapHeaders(line string) map[string]int {
	m := make(map[string]int)
	r := strings.Split(line, "\t")
	for index, header := range r {
		m[header] = index
	}
	return m
}

func MapLoadInt(r *[]string, mapping *map[string]int, field string) int {
	index, ok := (*mapping)[field]
	if ok {
		return dh.StringToInt(dh.EmptyToZero(dh.AsterToEmpty((*r)[index])))
	}
	return 0
}

func MapLoadString(r *[]string, mapping *map[string]int, field string) string {
	index, ok := (*mapping)[field]
	if ok {
		return dh.AsterToEmpty((*r)[index])
	}
	return ""
}

func MapLoadBool(r *[]string, mapping *map[string]int, field string) bool {
	return MapLoadInt(r, mapping, field) == 1
}

func MapLoadUint8(r *[]string, mapping *map[string]int, field string) uint8 {
	index, ok := (*mapping)[field]
	if ok {
		return dh.StringToUint8(dh.EmptyToZero(dh.AsterToEmpty((*r)[index])))
	}
	return 0
}
