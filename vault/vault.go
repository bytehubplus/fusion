package vault

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/bytehubplus/fusion/did"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	// Vault meta data prefix
	MetaPrefix []byte = []byte("VM")
	// Vault data entry prefix
	EntryPrefix []byte = []byte("VE")
)

// KvVault, Vault in KV database
type KvVault struct {
	db   *leveldb.DB
	Did  did.DID
	lock sync.RWMutex
}

// Controllers returns vault's controllers
func (k *KvVault) Controllers() []string {
	var result []string
	// rawData, err := k.Get("doc")
	// if err != nil {
	// 	return result
	// }

	return result
	// doc, err := did.ParseDocument(rawData)
	// for _, v := range doc.Controllers {

	// }
}

// PutEntry saves an entry data into vault, return entry's unique id if successful, otherwise return error
func (k *KvVault) PutEntry(entry []byte) ([]byte, error) {
	hash := sha256.Sum256(entry)
	key := fmt.Sprintf("%s%s", EntryPrefix, hash[:])
	err := k.Put(key, entry)
	if err != nil {
		return hash[:], nil
	}
	return nil, err
}

func (k *KvVault) GetEntry(Id string) ([]byte, error) {
	key := fmt.Sprintf("%s%s", EntryPrefix, Id)
	return k.Get(key)
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
func (p *Provider) CreateVault(doc did.Document) (Vault, error) {
	//create but not open existing
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/%s", p.RootFSPath, p.createVaultID(doc.ID)), &opt.Options{ErrorIfExist: true})
	if err != nil {
		return nil, err
	}

	vault := KvVault{db: db}
	didValue := doc.ID.String()
	vault.Put("did", []byte(didValue))
	raw, _ := doc.ID.MarshalJSON()
	vault.Put("doc", raw)

	return &vault, nil
}

func (p *Provider) createVaultID(did did.DID) string {
	hash := sha256.Sum256([]byte(did.String()))
	hexStr := hex.EncodeToString(hash[:])
	return hexStr
}
