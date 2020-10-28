package d2records

import (
	"github.com/gravestench/akara"
)

// UniqueMonsterPrefix is a representation of a possible name prefix for a unique monster instance
// eg _Blood_ Wing the Quick
type UniqueMonsterPrefix = UniqueMonsterAffixRecord

// UniqueMonsterSuffix is a representation of a possible name suffix for a unique monster instance.
// eg. Blood Wing _the Quick_
type UniqueMonsterSuffix = UniqueMonsterAffixRecord

// UniqueMonsterAffixes is a map of UniqueMonsterAffixRecords. The key is the string table lookup key.
type UniqueMonsterAffixes map[string]*UniqueMonsterAffixRecord

// UniqueMonsterAffixRecord is a string table key and a bit vector for the possible monster types
// that the suffix can be used with.
type UniqueMonsterAffixRecord struct {
	StringTableKey   string
	MonsterTypeFlags *akara.BitSet
}
