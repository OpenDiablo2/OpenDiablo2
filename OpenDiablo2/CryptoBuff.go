package OpenDiablo2

// CryptoBuffer contains the crypto bytes for filename hashing
var CryptoBuffer [0x500]uint32

// InitializeCryptoBuffer initializes the crypto buffer
func InitializeCryptoBuffer() {
	seed := uint32(0x00100001)
	for index1 := 0; index1 < 0x100; index1++ {
		index2 := index1
		for i := 0; i < 5; i++ {
			seed = (seed*125 + 3) % 0x2AAAAB
			temp1 := (seed & 0xFFFF) << 0x10
			seed = (seed*125 + 3) % 0x2AAAAB
			temp2 := (seed & 0xFFFF)
			CryptoBuffer[index2] = temp1 | temp2
			index2 += 0x100
		}
	}
}
