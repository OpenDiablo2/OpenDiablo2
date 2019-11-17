package tests

import (
	"log"
	"path"
	"strings"
	"testing"

	"github.com/OpenDiablo2/D2Shared/d2data/d2cof"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"

	"github.com/OpenDiablo2/D2Shared/d2common"
)

func TestMPQScanPerformance(t *testing.T) {
	log.SetFlags(log.Ldate | log.LUTC | log.Lmicroseconds | log.Llongfile)
	d2mpq.InitializeCryptoBuffer()
	d2common.ConfigBasePath = "../"
	config := d2common.LoadConfiguration()
	engine := d2core.CreateEngine()
	for _, fileName := range config.MpqLoadOrder {
		mpqFile := path.Join(config.MpqPath, fileName)
		archive, _ := d2mpq.Load(mpqFile)
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
				_ = d2cof.LoadCOF(archiveFile, engine)
			case "dcc":
				if strings.ContainsAny(archiveFile, "common") {
					continue
				}
				_ = d2dcc.LoadDCC(archiveFile, engine)
			}

			_, _ = archive.ReadFile(archiveFile)
		}
	}
}
