package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

type RSAPrivateKey struct {
	PrivKey *rsa.PrivateKey
	PubKey  RSAPublicKey
}

func (r *RSAPrivateKey) String() string {
	b, _ := r.Bytes()
	return fmt.Sprintf("%x", b)
}

// the key's raw byte
func (k *RSAPrivateKey) Bytes() ([]byte, error) {
	return x509.MarshalPKCS1PrivateKey(k.PrivKey), nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (r *RSAPrivateKey) PrivateKey() bool {
	return true
}

// symmetric returns true if this key is symmetric, otherwise false
func (r *RSAPrivateKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise error
func (k *RSAPrivateKey) PublicKey() (Key, error) {
	return &RSAPublicKey{&k.PrivKey.PublicKey}, nil
}

type RSAPublicKey struct {
	pubKey *rsa.PublicKey
}

// the key's raw byte
func (k *RSAPublicKey) Bytes() ([]byte, error) {
	return x509.MarshalPKCS1PublicKey(k.pubKey), nil
}

// PrivateKey returns true is this is a asymmetric private key or symmetric security key
func (r *RSAPublicKey) PrivateKey() bool {
	return false
}

// symmetric returns true if this key is symmetric, otherwise false
func (r *RSAPublicKey) Symmetric() bool {
	return false
}

// if this is a asymmetric key, returns the corresponding Public key, otherwise error
func (r *RSAPublicKey) PublicKey() (Key, error) {
	return r, nil
}

func NewRSAPrivateKey() (Key, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	pubKey := RSAPublicKey{pubKey: &privKey.PublicKey}
	return &RSAPrivateKey{PrivKey: privKey, PubKey: pubKey}, nil
}
