package d2enum

// PlayersRelationships represents players relationships
type PlayersRelationships int

// Players relationships
const (
	PlayerRelationNeutral PlayersRelationships = iota
	PlayerRelationFriend
	PlayerRelationEnemy
)
