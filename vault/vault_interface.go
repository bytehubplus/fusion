package vault

import "github.com/bytehubplus/did/did"

// vault is a sensetive data storage unit which can store one or more data entry in KV format
// did used as vault for being identification, authentication/authorization etc.
// each entry maybe encrypted
// a read reqeust contains:
// 1. vault ID,
// 2. entry id
// 3. encryption key, which can be used to encryp it before sending to requestor
type Vault interface {
	Get(key string) ([]byte, error)
	Put(Key string, value []byte) error
	Delete(key string) error
	VaultID() string
	Controllers() []string
	GetEntry(Id string) ([]byte, error)
	PutEntry(entry []byte) ([]byte, error)
}

type VaultProvider interface {
	CreateVault(config Config) (Vault, error)
	Open(id string) (Vault, error)
	OpenWithDid(did did.DID) (Vault, error)
}

type Config struct {
	// RootFSPath is the directory where vault is stored
	RootFSPath string
	// DID        did.DID
	// DBConfig   *DBConfig
}

type DBConfig struct {
	DBPath string
}
