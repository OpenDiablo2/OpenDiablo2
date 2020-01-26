package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

var MonStatsDictionary *d2common.DataDictionary

func LoadMonStats(fileProvider d2interface.FileProvider) {
	MonStatsDictionary = d2common.LoadDataDictionary(string(fileProvider.LoadFile(d2resource.MonStats)))
}
