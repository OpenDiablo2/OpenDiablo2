// Package d2animdata provides a file parser for AnimData files. AnimData files have the '.d2'
// file extension, but we do not call this package `d2d2` because multiple file types share this
// extension, and because the project namespace prefix makes it sound terrible.
package d2animdata

/*
The AnimData.d2 file is a binary file that contains speed and event data for animations.

The data is encoded as little-endian binary data

The file contents look like this:

type animData struct {
	blocks [256]struct{
		recordCount uint8
		records []struct{ // *see note below
			name               [8]byte    // last byte is always \0
			framesPerDirection uint32
			speed              uint16     // **see note below
			_                  uint16     // just padded 0's
			events             [144]uint8 // ***see not below
		}
	}
}

*NOTE: can contain 0 to 67 records

**NOTE: game fps is assumed to be 25, the speed is calculated as (25 * record.speed / 256)

**NOTE: Animation events can be one of `None`, `Attack`, `Missile`, `Sound`, or `Skill`

*/
