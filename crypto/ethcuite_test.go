package crypto

import "testing"

func TestRSAKeyLoad(t *testing.T) {
	ecp := &EthCryptoSuite{}
	ecp.LoadRSAPrivateKeyFromPEM("./testdata/rsa-key.pem", "")
	t.Logf("rsa key in hex: %s", ecp.RSAKey.String())
	t.Log("done")
}
func TestSecp256k1KeyLoad(t *testing.T) {
	ecp := &EthCryptoSuite{}
	ecp.LoadPrivateKeyFromPEM("./testdata/secp256k1-key.pem", "")
	t.Logf("rsa key in hex: %s", ecp.PrivKey.String())
	t.Log("done")
}
