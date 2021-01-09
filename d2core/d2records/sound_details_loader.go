package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// Loadrecords loads SoundEntries from sounds.txt
func soundDetailsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(SoundDetails)

	for d.Next() {
		entry := &SoundDetailsRecord{
			Handle:    d.String("Sound"),
			Index:     d.Number("Index"),
			FileName:  d.String("FileName"),
			Volume:    d.Number("Volume"),
			GroupSize: d.Number("Group Size"),
			Loop:      d.Bool("Loop"),
			FadeIn:    d.Number("Fade In"),
			FadeOut:   d.Number("Fade Out"),
			DeferInst: d.Bool("Defer Inst"),
			StopInst:  d.Bool("Stop Inst"),
			Duration:  d.Number("Duration"),
			Compound:  d.Number("Compound"),
			Reverb:    d.Number("Reverb"),
			Falloff:   d.Number("Falloff"),
			Cache:     d.Bool("Cache"),
			AsyncOnly: d.Bool("Async Only"),
			Priority:  d.Number("Priority"),
			Stream:    d.Bool("Stream"),
			Stereo:    d.Bool("Stereo"),
			Tracking:  d.Bool("Tracking"),
			Solo:      d.Bool("Solo"),
			MusicVol:  d.Bool("Music Vol"),
			Block1:    d.Number("Block 1"),
			Block2:    d.Number("Block 2"),
			Block3:    d.Number("Block 3"),
		}

		records[entry.Handle] = entry
	}

	if d.Err != nil {
		return d.Err
	}

	r.Sound.Details = records

	r.Logger.Infof("Loaded %d sound definitions", len(records))

	return nil
}
