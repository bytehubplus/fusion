package node

import (
	"crypto"
	"crypto/ed25519"

	"github.com/bytehubplus/fusion/core/signer"
)

type Node struct {
	signer.Signer
	privateKey crypto.PrivateKey
}

func (n *Node) Sig(message []byte) {
	ed25519.Sign(n.privateKey.(ed25519.PrivateKey), message)
}
