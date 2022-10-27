package crypto

import "hash"

type RSAPrivateKey struct {
	csi []byte
	// pub RSAPublicKey

}

// KeyGen generates a new key
func (r *RSAPrivateKey) KeyGen(opts KeyGenOpts) (Key, error) {
	panic("not implemented") // TODO: Implement
}

// GetKey returns the key
func (r *RSAPrivateKey) GetKey(keyInstance []byte) (Key, error) {
	panic("not implemented") // TODO: Implement
}

// Hash hashes a message
func (r *RSAPrivateKey) Hash(msg []byte, opts HashOpts) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// GetHash returns the instance of hash function
func (r *RSAPrivateKey) GetHash(opt HashOpts) (hash.Hash, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RSAPrivateKey) Encrypt(k Key, plaintext []byte, opts EncrypterOpts) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RSAPrivateKey) Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// Sign signs a message's hash
func (r *RSAPrivateKey) Sign(k Key, digest []byte, opts SignerOpts) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// Verify verifies a signature
func (r *RSAPrivateKey) Verify(k Key, signature []byte, digest []byte, opts SignerOpts) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// the key's raw byte
func (r *RSAPrivateKey) Bytes() ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (r *RSAPrivateKey) PrivateKey() bool {
	panic("not implemented") // TODO: Implement
}

// symmetric returns true if this key is symmetric, otherwise false
func (r *RSAPrivateKey) Symmetric() bool {
	panic("not implemented") // TODO: Implement
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise error
func (r *RSAPrivateKey) PublicKey() (Key, error) {
	panic("not implemented") // TODO: Implement
}

// type RSAPublicKey struct {
// 	csi []byte
//   pub xrsa.XRsa.
// }
