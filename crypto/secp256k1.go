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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"hash"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Secp256k1PrivateKey struct {
	privKey *ecdsa.PrivateKey
}

func (s *Secp256k1PrivateKey) String() string {
	return string(crypto.FromECDSA(s.privKey))
}

func GenSecp256k1KeyPair() (Key, Key, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	pubKey := privKey.PublicKey
	return &Secp256k1PrivateKey{privKey: privKey}, &Secp256k1PublicKey{pubKey: &pubKey}, nil
}

// Hash hashes a message
func (e *Secp256k1PrivateKey) Hash(msg []byte, opts HashOpts) ([]byte, error) {
	h, err := e.GetHash(opts)
	if err != nil {
		return nil, err
	}
	return h.Sum(msg), nil
}

// GetHash returns the instance of hash function
func (e *Secp256k1PrivateKey) GetHash(opt HashOpts) (hash.Hash, error) {
	return sha3.NewLegacyKeccak256(), nil
}

func (e *Secp256k1PrivateKey) Encrypt(k Key, plaintext []byte, opts EncrypterOpts) ([]byte, error) {
	return nil, errors.New("this key does NOT support encrypt")
}

func (e *Secp256k1PrivateKey) Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) ([]byte, error) {
	return nil, errors.New("this key does NOT support decrypt")
}

// Sign signs a message's hash
func (e *Secp256k1PrivateKey) Sign(k Key, digest []byte, opts SignerOpts) ([]byte, error) {
	return crypto.Sign(digest, k.(*Secp256k1PrivateKey).privKey)
}

// Verify verifies a signature
func (e *Secp256k1PrivateKey) Verify(k Key, signature []byte, digest []byte, opts SignerOpts) (bool, error) {
	return false, errors.New("Secp256k1 does NOT support verify with private key")
}

// the key's raw byte
func (e *Secp256k1PrivateKey) Bytes() ([]byte, error) {
	return elliptic.Marshal(e.privKey.Curve, e.privKey.X, e.privKey.Y), nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (e *Secp256k1PrivateKey) PrivateKey() bool {
	return true
}

// symmetric returns true if this key is symmetric, otherwise false
func (e *Secp256k1PrivateKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise false
func (e *Secp256k1PrivateKey) PublicKey() (Key, error) {
	return &Secp256k1PublicKey{&e.privKey.PublicKey}, nil
}

func (secKey *Secp256k1PrivateKey) SaveToPem(file string) error {
	pemFile, err := os.Create(file)
	if err != nil {
		return err
	}

	bytes, _ := x509.MarshalECPrivateKey(secKey.privKey)
	var pemKey = &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: bytes,
	}
	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return err
	}
	return pemFile.Close()
}

type Secp256k1PublicKey struct {
	pubKey *ecdsa.PublicKey
}

// Verify verifies a signature
func (s *Secp256k1PublicKey) Verify(k Key, signature []byte, digest []byte) (bool, error) {
	pub := crypto.FromECDSAPub(s.pubKey)
	return crypto.VerifySignature(pub, digest, signature), nil
}

// the key's raw byte
func (e *Secp256k1PublicKey) Bytes() ([]byte, error) {
	return elliptic.Marshal(e.pubKey.Curve, e.pubKey.X, e.pubKey.Y), nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (e *Secp256k1PublicKey) PrivateKey() bool {
	return false
}

// symmetric returns true if this key is symmetric, otherwise false
func (e *Secp256k1PublicKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise false
func (e *Secp256k1PublicKey) PublicKey() (Key, error) {
	return e, nil
}
