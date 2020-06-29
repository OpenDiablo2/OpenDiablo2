package d2dt1

// Block represents a DT1 block
type Block struct {
	X           int16
	Y           int16
	GridX       byte
	GridY       byte
	Format      BlockDataFormat
	EncodedData []byte
	Length      int32
	FileOffset  int32
}
