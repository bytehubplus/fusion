package controller

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	_ "embed"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Keys  string
	Token string
}

func Append(c *gin.Context) {
	keyInfo := "fusion"
	postdata := Message{}
	c.BindJSON(&postdata)
	//parse public Key
	pubKey, err := hex.DecodeString(postdata.Keys)
	if err != nil {
		log.Fatal(err)
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKey)
	lpuk := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	log.Println(lpuk)
	//将token字符串转换为token对象
	tokenInfo, _ := jwt.Parse(postdata.Token, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	//校验错误（基本）
	err = tokenInfo.Claims.Valid()
	if err != nil {
		log.Fatal(err)
	}

	finToken := tokenInfo.Claims.(jwt.MapClaims)
	//校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	if succ != true {
		c.JSON(400, "Token has expired. Please apply for the token again")
	}
	//获取token中保存的用户信息
	var vaultID string
	vaultID = fmt.Sprintf("%s", finToken["VaultID"])

	derPkix, err := x509.MarshalPKIXPublicKey(lpuk)
	if err != nil {
		panic(err)
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	fi, err := os.Create("./key/" + vaultID + "/" + postdata.Keys + ".pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(fi, block)
	if err != nil {
		panic(err)
	}
	c.JSON(200, "public key append success.")
}
