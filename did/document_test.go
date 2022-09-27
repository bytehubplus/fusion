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
package did

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"
)

var (
	//go:embed testdata/valid_doc.jsonld
	validDoc string
	//go:embed testdata/valid_doc_resolution.jsonld
	validDocResolution string
	//go:embed testdata/invalid_doc.jsonld
	invalidDoc string
	//go:embed testdata/valid_doc_v0.11.jsonld
	validDocV011 string
	//go:embed testdata/valid_doc_with_base.jsonld
	validDocWithBase string
	//go:embed testdata/did1.json
	did1Json string
)

func TestParseDocument(t *testing.T) {
	// doc, err := ParseDocument([]byte(did1Json))
	var doc Document
	err := json.Unmarshal([]byte(did1Json), &doc)
	if err != nil {
		t.Fatalf("parse document failed: %s", err)
	}

	t.Logf("id: %s", doc.ID.String())
	for k, v := range doc.Controller {
		t.Logf("Controller %d: %s", k, v.String())
	}
	for k, v := range doc.VerificationMethod {
		t.Logf("Verfication Method %d: %s", k, v.ID)
	}
	// for k, v := range doc.Authentication {
	// 	t.Logf("Authentication %d: %s", k, v.String())
	// }
	// for k, v := range doc.AssertionMethod {
	// 	t.Logf("Asseert Method %d: %s", k, v.String())
	// }
	// for k, v := range doc.CapabilityDelegation {
	// 	t.Logf("Capability Delegation %d: %s", k, v.String())
	// }
	for k, v := range doc.Service {
		t.Logf("Service %d: %s", k, v.ID.String())
	}

	data, err := json.Marshal(doc)
	err = os.WriteFile("./marshal.json", data, 0644)
	if err != nil {
		t.Errorf("write file failed: %s", err)
	}

}

func TestUnmarshalService(t *testing.T) {
	var input Service
	err := json.Unmarshal([]byte(`{
  		  "id":"did:example:123#linked-domain",
		  "type":"custom",
		  "serviceEndpoint": "did:nuts:123456"
}`), &input)

	if err != nil {
		t.Logf("Unmarshal service failed : %s", err)
	} else {
		t.Logf("Unmarshal service success. %s ", input.ServiceEndpoint[0])
	}

}
