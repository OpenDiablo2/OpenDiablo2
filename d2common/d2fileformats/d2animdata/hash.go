package d2animdata

import "strings"

type hashTable [numBlocks]byte

func hashName(name string) byte {
	hashBytes := []byte(strings.ToUpper(name))

	var hash uint32
	for hashByteIdx := range hashBytes {
		hash += uint32(hashBytes[hashByteIdx])
	}

	hash %= numBlocks

	return byte(hash)
}
