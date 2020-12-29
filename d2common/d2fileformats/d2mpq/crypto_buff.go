package d2mpq

var cryptoBuffer [0x500]uint32 //nolint:gochecknoglobals // will fix later..
var cryptoBufferReady bool     //nolint:gochecknoglobals // will fix later..

func cryptoLookup(index uint32) uint32 {
	if !cryptoBufferReady {
		cryptoInitialize()

		cryptoBufferReady = true
	}

	return cryptoBuffer[index]
}

//nolint:gomnd // magic cryptographic stuff here...
func cryptoInitialize() {
	seed := uint32(0x00100001)

	for index1 := 0; index1 < 0x100; index1++ {
		index2 := index1

		for i := 0; i < 5; i++ {
			seed = (seed*125 + 3) % 0x2AAAAB
			temp1 := (seed & 0xFFFF) << 0x10
			seed = (seed*125 + 3) % 0x2AAAAB
			temp2 := seed & 0xFFFF
			cryptoBuffer[index2] = temp1 | temp2
			index2 += 0x100
		}
	}
}
