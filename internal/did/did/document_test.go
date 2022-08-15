package did

import (
	_ "embed"
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
	// t.Logf("@context: %s", doc.Context[0].String())

}
