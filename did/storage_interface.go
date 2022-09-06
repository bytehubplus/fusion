package did

type Provider interface {
	OpenStore(name string) (Store, error)
	Close() error
}

type Store interface {
	Put(did string, data []byte) error
	Get(did string) ([]byte, error)
	Delete(did string) error
	Query(prefix string) (Iterator, error)
}

type Iterator interface {
	Next() (bool, error)
	Did() (DID, error)
	Data() ([]byte, error)
	Close() error
}
