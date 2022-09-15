package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type Proof struct {
	Request   string
	Vault_ID  string
	Entry_ID  string
	Signature string
}

func main() {
	r := gin.Default()
	r.POST("/proof", checkProof)
	r.Run(":8080")
}

func checkProof(c *gin.Context) {
	json := Proof{}
	c.BindJSON(&json)
	str := json.Request + json.Entry_ID + json.Vault_ID
	hash := crypto.Keccak256Hash([]byte(str))
	sig, _ := hexutil.Decode(json.Signature)
	sigNoID := sig[:len(sig)-1]
	//	TODO:
	//	According to the Vault_ID and Entry_ID, find the Signer's Public key from database
	pubk := "0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05"
	pubk_byte, _ := hexutil.Decode(pubk)
	verified := crypto.VerifySignature(pubk_byte, hash.Bytes(), sigNoID)
	if verified {
		c.JSON(200, "Verification pass. You have access")
		//	TODO:
		//	return data from database
	} else {
		c.JSON(400, "permission denied")
	}
}
