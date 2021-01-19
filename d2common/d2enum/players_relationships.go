package d2enum

// PlayersRelationships represents players relationships
type PlayersRelationships int

// Players relationships
const (
	PlayerRelationNeutral PlayersRelationships = iota
	PlayerRelationFriend
	PlayerRelationEnemy
)

// determinates a level, which both players should reach to go hostile
const (
	PlayersHostileLevel = 9
)

// determinates max players number for one game
const (
	MaxPlayersInGame = 8
)
