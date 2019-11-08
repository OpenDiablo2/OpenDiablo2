package tests

import (
	"path"
	"strings"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/mpq"

	"github.com/OpenDiablo2/OpenDiablo2/common"
)

func TestMPQScanPerformance(t *testing.T) {
	mpq.InitializeCryptoBuffer()
	common.ConfigBasePath = "../"
	config := common.LoadConfiguration()
	for _, fileName := range config.MpqLoadOrder {
		mpqFile := path.Join(config.MpqPath, fileName)
		archive, _ := mpq.Load(mpqFile)
		files, err := archive.GetFileList()
		if err != nil {
			continue
		}
		for _, archiveFile := range files {
			// Temporary until all audio formats are supported
			if strings.Contains(archiveFile, ".wav") || strings.Contains(archiveFile, ".pif") {
				continue
			}
			_, _ = archive.ReadFile(archiveFile)
		}
	}
}
