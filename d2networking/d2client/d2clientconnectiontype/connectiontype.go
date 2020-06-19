package d2clientconnectiontype

type ClientConnectionType int

const (
	Local ClientConnectionType = iota
	LANServer
	LANClient
)
