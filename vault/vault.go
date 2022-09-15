package vault

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/bytehubplus/did/did"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// / KvVault, Vault in KV database
type KvVault struct {
	db   *leveldb.DB
	Did  did.DID
	lock sync.RWMutex
}

func (k *KvVault) GetEntry(Id string) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

func (k *KvVault) Get(key string) ([]byte, error) {
	return k.db.Get([]byte(key), nil)

}

func (k *KvVault) Put(Key string, value []byte) error {
	return k.db.Put([]byte(Key), value, nil)

}

func (k *KvVault) Delete(key string) error {
	return k.db.Delete([]byte(key), nil)
}

func (k *KvVault) VaultID() string {
	hash := sha256.Sum256([]byte(k.Did.String()))
	return string(hash[:])
}

type Provider struct {
	RootFSPath string
	// Config     Config
}

func NewProvider(path string) *Provider {
	p := &Provider{RootFSPath: path}
	return p
}

func (p *Provider) Open(id string) (Vault, error) {
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/%s", p.RootFSPath, id), &opt.Options{ErrorIfMissing: true})
	if err != nil {
		return nil, err
	}

	vault := KvVault{db: db}
	return &vault, nil
}

func (p *Provider) OpenWithDid(did did.DID) (Vault, error) {
	return p.Open(did.String())
}

// CreateVault creates a new vault
// param
func (p *Provider) CreateVault(did did.DID) (Vault, error) {
	//create but not open existing
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/%s", p.RootFSPath, p.createVaultID(did)), &opt.Options{ErrorIfExist: true})
	if err != nil {
		return nil, err
	}

	vault := KvVault{db: db}
	didValue := did.String()
	vault.Put("did", []byte(didValue))

	return &vault, nil
}

func (p *Provider) createVaultID(did did.DID) string {
	hash := sha256.Sum256([]byte(did.String()))
	hexStr := hex.EncodeToString(hash[:])
	return hexStr
}
