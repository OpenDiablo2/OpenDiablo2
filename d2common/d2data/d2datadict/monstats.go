package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

var MonStatsDictionary *d2common.DataDictionary

func LoadMonStats(file []byte) {
	MonStatsDictionary = d2common.LoadDataDictionary(string(file))
}
