package types

type Client interface {
	ReadConnection()
	WriteConnection()
	AddToEgress(message []byte)
}

type Manager interface {
	AddClient(Client)
	RemoveClient(Client)
}
