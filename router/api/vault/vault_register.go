package vault

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	_ "embed"
	"encoding/hex"
	"encoding/pem"
	"log"
	"math/big"
	"os"

	vault "github.com/bytehubplus/fusion/node/vaultindex"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bytehubplus/fusion/cmd"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Keys string `json:"keys"`
	Sig  string `json:"sig"`
}

func RegistVault(c *gin.Context) {
	postdata := Message{}
	c.BindJSON(&postdata)
	//parse public Key
	pubKey, err := hex.DecodeString(postdata.Keys)
	if err != nil {
		log.Fatal(err)
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKey)
	lpuk := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	//get data signature hash
	data := []byte(postdata.Keys)
	hash := crypto.Keccak256Hash(data)
	//parse signature
	signature, err := hex.DecodeString(postdata.Sig)
	if err != nil {
		log.Fatal(err)
	}
	var rint, sint big.Int
	ab := bytes.Split(signature, []byte("+"))
	rint.UnmarshalText(ab[0])
	sint.UnmarshalText(ab[1])

	verify := ecdsa.Verify(lpuk, hash.Bytes(), &rint, &sint)
	if verify != true {
		c.JSON(400, "signature verify failed. please check youe signature.")
		panic(".go 51: code in vault_register.go failed. ")
	}
	//c.JSON(200, verify)
	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := vault.NewProvider(conf)
	//defer provider.CloseDB()
	vaultID, err := provider.RegisterVault(postdata.Keys)
	if err != nil {
		log.Fatal(err)
	}

	// create .key in ./key
	derPkix, err := x509.MarshalPKIXPublicKey(lpuk)
	if err != nil {
		panic(err)
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	fi, err := os.Create("./key/" + vaultID + ".pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(fi, block)

	// create DID document
	err = cmd.CreateDIDDoc("./key/"+vaultID+".pem", "./document/"+vaultID)
	if err != nil {
		log.Fatal(err)
	}

	respon := "Register Vault Success. The VaultID is:" + vaultID
	c.JSON(200, respon)

}
