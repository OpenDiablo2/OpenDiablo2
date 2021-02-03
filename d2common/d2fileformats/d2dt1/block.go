package d2dt1

// Block represents a DT1 block
type Block struct {
	unknown1    []byte
	unknown2    []byte
	X           int16
	Y           int16
	GridX       byte
	GridY       byte
	format      int16
	EncodedData []byte
	Length      int32
	FileOffset  int32
}

// Format returns block format
func (b *Block) Format() BlockDataFormat {
	if b.format == 1 {
		return BlockFormatIsometric
	}

	return BlockFormatRLE
}
