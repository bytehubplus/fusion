package node

import (
	"crypto"
	"crypto/ed25519"
	"log"

	"github.com/bytehubplus/fusion/router"
	check "github.com/bytehubplus/fusion/router/api/check_proof"
	"github.com/bytehubplus/fusion/router/api/vault_regist"
	"github.com/gin-gonic/gin"

	"github.com/bytehubplus/fusion/core/signer"
)

type Node struct {
	signer.Signer
	privateKey crypto.PrivateKey
}

func (n *Node) Sig(message []byte) {
	ed25519.Sign(n.privateKey.(ed25519.PrivateKey), message)
}

func main() {
	/*
		routersInit := routers.InitRouter()
		endPoint := "http://127.0.0.1:8000"

		server := &http.Server{
			Addr:    endPoint,
			Handler: routersInit,
		}

		log.Printf("[info] start http server listening %s", endPoint)

		server.ListenAndServe()
	*/

	// Read Key-pair from folder key
	publicKey, privateKey := router.Readkey()

	// Initialize an xrsa instance and decrypt the received message with node's public key.
	// function CreateKeys() will create 2 file to save public and private key.
	/*
		publicKey := bytes.NewBufferString("")
		privateKey := bytes.NewBufferString("")
		err := xrsa.CreateKeys(publicKey, privateKey, 2048)
		if err != nil {
			log.Println(err)
			return
		}
	*/

	err := vault_regist.CreateRSAentry(publicKey, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/register", vault_regist.RegistVault)
	r.POST("/register_safe", vault_regist.RegistVault_safe)
	r.POST("/checkproof", check.CheckProof)
	r.Run(":8080")
}
