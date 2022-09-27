// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com

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

package vaultindex

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	Scheme string
	Method string
	DBPath string
}

type IndexProvider struct {
	Config Config
	db     *leveldb.DB
}

// shorten and undetectable
func (p *IndexProvider) generateVaultID(id string) string {
	did := fmt.Sprintf("%s:%s:%s", p.Config.Scheme, p.Config.Method, id)
	h := sha256.Sum256([]byte(did))
	return fmt.Sprintf("%x", h[:20])
}

// register a vault, return vault ID
func (i *IndexProvider) RegisterVault(id string) (string, error) {
	vaultID := i.generateVaultID(id)
	err := i.db.Put([]byte(vaultID), []byte(id), nil)
	if err != nil {
		return "", err
	}

	return vaultID, nil
}

// unregiste a vault
func (i *IndexProvider) UnregisterVault(id string) error {
	vaultID := i.generateVaultID(id)
	err := i.db.Delete([]byte(vaultID), nil)
	if err != nil {
		return errors.New(fmt.Sprintf("unregister vault failed: %s", err))
	}
	return nil
}

// Check if a vault exist
func (p *IndexProvider) VaultExits(id string) bool {
	// did := fmt.Sprintf("%s:%s:%s", p.Config.Scheme, p.Config.Method, id)
	exist, err := p.db.Get([]byte(id), nil)
	if err != nil {
		return false
	}
	return exist != nil || len(exist) > 0
}

func NewProvider(conf Config) (*IndexProvider, error) {
	l, err := leveldb.OpenFile(conf.DBPath, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open database: %s", err))
	}
	defer l.Close()
	p := &IndexProvider{Config: conf, db: l}
	return p, nil
}
