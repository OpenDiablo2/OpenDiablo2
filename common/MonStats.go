package common

import "github.com/OpenDiablo2/OpenDiablo2/resourcepaths"

var MonStatsDictionary *DataDictionary

func LoadMonStats(fileProvider FileProvider) {
	MonStatsDictionary = LoadDataDictionary(string(fileProvider.LoadFile(resourcepaths.MonStats)))
}
