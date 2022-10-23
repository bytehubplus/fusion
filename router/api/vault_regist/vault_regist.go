package vault_regist

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"xrsa"

	"github.com/bytehubplus/fusion/did"
	store "github.com/bytehubplus/fusion/did/store"
	"github.com/gin-gonic/gin"
)

var (
	//go:embed test/did1.json
	did1Json string
)

type Rsa_node struct {
	Xrsa *xrsa.XRsa
}

type Message struct {
	Email  string
	Phone  string
	DID    string
	PubKey string
}

// Global Rsa Object
var RSAENTRY = &Rsa_node{}

func main() {

	file, err := os.Open("private.pem")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	// Get file information
	info, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	// Create a slice to store the public key information read from the file
	privateKey := make([]byte, info.Size())
	// Read public key file
	file.Read(privateKey)
	file.Close()

	file, err = os.Open("public.pem")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	// Get file information
	info, err = file.Stat()
	if err != nil {
		log.Println(err)
		return
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
		}*/

	RSAENTRY.Xrsa, err = xrsa.NewXRsa(publicKey, privateKey)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(publicKey.String())
	//log.Println(privateKey.String())

	r := gin.Default()
	r.POST("/register", RegistVault)
	r.POST("/register_safe", RegistVault_safe)
	r.Run(":8080")

}

func CreateRSAentry(pub, pri []byte) error {
	var err error
	RSAENTRY.Xrsa, err = xrsa.NewXRsa(pub, pri)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func RegistVault_safe(c *gin.Context) {
	postdata := Message{}
	c.BindJSON(&postdata)

	//log.Println(reflect.TypeOf(postdata.pubKey))
	decrypted_node, err := RSAENTRY.Xrsa.PrivateDecrypt(postdata.DID)
	if err != nil {
		log.Println(err)
		return
	}
	xrsa_client, err := xrsa.NewXRsa([]byte(postdata.PubKey), nil)
	if err != nil {
		log.Println(err)
		return
	}
	decrypted, err := xrsa_client.PublicDecrypt(decrypted_node)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(decrypted)
	de_byte := []byte(decrypted)
	//log.Println(de_byte)
	var doc did.Document
	_ = json.Unmarshal(de_byte, &doc)
	//log.Println(doc)

	conf := &store.StoreConfig{
		DBPath: "./data",
		Schema: "did",
		Method: "rich",
	}
	sp, err := store.NewProvider(*conf)
	if err != nil {
		c.JSON(200, "create store provider failed")
		return
	}
	sto, err := sp.OpenStore()
	defer sp.CloseStore()
	key, err := sto.SaveDocument(doc)
	if err != nil {
		c.JSON(200, "Save did-document failed. please retry")
		return
	}
	resp := "Save did-document done. the key is : "
	resp = resp + key
	c.JSON(http.StatusOK, resp)
}

func RegistVault(c *gin.Context) {
	postdata := did.Document{}
	c.BindJSON(&postdata)

	conf := &store.StoreConfig{
		DBPath: "./data",
		Schema: "did",
		Method: "rich",
	}
	sp, err := store.NewProvider(*conf)
	if err != nil {
		c.JSON(200, "create store provider failed")
		return
	}

	sto, err := sp.OpenStore()
	defer sp.CloseStore()
	key, err := sto.SaveDocument(postdata)
	if err != nil {
		c.JSON(200, "Save did-document failed. please retry")
		return
	}
	resp := "Save did-document done. the key is : "
	resp = resp + key
	c.JSON(http.StatusOK, resp)
}
