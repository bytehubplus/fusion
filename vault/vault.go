package vault

import (
	"github.com/bytehubplus/did/did"
)

type Vault struct {
}

func (v *Vault) Read(did did.DID, proof did.DID) ([]byte, error) {
	return nil, nil
}
