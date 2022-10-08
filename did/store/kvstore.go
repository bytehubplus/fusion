// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package store

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/bytehubplus/fusion/did"
	"github.com/syndtr/goleveldb/leveldb"
)

type StoreConfig struct {
	Schema string
	Method string
	DBPath string
}

type KvStore struct {
	db *leveldb.DB
}

type StoreProvider struct {
	Config StoreConfig
	Store  KvStore
}

func (s *StoreProvider) OpenStore() (Store, error) {
	return &s.Store, nil
}

func NewProvider(conf StoreConfig) (*StoreProvider, error) {
	db, err := leveldb.OpenFile(conf.DBPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}
	store := &KvStore{db: db}
	sp := &StoreProvider{
		Config: conf,
		Store:  *store,
	}
	defer db.Close()
	return sp, nil
}

func (sp *StoreProvider) CloseStore() {
	sp.Store.db.Close()
}

func (k *KvStore) SaveDocument(doc did.Document) (string, error) {
	key := k.KeyGenerate([]byte(doc.ID.String()))
	rawData, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}
	if err := k.db.Put(key, rawData, nil); err != nil {
		return "", fmt.Errorf("save document failed: %s", err)
	}

	return bytes.NewBuffer(key).String(), nil
}

func (k *KvStore) KeyGenerate(value []byte) []byte {
	h := sha256.Sum256(value)
	return []byte(fmt.Sprintf("%x", h[:20]))
}

func (k *KvStore) LoadDocument(key string) (did.Document, error) {
	b, err := k.db.Get([]byte(key), nil)
	doc := &did.Document{}

	if err != nil {
		return *doc, err
	}
	err = json.Unmarshal(b, doc)
	// doc, err = did.ParseDocument(b)
	return *doc, err
}
