package crypto

import (
	"bytes"
	"testing"
)

func TestEncrypt(t *testing.T) {
	aesKey := &AESKey{secKey: []byte("richzhaorichzhaorichzhaorichzhao")}
	plaintext := []byte("richzhao @ bytehub+")
	ciphertext, err := aesKey.Encrypt(aesKey, plaintext)
	if err != nil {
		t.Log(err)
	}

	t.Logf("ciphertext = %x", ciphertext)

	decypter, err := aesKey.Decrypt(aesKey, ciphertext)
	if err != nil {
		t.Log(err)
	}
	t.Logf("ciphertext = %x", decypter)
	if bytes.Compare(plaintext, decypter) != 0 {
		t.Error("failed")
	}
}
