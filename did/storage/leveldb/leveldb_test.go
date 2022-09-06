package leveldb

import (
	_ "embed"
	"testing"
)

//nolint:gochecknoglobals
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
)

func TestNewProvider(t *testing.T) {
	p := NewProvider(".")
	t.Logf("Store database path: %s", p.dbPath)
}

func TestOpenStore(t *testing.T) {
	p := NewProvider("data")
	p.OpenStore("did")
	t.Logf("store %s opened", p.dbs["did"].name)
}

func TestSaveDiD(t *testing.T) {
	p := NewProvider("data")
	p.OpenStore("did")
	s := p.dbs["did"]
	s.Put("did:example:21tDAKCERh95uGgKbJNHYp", []byte(validDoc))

	doc, err := s.Get("did:example:21tDAKCERh95uGgKbJNHYp")
	if err != nil {
		t.Fatalf("cannot read did from store")
	}

	t.Logf("did content:\n%s", doc)

}
