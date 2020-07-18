package d2gamescreen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
)

// Navigator is used for transitioning between game screens
type Navigator interface {
	ToMainMenu()
	ToSelectHero(connType d2clientconnectiontype.ClientConnectionType, connHost string)
	ToCreateGame(filePath string, connType d2clientconnectiontype.ClientConnectionType, connHost string)
	ToCharacterSelect(connType d2clientconnectiontype.ClientConnectionType, connHost string)
	ToMapEngineTest(region int, level int)
	ToCredits()
}
