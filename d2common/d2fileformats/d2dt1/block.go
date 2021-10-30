package d2dt1

// Block represents a DT1 block
type Block struct {
	EncodedData []byte
	Length      int32
	FileOffset  int32
	format      int16
	X           int16
	Y           int16
	GridX       byte
	GridY       byte
}

// Format returns block format
func (b *Block) Format() BlockDataFormat {
	if b.format == 1 {
		return BlockFormatIsometric
	}

	return BlockFormatRLE
}

func (b *Block) unknown1() []byte {
	return make([]byte, numUnknownBlockBytes)
}

func (b *Block) unknown2() []byte {
	return make([]byte, numUnknownBlockBytes)
}
