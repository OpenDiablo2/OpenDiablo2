package d2player

type InputCallbackListener interface {
	OnPlayerMove(x, y float64)
	OnPlayerCast(missleID int, x, y float64)
}
