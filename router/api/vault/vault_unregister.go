package vault

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"

	vault "github.com/bytehubplus/fusion/node/vaultindex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type UGmsg struct {
	VaultID string
	Sig     string
}

func Unregister(c *gin.Context) {

	postdata := UGmsg{}
	c.BindJSON(&postdata)
	publicKey, err := ioutil.ReadFile("./key/" + postdata.VaultID + ".pem")
	if publicKey == nil {
		panic("public key file is nil ")
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		panic(".go 59  code run failed in create.go")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	pub := pubInterface.(*ecdsa.PublicKey)

	data := []byte(postdata.VaultID)
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

	verify := ecdsa.Verify(pub, hash.Bytes(), &rint, &sint)
	if verify != true {
		c.JSON(400, "signature verify failed. please check youe signature.")
		panic(".go 56: code in vault_unregister.go failed. ")
	}

	vaultID := postdata.VaultID

	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := vault.NewProvider(conf)
	exist := provider.VaultExits(postdata.VaultID)
	if exist == false {
		c.JSON(200, "VaultID not exist. please check your ID")
	} else {
		if err := provider.UnregisterVault(vaultID); err != nil {
			c.JSON(200, "Unregister failed.")
		} else {
			c.JSON(200, "Unregister Success.")
		}
	}

}
