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
	doc, err := ParseDocument([]byte(did1Json))
	if err != nil {
		t.Errorf("parse document failed: %s", err)
	}

	t.Logf("id: %s", doc.ID.String())
	for k, v := range doc.Controller {
		t.Logf("Controller %d: %s", k, v.String())
	}
	for k, v := range doc.VerificationMethod {
		t.Logf("Verfication Method %d: %s", k, v.ID)
	}
	for k, v := range doc.Authentication {
		t.Logf("Authentication %d: %s", k, v.String())
	}
	for k, v := range doc.AssertionMethod {
		t.Logf("Asseert Method %d: %s", k, v.String())
	}
	for k, v := range doc.CapabilityDelegation {
		t.Logf("Capability Delegation %d: %s", k, v.String())
	}
	for k, v := range doc.Service {
		t.Logf("Service %d: %s", k, v)
	}

	data, err := json.Marshal(doc)
	err = os.WriteFile("./marshal.json", data, 0644)
	if err != nil {
		t.Errorf("write file failed: %s", err)
	}

}
