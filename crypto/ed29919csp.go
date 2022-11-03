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

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"hash"
)

type Ed25519PrivateKey struct {
	csp []byte
	pub Ed25519PublicKey
}

// KeyGen generates a new key
func (e *Ed25519PrivateKey) KeyGen(opts KeyGenOpts) (Key, error) {
	return e.generateEd25519Key()
}

func (k *Ed25519PrivateKey) generateEd25519Key() (Key, error) {
	pub, ski, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, errors.New("failed to genenate ed25519 key")
	}
	key := &Ed25519PrivateKey{ski, Ed25519PublicKey{ski, &pub}}
	return key, nil
}

// Hash hashes a message
func (e *Ed25519PrivateKey) Hash(msg []byte, opts HashOpts) ([]byte, error) {
	h, err := e.GetHash(opts)
	if err != nil {
		return nil, err
	}
	return h.Sum(msg), nil
}

// GetHash returns the instance of hash function
func (e *Ed25519PrivateKey) GetHash(opt HashOpts) (hash.Hash, error) {
	return sha512.New(), nil
}

func (e *Ed25519PrivateKey) Encrypt(k Key, plaintext []byte, opts EncrypterOpts) ([]byte, error) {
	return nil, errors.New("this key does NOT support encrypt")
}

func (e *Ed25519PrivateKey) Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) ([]byte, error) {
	return nil, errors.New("this key does NOT support decrypt")
}

// Sign signs a message's hash
func (e *Ed25519PrivateKey) Sign(k Key, digest []byte, opts SignerOpts) ([]byte, error) {
	return ed25519.Sign(e.csp, digest), nil
}

// Verify verifies a signature
func (e *Ed25519PrivateKey) Verify(k Key, signature []byte, digest []byte, opts SignerOpts) (bool, error) {
	return ed25519.Verify(e.pub.csi, digest, signature), nil
}

// the key's raw byte
func (e *Ed25519PrivateKey) Bytes() ([]byte, error) {
	return e.csp, nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (e *Ed25519PrivateKey) PrivateKey() bool {
	return true
}

// symmetric returns true if this key is symmetric, otherwise false
func (e *Ed25519PrivateKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise false
func (e *Ed25519PrivateKey) PublicKey() (Key, error) {
	return &e.pub, nil
}

type Ed25519PublicKey struct {
	csi []byte
	pub *ed25519.PublicKey
}

// the key's raw byte
func (e *Ed25519PublicKey) Bytes() ([]byte, error) {
	return e.csi, nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (e *Ed25519PublicKey) PrivateKey() bool {
	return false
}

// symmetric returns true if this key is symmetric, otherwise false
func (e *Ed25519PublicKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise false
func (e *Ed25519PublicKey) PublicKey() (Key, error) {
	return e, nil
}
