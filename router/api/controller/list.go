package controller

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type Message_list struct {
	Token string
}

func List(c *gin.Context) {
	keyInfo := "fusion"
	postdata := Message_list{}
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

	var s []string
	s, _ = GetAllFile("./key/"+vaultID, s)

	c.JSON(200, s)
}
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
