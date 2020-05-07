package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"log"
)

var MagicPrefixDictionary *d2common.DataDictionary
var MagicSuffixDictionary *d2common.DataDictionary

var RarePrefixDictionary *d2common.DataDictionary
var RareSuffixDictionary *d2common.DataDictionary

var UniquePrefixDictionary *d2common.DataDictionary
var UniqueSuffixDictionary *d2common.DataDictionary

func loadAffixes(file []byte, dst *d2common.DataDictionary, name string) {
	dst = d2common.LoadDataDictionary(string(file))
	log.Printf("Loaded %d %s records", len(dst.Data), name)
}

func LoadMagicPrefix(file []byte) {
	loadAffixes(file, MagicPrefixDictionary, "MagicPrefix")
}

func LoadMagicSuffix(file []byte) {
	loadAffixes(file, MagicSuffixDictionary, "MagicSuffix")
}

func LoadRarePrefix(file []byte) {
	loadAffixes(file, RarePrefixDictionary, "RarePrefix")
}

func LoadRareSuffix(file []byte) {
	loadAffixes(file, RareSuffixDictionary, "RareSuffix")
}

func LoadUniquePrefix(file []byte) {
	loadAffixes(file, UniquePrefixDictionary, "UniquePrefix")
}

func LoadUniqueSuffix(file []byte) {
	loadAffixes(file, UniqueSuffixDictionary, "UniqueSuffix")
}
