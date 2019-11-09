package tests

import (
	"log"
	"path"
	"strings"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/core"

	"github.com/OpenDiablo2/OpenDiablo2/mpq"

	"github.com/OpenDiablo2/OpenDiablo2/common"
)

func TestMPQScanPerformance(t *testing.T) {
	log.SetFlags(log.Ldate | log.LUTC | log.Lmicroseconds | log.Llongfile)
	mpq.InitializeCryptoBuffer()
	common.ConfigBasePath = "../"
	config := common.LoadConfiguration()
	engine := core.CreateEngine()
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
			parts := strings.Split(archiveFile, ".")
			switch strings.ToLower(parts[len(parts)-1]) {
			case "coff":
				_ = common.LoadCof(archiveFile, engine)
			case "dcc":
				if strings.ContainsAny(archiveFile, "common") {
					continue
				}
				_ = common.LoadDCC(archiveFile, engine)
			}

			_, _ = archive.ReadFile(archiveFile)
		}
	}
}
