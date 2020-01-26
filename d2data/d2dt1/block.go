package d2dt1

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
