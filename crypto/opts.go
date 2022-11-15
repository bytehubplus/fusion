// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhu@gmail.com>

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

const (
	// ESECP256K1 p256k1 clliptic curve, for signing and verifying, not for encrypt nor decrypt
	SECP256K1 = "ESECP256K1"

	RSA = "RSA"
	AES = "AES"

	SHA256 = "KECCAK256"
)

type SECP256K1GenOpts struct {
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SECP256K1GenOpts) Algorithm() string {
	return SECP256K1
}

// AESKeyGenOpts contains options for AES key generation at default security level
type AESKeyGenOpts struct {
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *AESKeyGenOpts) Algorithm() string {
	return AES
}

// SHAOpts contains options for computing SHA.
type SHAOpts struct{}

// Algorithm returns the hash algorithm identifier (to be used).
func (opts *SHAOpts) Algorithm() string {
	return SHA256
}

type Keccak256Hash struct {
}

func (k *Keccak256Hash) Algorithm() string {
	return SHA256
}
