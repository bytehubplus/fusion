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

package signer

import (
	"crypto"
	"crypto/ed25519"
	"fmt"

	"github.com/bytehubplus/fusion/did"
)

type Signer struct {
	did.Document
}

func (s Signer) PublicKey(index int) (string, crypto.PublicKey, error) {
	if index > len(s.VerificationMethod) {
		return "", nil, fmt.Errorf("index is greater than total verification method: %d", len(s.VerificationMethod))
	}

	vm := s.VerificationMethod[index]
	keyType := string(vm.Type)
	if keyType != "ED25519VerificationKey2018" {
		return "", nil, fmt.Errorf("invalid public key type")
	}
	key, err := s.VerificationMethod[index].PublicKey()
	if err != nil {
		return "", nil, err
	}
	return keyType, key, nil
}

func (s Signer) Verify(message, sig []byte) bool {
	for _, vm := range s.VerificationMethod {
		pubKey, err := vm.PublicKey()
		if err != nil {
			return false
		}
		if ed25519.Verify(pubKey.(ed25519.PublicKey), message, sig) {
			return true
		}
	}

	return false
}
