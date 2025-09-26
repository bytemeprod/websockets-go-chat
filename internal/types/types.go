package types

type Client interface {
	ReadConnection()
	WriteConnection()
}

type Manager interface {
	AddClient(Client)
	RemoveClient(Client)
}
