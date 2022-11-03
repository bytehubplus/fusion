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

package crypto

import "testing"

func TestKeyGen(t *testing.T) {
	keyGenOpts := &Ed25519PrivateKey{}
	key, err := keyGenOpts.KeyGen(nil)
	b, _ := key.Bytes()
	if err != nil {
		t.Logf("key = %x", b)
	}
}

func TestSign(t *testing.T) {
	keyGenOpts := &Ed25519PrivateKey{}
	key, err := keyGenOpts.KeyGen(nil)
	if err != nil {
		t.Logf("key failed")
	}

	msg := []byte("Rich Zhao")
	signature, err := key.(*Ed25519PrivateKey).Sign(key, msg, nil)
	if err != nil {
		t.Logf("sign failed")
	}

	t.Logf("signature = %x", signature)
}
