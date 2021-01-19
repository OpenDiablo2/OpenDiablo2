package d2enum

// PlayersRelationships represents players relationships
type PlayersRelationships int

// Players relationships
const (
	PlayerRelationNeutral PlayersRelationships = iota
	PlayerRelationFriend
	PlayerRelationEnemy
)

const (
	PlayersHostileLevel = 9
)

const (
	MaxPlayersInGame = 8
)
