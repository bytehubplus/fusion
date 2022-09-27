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
	"encoding/json"
	"testing"

	_ "embed"

	"github.com/bytehubplus/fusion/did"
)

var (
	//go:embed test/did1.json
	did1Json string
)

func TestSaveDocument(t *testing.T) {
	conf := &StoreConfig{
		DBPath: "./data",
		Schema: "did",
		Method: "rich",
	}
	sp, err := NewProvider(*conf)
	if err != nil {
		t.Logf("create store provider failed")
	}

	store, err := sp.OpenStore()
	var doc did.Document
	json.Unmarshal([]byte(did1Json), &doc)
	// doc, err := did.ParseDocument([]byte(did1Json))
	key, err := store.SaveDocument(doc)
	sp.CloseStore()
	t.Logf("DID Document saved. key : %s\n", key)
}

func TestLoadDocument(t *testing.T) {

	conf := &StoreConfig{
		DBPath: "./data",
		Schema: "did",
		Method: "rich",
	}
	sp, err := NewProvider(*conf)
	if err != nil {
		t.Logf("create store provider failed")
	}

	store, err := sp.OpenStore()

	key := "9be627fdc3dae7d17f13fafa6df69d0f5325ba69"
	document, err := store.LoadDocument(key)

	sp.CloseStore()
	t.Logf("did : %s", document.ID.String())
}
