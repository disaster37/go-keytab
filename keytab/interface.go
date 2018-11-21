package keytab

type keytabInterface interface {
	IsExist() (bool, error)
	Create() (*Keytab, error)
	Delete() error
}
