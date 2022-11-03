package did

import (
	ssi "github.com/nuts-foundation/go-did"
	"github.com/nuts-foundation/go-did/did"
)

type DID did.DID

type Document did.Document

type URI ssi.URI

type KeyType ssi.KeyType

const (
	DIIDContext = did.DIDContextV1
)
