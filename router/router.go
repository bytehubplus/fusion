package router

import (
	"log"
	"os"

	check "github.com/bytehubplus/fusion/router/api/check_proof"
	"github.com/bytehubplus/fusion/router/api/vault_regist"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	publicKey, privateKey := Readkey()

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

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	//此处进行参数解析操作
	//api.Use(jwt.JWT())
	{
		api.POST("/register", vault_regist.RegistVault)
		api.POST("/register_safe", vault_regist.RegistVault_safe)
		api.POST("/checkproof", check.CheckProof)
	}
	return r
}

func Readkey() ([]byte, []byte) {
	file, err := os.Open("key/private.pem")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Get file information
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice to store the public key information read from the file
	privateKey := make([]byte, info.Size())
	// Read public key file
	file.Read(privateKey)
	file.Close()

	file, err = os.Open("key/public.pem")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Get file information
	info, err = file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice to store the public key information read from the file
	publicKey := make([]byte, info.Size())
	// Read public key file
	file.Read(publicKey)
	file.Close()
	return publicKey, privateKey
}
