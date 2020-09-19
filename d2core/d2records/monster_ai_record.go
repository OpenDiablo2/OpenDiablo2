package d2records

// MonsterAI holds the MonsterAIRecords, The monai.txt file is a lookup table for unit AI codes
type MonsterAI map[string]*MonsterAIRecord

// MonsterAIRecord represents a single row from monai.txt
type MonsterAIRecord struct {
	AI string
}
