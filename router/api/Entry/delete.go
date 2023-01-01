package entry

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	vault "github.com/bytehubplus/fusion/node/vaultindex"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type Message_Delete struct {
	Entry_ID string
	Token    string
}

func Delete(c *gin.Context) {
	keyInfo := "fusion"
	postdata := Message_Get{}
	c.BindJSON(&postdata)
	//将token字符串转换为token对象
	tokenInfo, _ := jwt.Parse(postdata.Token, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
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

	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/" + vaultID,
	}
	provider, _ := vault.NewProvider(conf)
	err = provider.Delete(postdata.Entry_ID)
	if err != nil {
		c.JSON(400, "Delete data failed.")
	} else {
		c.JSON(200, "Delete data success.")
	}

}
