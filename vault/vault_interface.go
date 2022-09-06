package vault

import "github.com/syndtr/goleveldb/leveldb"

type VaultDBConfig struct {
	Database string
	VaultID  string
	DBConfig *LevelDBConfig
}

type LevelDBConfig struct {
	Path     string
	Database *leveldb.DB
}

type VaultConfig struct {
	DBConfig *VaultDBConfig
}

type VaultInterface interface {
}
