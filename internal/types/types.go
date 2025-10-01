package types

type Client interface {
	ReadConnection()
	WriteConnection()
	AddToEgress(message []byte)
	GetUsername() string
}

type Manager interface {
	AddClient(Client)
	RemoveClient(Client)
}
