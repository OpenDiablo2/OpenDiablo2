package d2clientconnectiontype

// ClientConnectionType is an enum referring to types implementing
// d2server.ClientConnection and d2client.ClientConnection.
type ClientConnectionType int

//
const (
	Local     ClientConnectionType = iota // Local client
	LANServer                             // Server
	LANClient                             // Remote client
)
