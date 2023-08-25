package interfaces

type ITransaction interface {
	Transfer(fromid int, toid int, amount int) (string, error)
}
