package main

import (
	"log"
	"os"

	check "github.com/bytehubplus/fusion/node/checkProof"
	"github.com/bytehubplus/fusion/node/vaultRegist"
	"github.com/gin-gonic/gin"
)

func main() {

	file, err := os.Open("private.pem")
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

	file, err = os.Open("public.pem")
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

	err = vaultRegist.CreateRSAentry(publicKey, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/register", vaultRegist.RegistVault)
	r.POST("/register_safe", vaultRegist.RegistVault_safe)
	r.POST("/checkproof", check.CheckProof)
	r.Run(":8080")
}
