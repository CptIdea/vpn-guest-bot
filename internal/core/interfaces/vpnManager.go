package interfaces

type VpnManager interface {
	Create(name string) (password string, err error)
	Delete(name string) error
}
