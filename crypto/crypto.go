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
	gocrypto "crypto"
	"hash"
)

type Key interface {
	//the key's raw byte
	Bytes() ([]byte, error)
	//PrivateKey returns true is this is a asymmetric private key or symmetric security key
	PrivateKey() bool
	//symmetric returns true if this key is symmetric, otherwise false
	Symmetric() bool
	//if this is a asymmetric key, returns the corresponding Public key, otherwise error
	PublicKey() (Key, error)
}

// Key generation options for BHPCSP
type KeyGenOpts interface {
	Algorithem() string
}

// HashOpts contains hash options for BHPCSP
type HashOpts interface {
	Algorithem() string
}

// EncrypterOpts contains encrypting options
type EncrypterOpts interface {
}

// DecrypterOpts contains decrypting options
type DecrypterOpts interface {
}

// SignerOpts contain signing options
type SignerOpts interface {
	gocrypto.SignerOpts
}

// bytehub+ crytograhic service provider
type BHPCSP interface {
	//KeyGen generates a new key
	KeyGen(opts KeyGenOpts) (Key, error)

	//GetKey returns the key
	GetKey(keyInstance []byte) (Key, error)

	//Hash hashes a message
	Hash(msg []byte, opts HashOpts) ([]byte, error)

	//GetHash returns the instance of hash function
	GetHash(opt HashOpts) (hash.Hash, error)

	Encrypt(k Key, plaintext []byte, opts EncrypterOpts) ([]byte, error)
	Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) ([]byte, error)

	//Sign signs a message's hash
	Sign(k Key, digest []byte, opts SignerOpts) ([]byte, error)
	//Verify verifies a signature
	Verify(k Key, signature, digest []byte, opts SignerOpts) (bool, error)
}
